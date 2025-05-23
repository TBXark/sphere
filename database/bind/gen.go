package bind

import (
	"fmt"
	"reflect"
	"strings"
	"text/template"
	"unicode"

	"github.com/iancoleman/strcase"
)

func getPublicFields(obj interface{}, keyMapper func(s string) string) ([]string, map[string]reflect.StructField) {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil, nil
	}
	typ := val.Type()
	keys := make([]string, 0)
	fields := make(map[string]reflect.StructField)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if unicode.IsUpper(rune(field.Name[0])) && !field.Anonymous {
			k := field.Name
			if keyMapper != nil {
				k = keyMapper(k)
			}
			keys = append(keys, k)
			fields[k] = field
		}
	}
	return keys, fields
}

func getPublicMethods(obj interface{}, keyMapper func(s string) string) ([]string, map[string]reflect.Method) {
	typ := reflect.TypeOf(obj)

	if typ == nil || (typ.Kind() != reflect.Struct && (typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct)) {
		return nil, nil
	}

	keys := make([]string, 0)
	methods := make(map[string]reflect.Method)

	structType := typ
	ptrType := typ
	if typ.Kind() == reflect.Ptr {
		structType = typ.Elem()
	} else {
		ptrType = reflect.PointerTo(typ)
	}

	for i := 0; i < structType.NumMethod(); i++ {
		method := structType.Method(i)
		if unicode.IsUpper(rune(method.Name[0])) {
			k := method.Name
			if keyMapper != nil {
				k = keyMapper(k)
			}
			keys = append(keys, k)
			methods[k] = method
		}
	}

	for i := 0; i < ptrType.NumMethod(); i++ {
		method := ptrType.Method(i)
		k := method.Name
		if keyMapper != nil {
			k = keyMapper(k)
		}
		if _, exists := methods[k]; !exists && unicode.IsUpper(rune(method.Name[0])) {
			keys = append(keys, k)
			methods[k] = method
		}
	}

	return keys, methods
}

func getStructName(value any) string {
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	if v.Kind() == reflect.Struct {
		return t.Name()
	}
	return "Unknown"
}

func genZeroCheck(sourceName string, field reflect.StructField) string {
	if field.Type.Kind() == reflect.Ptr {
		return fmt.Sprintf("%s.%s == nil", sourceName, field.Name)
	}
	switch field.Type.Kind() {
	case reflect.String:
		return fmt.Sprintf("%s.%s == \"\"", sourceName, field.Name)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%s.%s == 0", sourceName, field.Name)
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%s.%s == 0.0", sourceName, field.Name)
	case reflect.Bool:
		return fmt.Sprintf("!%s.%s", sourceName, field.Name)
	case reflect.Slice:
		return fmt.Sprintf("%s.%s == nil", sourceName, field.Name)
	default:
		return fmt.Sprintf("reflect.ValueOf(%s.%s).IsZero()", sourceName, field.Name)
	}
}

type GenOptions struct {
	ignoreSetZeroFields map[string]struct{}
	clearOnNilFields    map[string]struct{}
	ignoreFields        map[string]struct{}
	keepFieldsOnly      map[string]struct{}
}

func NewGenOptions(options ...Options) *GenOptions {
	o := &GenOptions{}
	for _, opt := range options {
		opt(o)
	}
	return o
}

type Options func(*GenOptions)

func IgnoreSetZeroField(fields ...string) Options {
	return func(o *GenOptions) {
		if o.ignoreSetZeroFields == nil {
			o.ignoreSetZeroFields = make(map[string]struct{}, len(fields))
		}
		for _, field := range fields {
			o.ignoreSetZeroFields[field] = struct{}{}
		}
	}
}

func ClearOnNilField(fields ...string) Options {
	return func(o *GenOptions) {
		if o.clearOnNilFields == nil {
			o.clearOnNilFields = make(map[string]struct{}, len(fields))
		}
		for _, field := range fields {
			o.clearOnNilFields[field] = struct{}{}
		}
	}
}

func IgnoreField(fields ...string) Options {
	return func(o *GenOptions) {
		if o.ignoreFields == nil {
			o.ignoreFields = make(map[string]struct{}, len(fields))
		}
		for _, field := range fields {
			o.ignoreFields[field] = struct{}{}
		}
	}
}

func KeepFieldsOnly(fields ...string) Options {
	return func(o *GenOptions) {
		if o.keepFieldsOnly == nil {
			o.keepFieldsOnly = make(map[string]struct{}, len(fields))
		}
		for _, field := range fields {
			o.keepFieldsOnly[field] = struct{}{}
		}
	}
}

func (o *GenOptions) ClearOnNil(field string) bool {
	if o.clearOnNilFields == nil {
		return false
	}
	_, ok := o.clearOnNilFields[field]
	return ok
}

func (o *GenOptions) IgnoreSetZero(field string) bool {
	if o.ignoreSetZeroFields == nil {
		return false
	}
	_, ok := o.ignoreSetZeroFields[field]
	return ok
}

func (o *GenOptions) CanSetField(field string) bool {
	if o.keepFieldsOnly != nil {
		_, ok := o.keepFieldsOnly[field]
		return ok
	}
	if o.ignoreFields == nil {
		return true
	}
	_, ok := o.ignoreFields[field]
	return !ok
}

type GenConf struct {
	source        any
	target        any
	action        any
	IgnoreFields  []string
	SourcePkgName string
	TargetPkgName string
}

func NewGenConf(source, target, action any) *GenConf {
	return &GenConf{
		source:        source,
		target:        target,
		action:        action,
		IgnoreFields:  nil,
		SourcePkgName: "ent",
		TargetPkgName: "entpb",
	}
}

func (c *GenConf) WithSourcePkgName(pkgName string) *GenConf {
	c.SourcePkgName = pkgName
	return c
}

func (c *GenConf) WithTargetPkgName(pkgName string) *GenConf {
	c.TargetPkgName = pkgName
	return c
}

func (c *GenConf) WithIgnoreFields(fields ...string) *GenConf {
	c.IgnoreFields = fields
	return c
}

func Gen(conf *GenConf) string {
	actionName := getStructName(conf.action)
	sourceName := getStructName(conf.source)
	targetName := getStructName(conf.target)
	funcName := strings.Replace(actionName, sourceName, "", 1) + sourceName

	keys, sourceFields := getPublicFields(conf.source, strcase.ToSnake)
	_, targetFields := getPublicFields(conf.target, strcase.ToSnake)
	_, actionMethods := getPublicMethods(conf.action, strcase.ToSnake)

	context := GenContext{
		SourcePkgName: conf.SourcePkgName,
		TargetPkgName: conf.TargetPkgName,

		ActionName: actionName,
		SourceName: sourceName,
		TargetName: targetName,
		FuncName:   funcName,
		Fields:     make([]GenFieldContext, 0),
	}

	ignoreFields := make(map[string]bool, len(conf.IgnoreFields))
	for _, field := range conf.IgnoreFields {
		ignoreFields[strings.ToLower(field)] = true
	}

	for _, n := range keys {
		if ignoreFields[n] {
			continue
		}
		targetField, ok := targetFields[n]
		if !ok {
			continue
		}
		sourceField, ok := sourceFields[n]
		if !ok {
			continue
		}

		setter, hasSetter := actionMethods[strcase.ToSnake(fmt.Sprintf("Set%s", targetField.Name))]
		if !hasSetter {
			continue
		}
		settNillable, hasSettNillable := actionMethods[strcase.ToSnake(fmt.Sprintf("SetNillable%s", targetField.Name))]
		clearOnNil, hasClearOnNil := actionMethods[strcase.ToSnake(fmt.Sprintf("Clear%s", targetField.Name))]
		targetFieldIsPtr := targetField.Type.Kind() == reflect.Ptr

		field := GenFieldContext{
			TargetField: targetField,
			SourceField: sourceField,

			SetterFuncName:       setter.Name,
			SettNillableFuncName: settNillable.Name,
			ClearOnNilFuncName:   clearOnNil.Name,

			CanSettNillable:        hasSettNillable,
			CanClearOnNil:          hasClearOnNil,
			TargetFieldIsPtr:       targetFieldIsPtr,
			TargetSourceIsSomeType: false,
		}
		if targetFieldIsPtr {
			elem := targetField.Type.Elem()
			field.TargetSourceIsSomeType = elem.Kind() == sourceField.Type.Kind() && elem.String() == sourceField.Type.String()
		} else {
			field.TargetSourceIsSomeType = targetField.Type.Kind() == sourceField.Type.Kind() && targetField.Type.String() == sourceField.Type.String()
		}
		context.Fields = append(context.Fields, field)
	}

	parse, err := template.New("gen").Funcs(template.FuncMap{
		"GenZeroCheck": genZeroCheck,
		"ToSnakeCase":  strcase.ToSnake,
	}).Parse(genTemplate)
	if err != nil {
		return ""
	}
	var builder strings.Builder
	err = parse.Execute(&builder, context)
	if err != nil {
		return ""
	}
	return builder.String()
}

type GenContext struct {
	SourcePkgName string
	TargetPkgName string

	ActionName string
	SourceName string
	TargetName string
	FuncName   string

	Fields []GenFieldContext
}

type GenFieldContext struct {
	TargetField reflect.StructField
	SourceField reflect.StructField

	SetterFuncName       string
	SettNillableFuncName string
	ClearOnNilFuncName   string

	CanSettNillable bool
	CanClearOnNil   bool

	TargetFieldIsPtr       bool
	TargetSourceIsSomeType bool
}

const genTemplate = `
func {{.FuncName}}(source *{{.SourcePkgName}}.{{.ActionName}}, target *{{.TargetPkgName}}.{{.TargetName}}, options ...bind.Options) *{{.SourcePkgName}}.{{.ActionName}} {
	option := bind.NewGenOptions(options...)
{{- range .Fields}}
	if option.CanSetField("{{ToSnakeCase .SourceField.Name}}") {
		{{- if .TargetFieldIsPtr}} {{/* 当目标字段是指针类型 */}}
			{{- if .CanSettNillable}} {{/* 如果存在SetNillable方法，直接使用 */}}
				{{- if .CanClearOnNil}} {{/* 如果存在ClearOnNil，判断是否需要使用 */}}
					if target.{{.TargetField.Name}} == nil && option.ClearOnNil("{{ToSnakeCase .SourceField.Name}}") {
						source.{{.ClearOnNilFuncName}}()
					} else {
						source.{{.SettNillableFuncName}}(target.{{.TargetField.Name}})
					}
				{{- else}}
					source.{{.SettNillableFuncName}}(target.{{.TargetField.Name}})
				{{- end}}
			{{- else}} {{/* 否则使用普通Setter方法，但需要解引用 */}}
				if target.{{.TargetField.Name}} != nil {
        			{{- if .TargetSourceIsSomeType}} {{/* 如果源和目标是相同类型，直接赋值 */}}
						source.{{.SetterFuncName}}(*target.{{.TargetField.Name}}) 
        			{{- else}} {{/* 如果类型不同，需要进行类型转换 */}}
						source.{{.SetterFuncName}}({{.SourceField.Type.String}}(*target.{{.TargetField.Name}}))
        			{{- end}}
				}
			{{- end}}
		{{- else -}} {{/* 当目标字段不是指针类型 */}}
			if !option.IgnoreSetZero("{{ToSnakeCase .SourceField.Name}}") || !({{GenZeroCheck "target" .TargetField}}) {
        		{{- if .TargetSourceIsSomeType}} {{/* 如果源和目标是相同类型，直接赋值 */}}
					source.{{.SetterFuncName}}(target.{{.TargetField.Name}}) 
        		{{- else}} {{/* 如果类型不同，需要进行类型转换 */}}
					source.{{.SetterFuncName}}({{.SourceField.Type.String}}(target.{{.TargetField.Name}}))
        		{{- end}}
    		}
		{{- end}}
	}
{{- end}}
	return source
}
`
