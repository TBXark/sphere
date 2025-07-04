package main

import (
	"flag"
	"fmt"

	"github.com/TBXark/sphere/cmd/protoc-gen-route/generate/route"
	"github.com/TBXark/sphere/cmd/protoc-gen-route/generate/template"
	"github.com/TBXark/sphere/internal/protogo"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var (
	showVersion = flag.Bool("version", false, "print the version and exit")

	optionsKey   = flag.String("options_key", "route", "options key in proto")
	fileSuffix   = flag.String("file_suffix", "_route.pb.go", "generated file suffix")
	templateFile = flag.String("template_file", "", "template file, if not set, use default template")

	requestModel   = flag.String("request_model", "", "request model")
	responseModel  = flag.String("response_model", "", "response model")
	extraDataModel = flag.String("extra_data_model", "", "extra data model")

	extraDataConstructor = flag.String("extra_data_constructor", "", "extra data constructor, and return a pointer of extra data")
)

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-route %v\n", "0.0.1")
		return
	}
	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gen *protogen.Plugin) error {
		conf := route.Config{
			OptionsKey:   *optionsKey,
			FileSuffix:   *fileSuffix,
			TemplateFile: *templateFile,

			RequestType:  protogo.NewGoIdent(*requestModel),
			ResponseType: protogo.NewGoIdent(*responseModel),
			ExtraType:    protogo.NewGoIdent(*extraDataModel),

			ExtraConstructor: protogo.NewGoIdent(*extraDataConstructor, protogo.IsFunc()),
		}
		if conf.RequestType == nil {
			return fmt.Errorf("flag request_model must be set")
		}
		if conf.ResponseType == nil {
			return fmt.Errorf("flag response_model must be set")
		}
		if conf.ExtraType != nil && conf.ExtraConstructor == nil {
			return fmt.Errorf("flag extra_data_constructor must be set if extra_data_model is set")
		}
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		template.ReplaceTemplateIfNeed(conf.TemplateFile)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			route.GenerateFile(gen, f, &conf)
		}
		return nil
	})
}
