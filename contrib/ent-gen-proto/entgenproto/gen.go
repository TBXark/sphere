package entgenproto

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/protobuf/types/descriptorpb"
	"log"
	"path/filepath"
	"reflect"
	"sort"
	_ "unsafe"
)

//go:linkname generate entgo.io/contrib/entproto.(*Extension).generate
func generate(extension *entproto.Extension, g *gen.Graph) error

//go:linkname wktsPaths entgo.io/contrib/entproto.wktsPaths
var wktsPaths map[string]string

type Options struct {
	SchemaPath             string
	ProtoDir               string
	AllFieldsRequired      bool
	IgnoreUnsupportedJson  bool
	IgnoreUnsupportedType  bool
	UsageAnyForUnsupported bool
	AutoAddAnnotation      bool
	EnumUseRawType         bool
	TimeUseProtoType       string
}

func Generate(options *Options) {
	injectGoogleProtobufAny()
	abs, err := filepath.Abs(options.SchemaPath)
	if err != nil {
		log.Fatalf("entproto: failed getting absolute path: %v", err)
	}
	graph, err := entc.LoadGraph(options.SchemaPath, &gen.Config{
		Target: filepath.Dir(abs),
	})
	if err != nil {
		log.Fatalf("entproto: failed loading ent graph: %v", err)
	}
	if options.AutoAddAnnotation {
		for i := 0; i < len(graph.Nodes); i++ {
			addAnnotationForNode(graph.Nodes[i], options)
		}
	}
	extension, err := entproto.NewExtension(
		entproto.WithProtoDir(options.ProtoDir),
		entproto.SkipGenFile(),
	)
	if err != nil {
		log.Fatalf("entproto: failed creating entproto extension: %v", err)
	}
	err = generate(extension, graph)
	if err != nil {
		log.Fatalf("entproto: failed generating protos: %s", err)
	}
}

func addAnnotationForNode(node *gen.Type, options *Options) {
	if node.Annotations == nil {
		node.Annotations = make(map[string]interface{}, 1)
	}
	if node.Annotations[entproto.MessageAnnotation] != nil {
		return
	}

	maxExistNum := 0
	existNums := map[int]struct{}{}
	for _, fd := range node.Fields {
		if fd.Annotations != nil {
			if obj, exist := fd.Annotations[entproto.FieldAnnotation]; exist {
				pbField := struct {
					Number int
				}{}
				err := mapstructure.Decode(obj, &pbField)
				if err != nil {
					log.Fatalf("entproto: failed decoding field annotation: %v", err)
				}
				existNums[pbField.Number] = struct{}{}
				if pbField.Number > maxExistNum {
					maxExistNum = pbField.Number
				}
			}
		}
	}

	idGenerator := &fileIDGenerator{exist: existNums}

	// If the node does not have the message annotation, add it.
	node.Annotations[entproto.MessageAnnotation] = entproto.Message()
	if node.ID.Annotations == nil {
		node.ID.Annotations = make(map[string]interface{}, 1)
	}
	node.ID.Annotations[entproto.FieldAnnotation] = entproto.Field(idGenerator.MustNext())
	sort.Slice(node.Fields, func(i, j int) bool {
		if node.Fields[i].Position.MixedIn != node.Fields[j].Position.MixedIn {
			// MixedIn fields should be at the end of the list.
			return !node.Fields[i].Position.MixedIn
		}
		return node.Fields[i].Position.Index < node.Fields[j].Position.Index
	})

	removeFields := make([]int, 0)
	for j := 0; j < len(node.Fields); j++ {
		fd := node.Fields[j]
		if fd.Annotations == nil {
			fd.Annotations = make(map[string]interface{}, 1)
		}
		if fd.IsEnum() {
			fixEnumType(fd, options.EnumUseRawType)
		}
		if fd.Annotations[entproto.FieldAnnotation] != nil {
			continue
		}
		if fd.Annotations[entproto.SkipAnnotation] != nil {
			removeFields = append(removeFields, j)
			continue
		}
		if options.IgnoreUnsupportedType && (fd.Type.Type == field.TypeOther || fd.Type.Type == field.TypeInvalid) {
			removeFields = append(removeFields, j)
			continue
		}
		newType := fixFieldType(fd, options)
		if fd.Type.Type == field.TypeJSON && newType == field.TypeBytes {
			if options.UsageAnyForUnsupported {
				newType = field.TypeJSON
			} else if options.IgnoreUnsupportedJson {
				removeFields = append(removeFields, j)
			}
		}
		fd.Type.Type = newType
		var fieldOptions []entproto.FieldOption
		if fd.Type.Type == field.TypeJSON || fd.Type.Type == field.TypeOther || fd.Type.Type == field.TypeInvalid {
			if _, exist := buildInTypeSlice[fd.Type.RType.Ident]; !exist {
				fieldOptions = append(fieldOptions,
					entproto.Type(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE),
					entproto.TypeName("google.protobuf.Any"),
				)
			}
		}
		fd.Annotations[entproto.FieldAnnotation] = entproto.Field(idGenerator.MustNext(), fieldOptions...)
		if fd.Optional && options.AllFieldsRequired {
			fd.Optional = false
		}
	}
	for i := len(removeFields) - 1; i >= 0; i-- {
		node.Fields = append(node.Fields[:removeFields[i]], node.Fields[removeFields[i]+1:]...)
	}
}

func fixEnumType(fd *gen.Field, enumUseRawType bool) {
	if enumUseRawType {
		if fd.HasGoType() {
			fd.Type.Type = reflectKind2FieldType[fd.Type.RType.Kind]
		} else {
			fd.Type.Type = field.TypeString
		}
	} else {
		enums := make(map[string]int32, len(fd.Enums))
		for index, enum := range fd.Enums {
			enums[enum.Value] = int32(index) + 1
		}
		fd.Annotations[entproto.EnumAnnotation] = entproto.Enum(enums, entproto.OmitFieldPrefix())
	}
}

func fixFieldType(fd *gen.Field, options *Options) field.Type {
	switch fd.Type.Type {
	case field.TypeOther, field.TypeInvalid:
		// unsupported type will be converted to bytes
		return field.TypeBytes
	case field.TypeJSON:
		if fd.HasGoType() {
			switch fd.Type.RType.Kind {
			case reflect.Slice:
				// support for built-in types slice only
				if _, ok := buildInTypeSlice[fd.Type.RType.Ident]; ok {
					return field.TypeJSON
				}
				return field.TypeBytes
			default:
				return field.TypeBytes
			}
		}
		return field.TypeBytes
	case field.TypeUUID:
		return field.TypeString
	case field.TypeTime:
		switch options.TimeUseProtoType {
		case "int64":
			return field.TypeInt64
		case "string":
			return field.TypeString
		default:
			return field.TypeTime
		}
	default:
		return fd.Type.Type
	}
}

var reflectKind2FieldType = map[reflect.Kind]field.Type{
	reflect.Bool:          field.TypeBool,
	reflect.Int:           field.TypeInt,
	reflect.Int8:          field.TypeInt8,
	reflect.Int16:         field.TypeInt16,
	reflect.Int32:         field.TypeInt32,
	reflect.Int64:         field.TypeInt64,
	reflect.Uint:          field.TypeUint,
	reflect.Uint8:         field.TypeUint8,
	reflect.Uint16:        field.TypeUint16,
	reflect.Uint32:        field.TypeUint32,
	reflect.Uint64:        field.TypeUint64,
	reflect.Uintptr:       field.TypeUint,
	reflect.Float32:       field.TypeFloat32,
	reflect.Float64:       field.TypeFloat64,
	reflect.Complex64:     field.TypeOther,
	reflect.Complex128:    field.TypeOther,
	reflect.Array:         field.TypeJSON,
	reflect.Chan:          field.TypeOther,
	reflect.Func:          field.TypeOther,
	reflect.Interface:     field.TypeJSON,
	reflect.Map:           field.TypeJSON,
	reflect.Pointer:       field.TypeJSON,
	reflect.Slice:         field.TypeJSON,
	reflect.String:        field.TypeString,
	reflect.Struct:        field.TypeJSON,
	reflect.UnsafePointer: field.TypeOther,
}

var buildInTypeSlice = map[string]struct{}{
	"[]int":     {},
	"[]int8":    {},
	"[]int16":   {},
	"[]int32":   {},
	"[]int64":   {},
	"[]uint":    {},
	"[]uint8":   {},
	"[]uint16":  {},
	"[]uint32":  {},
	"[]uint64":  {},
	"[]float32": {},
	"[]float64": {},
	"[]string":  {},
	"[]bool":    {},
}

type fileIDGenerator struct {
	current int
	exist   map[int]struct{}
}

func (f *fileIDGenerator) Next() (int, error) {
	f.current++
	for {
		if _, ok := f.exist[f.current]; ok {
			f.current++
			continue
		}
		if f.current > 536870911 {
			return 0, fmt.Errorf("entproto: field number exceed the maximum value 536870911")
		}
		break
	}
	return f.current, nil
}

func (f *fileIDGenerator) MustNext() int {
	num, err := f.Next()
	if err != nil {
		panic(err)
	}
	return num
}

func injectGoogleProtobufAny() {
	wktsPaths["google.protobuf.Any"] = "google/protobuf/any.proto"
}
