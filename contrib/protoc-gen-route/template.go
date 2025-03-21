package main

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"
)

//go:embed template.go.tpl
var routeTemplate string

type serviceDesc struct {
	ServiceType string // Greeter
	ServiceName string // helloworld.Greeter
	Metadata    string // api/helloworld/helloworld.proto
	Methods     []*methodDesc
	MethodSets  map[string]*methodDesc

	OptionsKey       string
	RequestType      string
	ResponseType     string
	ExtraDataType    string
	NewExtraDataFunc string
}

type methodDesc struct {
	// method
	Name         string
	OriginalName string // The parsed original name
	Num          int
	Request      string
	Reply        string
	Comment      string

	// extra data
	Extra map[string]string
}

func (s *serviceDesc) execute() string {
	s.MethodSets = make(map[string]*methodDesc)
	for _, m := range s.Methods {
		s.MethodSets[m.Name] = m
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("route").Parse(strings.TrimSpace(routeTemplate))
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(buf, s)
	if err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}
