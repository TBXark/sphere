package route

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/TBXark/sphere/cmd/protoc-gen-route/generate/template"
	"github.com/TBXark/sphere/proto/options/sphere/options"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	deprecationComment = "// Deprecated: Do not use."
)

const (
	contextPackage = protogen.GoImportPath("context")
)

var methodSets = make(map[string]int)

func GenerateFile(gen *protogen.Plugin, file *protogen.File, conf *Config) *protogen.GeneratedFile {
	if len(file.Services) == 0 || (!hasOptionsRule(file.Services, conf.OptionsKey)) {
		return nil
	}
	filename := file.GeneratedFilenamePrefix + conf.FileSuffix
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	generateFileHeader(gen, file, g)
	generateFileContent(file, g, conf)
	return g
}

func generateFileHeader(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile) {
	g.P("// Code generated by protoc-gen-route. DO NOT EDIT.")
	g.P("// versions:")
	g.P("// - protoc             ", protocVersion(gen))
	if file.Proto.GetOptions().GetDeprecated() {
		g.P("// ", file.Desc.Path(), " is a deprecated file.")
	} else {
		g.P("// source: ", file.Desc.Path())
	}
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
}

func generateFileContent(file *protogen.File, g *protogen.GeneratedFile, conf *Config) {
	if len(file.Services) == 0 {
		return
	}
	generateGoImport(g, conf)
	packageDesc := &template.PackageDesc{
		RequestType:  g.QualifiedGoIdent(conf.RequestType),
		ResponseType: g.QualifiedGoIdent(conf.ResponseType),
	}
	if conf.ExtraType.GoName != "" {
		packageDesc.ExtraDataType = g.QualifiedGoIdent(conf.ExtraType)
		packageDesc.NewExtraDataFunc = g.QualifiedGoIdent(conf.ExtraConstructor)
	}
	for _, service := range file.Services {
		generateService(g, service, conf, packageDesc)
	}
}

func generateGoImport(g *protogen.GeneratedFile, conf *Config) {
	didImport := make(map[protogen.GoImportPath]bool)
	g.P("var _ = new(", contextPackage.Ident("Context"), ")")
	for _, ident := range []protogen.GoIdent{conf.RequestType, conf.ResponseType, conf.ExtraType} {
		if !didImport[ident.GoImportPath] {
			didImport[ident.GoImportPath] = true
			g.P("var _ = new(", ident, ")")
		}
	}
	for _, ident := range []protogen.GoIdent{conf.ExtraConstructor} {
		if !didImport[ident.GoImportPath] {
			didImport[ident.GoImportPath] = true
			g.P("var _ = ", ident)
		}
	}
	g.P()
}

func generateService(g *protogen.GeneratedFile, service *protogen.Service, conf *Config, packageDesc *template.PackageDesc) {
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P("//")
		g.P(deprecationComment)
	}
	sd := &template.ServiceDesc{
		OptionsKey:  pascalCase(conf.OptionsKey),
		ServiceType: service.GoName,
		ServiceName: string(service.Desc.FullName()),
		Package:     packageDesc,
	}

	for _, method := range service.Methods {
		rule := extractOptionsRule(method, conf.OptionsKey)
		if rule == nil {
			continue
		}
		sd.Methods = append(sd.Methods, &template.MethodDesc{
			Name:         method.GoName,
			OriginalName: string(method.Desc.Name()),
			Num:          methodSets[method.GoName],
			Request:      g.QualifiedGoIdent(method.Input.GoIdent),
			Reply:        g.QualifiedGoIdent(method.Output.GoIdent),
			Comment:      methodCommend(method),
			Extra:        rule.Extra,
		})
		methodSets[method.GoName] += 1
	}
	if len(sd.Methods) != 0 {
		g.P(sd.Execute())
	}
}

func methodCommend(m *protogen.Method) string {
	leading := string(m.Comments.Leading)
	if leading == "" {
		return ""
	}
	cmp := strings.Split(strings.TrimSuffix(leading, "\n"), "\n")
	if len(cmp) == 0 {
		return ""
	}
	var lines []string
	lines = append(lines, fmt.Sprintf("// %s %s", m.Desc.Name(), strings.TrimSpace(cmp[0])))
	if len(cmp) > 1 {
		for _, line := range cmp[1:] {
			if strings.TrimSpace(line) == "" {
				continue
			}
			lines = append(lines, fmt.Sprintf("// %s", strings.TrimSpace(line)))
		}
	}
	return strings.Join(lines, "\n")
}

func hasOptionsRule(services []*protogen.Service, key string) bool {
	for _, service := range services {
		for _, method := range service.Methods {
			if extractOptionsRule(method, key) != nil {
				return true
			}
		}
	}
	return false
}

func extractOptionsRule(method *protogen.Method, key string) *options.KeyValuePair {
	if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
		return nil
	}
	if !proto.HasExtension(method.Desc.Options(), options.E_Options) {
		return nil
	}
	rules, ok := proto.GetExtension(method.Desc.Options(), options.E_Options).([]*options.KeyValuePair)
	if rules == nil || !ok {
		return nil
	}
	for _, rule := range rules {
		if rule.GetKey() == key {
			return rule
		}
	}
	return nil
}

func pascalCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '_' || r == '-'
	})
	if len(words) == 0 {
		return ""
	}
	result := ""
	for i := 0; i < len(words); i++ {
		word := words[i]
		if word == "" {
			continue
		}
		word = strings.ToLower(word)
		word = string(unicode.ToUpper(rune(word[0]))) + word[1:]
		result += word
	}
	return result
}

func protocVersion(gen *protogen.Plugin) string {
	v := gen.Request.GetCompilerVersion()
	if v == nil {
		return "(unknown)"
	}
	var suffix string
	if s := v.GetSuffix(); s != "" {
		suffix = "-" + s
	}
	return fmt.Sprintf("v%d.%d.%d%s", v.GetMajor(), v.GetMinor(), v.GetPatch(), suffix)
}
