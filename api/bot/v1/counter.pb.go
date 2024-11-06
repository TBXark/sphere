// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: bot/v1/counter.proto

package botv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	mi := &file_bot_v1_counter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bot_v1_counter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRequest.ProtoReflect.Descriptor instead.
func (*StartRequest) Descriptor() ([]byte, []int) {
	return file_bot_v1_counter_proto_rawDescGZIP(), []int{0}
}

func (x *StartRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type StartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StartResponse) Reset() {
	*x = StartResponse{}
	mi := &file_bot_v1_counter_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartResponse) ProtoMessage() {}

func (x *StartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bot_v1_counter_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartResponse.ProtoReflect.Descriptor instead.
func (*StartResponse) Descriptor() ([]byte, []int) {
	return file_bot_v1_counter_proto_rawDescGZIP(), []int{1}
}

func (x *StartResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CounterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Step  int32 `protobuf:"varint,2,opt,name=step,proto3" json:"step,omitempty"`
}

func (x *CounterRequest) Reset() {
	*x = CounterRequest{}
	mi := &file_bot_v1_counter_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CounterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CounterRequest) ProtoMessage() {}

func (x *CounterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bot_v1_counter_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CounterRequest.ProtoReflect.Descriptor instead.
func (*CounterRequest) Descriptor() ([]byte, []int) {
	return file_bot_v1_counter_proto_rawDescGZIP(), []int{2}
}

func (x *CounterRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CounterRequest) GetStep() int32 {
	if x != nil {
		return x.Step
	}
	return 0
}

type CounterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CounterResponse) Reset() {
	*x = CounterResponse{}
	mi := &file_bot_v1_counter_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CounterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CounterResponse) ProtoMessage() {}

func (x *CounterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bot_v1_counter_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CounterResponse.ProtoReflect.Descriptor instead.
func (*CounterResponse) Descriptor() ([]byte, []int) {
	return file_bot_v1_counter_proto_rawDescGZIP(), []int{3}
}

func (x *CounterResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type UnknownRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnknownRequest) Reset() {
	*x = UnknownRequest{}
	mi := &file_bot_v1_counter_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnknownRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnknownRequest) ProtoMessage() {}

func (x *UnknownRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bot_v1_counter_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnknownRequest.ProtoReflect.Descriptor instead.
func (*UnknownRequest) Descriptor() ([]byte, []int) {
	return file_bot_v1_counter_proto_rawDescGZIP(), []int{4}
}

type UnknownResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnknownResponse) Reset() {
	*x = UnknownResponse{}
	mi := &file_bot_v1_counter_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnknownResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnknownResponse) ProtoMessage() {}

func (x *UnknownResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bot_v1_counter_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnknownResponse.ProtoReflect.Descriptor instead.
func (*UnknownResponse) Descriptor() ([]byte, []int) {
	return file_bot_v1_counter_proto_rawDescGZIP(), []int{5}
}

var File_bot_v1_counter_proto protoreflect.FileDescriptor

var file_bot_v1_counter_proto_rawDesc = []byte{
	0x0a, 0x14, 0x62, 0x6f, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x62, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x22,
	0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x0a,
	0x0e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x22, 0x27, 0x0a, 0x0f, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xbe, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x14, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x6f, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3a, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x62, 0x6f,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x07,
	0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x12, 0x16, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x7e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e,
	0x62, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x62, 0x78, 0x61, 0x72, 0x6b, 0x2f, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6f, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x6f, 0x74, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x42, 0x6f, 0x74, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x06, 0x42, 0x6f, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12, 0x42, 0x6f, 0x74, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x07, 0x42, 0x6f, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bot_v1_counter_proto_rawDescOnce sync.Once
	file_bot_v1_counter_proto_rawDescData = file_bot_v1_counter_proto_rawDesc
)

func file_bot_v1_counter_proto_rawDescGZIP() []byte {
	file_bot_v1_counter_proto_rawDescOnce.Do(func() {
		file_bot_v1_counter_proto_rawDescData = protoimpl.X.CompressGZIP(file_bot_v1_counter_proto_rawDescData)
	})
	return file_bot_v1_counter_proto_rawDescData
}

var file_bot_v1_counter_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_bot_v1_counter_proto_goTypes = []any{
	(*StartRequest)(nil),    // 0: bot.v1.StartRequest
	(*StartResponse)(nil),   // 1: bot.v1.StartResponse
	(*CounterRequest)(nil),  // 2: bot.v1.CounterRequest
	(*CounterResponse)(nil), // 3: bot.v1.CounterResponse
	(*UnknownRequest)(nil),  // 4: bot.v1.UnknownRequest
	(*UnknownResponse)(nil), // 5: bot.v1.UnknownResponse
}
var file_bot_v1_counter_proto_depIdxs = []int32{
	0, // 0: bot.v1.CounterService.Start:input_type -> bot.v1.StartRequest
	2, // 1: bot.v1.CounterService.Counter:input_type -> bot.v1.CounterRequest
	4, // 2: bot.v1.CounterService.Unknown:input_type -> bot.v1.UnknownRequest
	1, // 3: bot.v1.CounterService.Start:output_type -> bot.v1.StartResponse
	3, // 4: bot.v1.CounterService.Counter:output_type -> bot.v1.CounterResponse
	5, // 5: bot.v1.CounterService.Unknown:output_type -> bot.v1.UnknownResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bot_v1_counter_proto_init() }
func file_bot_v1_counter_proto_init() {
	if File_bot_v1_counter_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bot_v1_counter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bot_v1_counter_proto_goTypes,
		DependencyIndexes: file_bot_v1_counter_proto_depIdxs,
		MessageInfos:      file_bot_v1_counter_proto_msgTypes,
	}.Build()
	File_bot_v1_counter_proto = out.File
	file_bot_v1_counter_proto_rawDesc = nil
	file_bot_v1_counter_proto_goTypes = nil
	file_bot_v1_counter_proto_depIdxs = nil
}
