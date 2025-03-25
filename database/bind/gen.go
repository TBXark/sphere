package bind

import (
	"fmt"
	"reflect"
	"strings"
	"text/template"
	"unicode"
)

func getPublicFields(obj interface{}) ([]string, map[string]reflect.StructField) {
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
			k := strings.ToLower(field.Name)
			keys = append(keys, k)
			fields[k] = field
		}
	}
	return keys, fields
}

func getPublicMethods(obj interface{}) ([]string, map[string]reflect.Method) {
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
		ptrType = reflect.PtrTo(typ)
	}

	for i := 0; i < structType.NumMethod(); i++ {
		method := structType.Method(i)
		if unicode.IsUpper(rune(method.Name[0])) {
			k := strings.ToLower(method.Name)
			keys = append(keys, k)
			methods[k] = method
		}
	}

	for i := 0; i < ptrType.NumMethod(); i++ {
		method := ptrType.Method(i)
		k := strings.ToLower(method.Name)
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
		return fmt.Sprintf("len(%s.%s) == 0", sourceName, field.Name)
	default:
		return fmt.Sprintf("reflect.ValueOf(%s.%s).IsZero()", sourceName, field.Name)
	}
}

type GenOptions struct {
	IgnoreSetZeroFields map[string]struct{}
	ClearOnNilFields    map[string]struct{}
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
		if o.IgnoreSetZeroFields == nil {
			o.IgnoreSetZeroFields = make(map[string]struct{}, len(fields))
		}
		for _, field := range fields {
			o.IgnoreSetZeroFields[field] = struct{}{}
		}
	}
}

func ClearOnNilField(fields ...string) Options {
	return func(o *GenOptions) {
		if o.ClearOnNilFields == nil {
			o.ClearOnNilFields = make(map[string]struct{}, len(fields))
		}
		for _, field := range fields {
			o.ClearOnNilFields[field] = struct{}{}
		}
	}
}

func (o *GenOptions) ClearOnNil(field string) bool {
	if o.ClearOnNilFields == nil {
		return false
	}
	_, ok := o.ClearOnNilFields[field]
	return ok
}

func (o *GenOptions) IgnoreSetZero(field string) bool {
	if o.IgnoreSetZeroFields == nil {
		return false
	}
	_, ok := o.IgnoreSetZeroFields[field]
	return ok
}

type GenConf struct {
	source       any
	target       any
	action       any
	ignoreFields []string
}

func NewGenConf(source, target, action any, ignoreFields ...string) GenConf {
	return GenConf{
		source:       source,
		target:       target,
		action:       action,
		ignoreFields: ignoreFields,
	}
}

func Gen(conf GenConf) string {
	actionName := getStructName(conf.action)
	sourceName := getStructName(conf.source)
	targetName := getStructName(conf.target)
	funcName := strings.Replace(actionName, sourceName, "", 1) + sourceName

	keys, sourceFields := getPublicFields(conf.source)
	_, targetFields := getPublicFields(conf.target)
	_, actionMethods := getPublicMethods(conf.action)

	context := GenContext{
		ActionName: actionName,
		SourceName: sourceName,
		TargetName: targetName,
		FuncName:   funcName,
		Fields:     make([]GenFieldContext, 0),
	}

	ignoreFields := make(map[string]bool, len(conf.ignoreFields))
	for _, field := range conf.ignoreFields {
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

		setter, hasSetter := actionMethods[strings.ToLower(fmt.Sprintf("Set%s", targetField.Name))]
		if !hasSetter {
			continue
		}
		settNillable, hasSettNillable := actionMethods[strings.ToLower(fmt.Sprintf("SetNillable%s", targetField.Name))]
		clearOnNil, hasClearOnNil := actionMethods[strings.ToLower(fmt.Sprintf("Clear%s", targetField.Name))]
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
			TargetSourceIsSomeType: targetField.Type.Kind() == sourceField.Type.Kind(),
		}
		context.Fields = append(context.Fields, field)
	}

	parse, err := template.New("gen").Funcs(template.FuncMap{
		"GenZeroCheck": genZeroCheck,
	}).Parse(genTemplate)
	if err != nil {
		return ""
	}
	var builder strings.Builder
	err = parse.Execute(&builder, context)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(builder.String())
}

type GenContext struct {
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
func {{.FuncName}}(source *ent.{{.ActionName}}, target *entpb.{{.TargetName}}, options ...bind.Options) *ent.{{.ActionName}} {
	option := NewGenOptions(options...)
{{- range .Fields -}}
{{- if .TargetFieldIsPtr}} {{/* 当目标字段是指针类型 */}}
	{{- if .CanSettNillable}} {{/* 如果存在SetNillable方法，直接使用 */}}
	{{- if .CanClearOnNil}} {{/* 如果存在ClearOnNil，判断是否需要使用 */}}
	if target.{{.TargetField.Name}} == nil && option.ClearOnNil("{{.SourceField.Name}}") {
		source.{{.ClearOnNilFuncName}}()
	} else {
		source.{{.SettNillableFuncName}}(target.{{.TargetField.Name}})
	}
	{{- else}}
		source.{{.SettNillableFuncName}}(target.{{.TargetField.Name}})
	{{- end}}
	{{- else}} {{/* 否则使用普通Setter方法，但需要解引用 */}}
	if target.{{.TargetField.Name}} != nil {
		source.{{.SetterFuncName}}(*target.{{.TargetField.Name}})
	}
	{{- end}}
{{- else -}} {{/* 当目标字段不是指针类型 */}}
    if !option.IgnoreSetZero("{{.SourceField.Name}}") || !({{GenZeroCheck "target" .TargetField}}) {
        {{- if .TargetSourceIsSomeType}} {{/* 如果源和目标是相同类型，直接赋值 */}}
		source.{{.SetterFuncName}}(target.{{.TargetField.Name}}) 
        {{- else}} {{/* 如果类型不同，需要进行类型转换 */}}
		source.{{.SetterFuncName}}({{.SourceField.Type.String}}(target.{{.TargetField.Name}}))
        {{- end}}
    }
{{- end}}
{{- end}}
	return source
}
`
