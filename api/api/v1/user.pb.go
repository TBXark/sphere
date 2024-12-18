// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: api/v1/user.proto

package apiv1

import (
	v1 "github.com/TBXark/sphere/api/shared/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type MeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MeRequest) Reset() {
	*x = MeRequest{}
	mi := &file_api_v1_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeRequest) ProtoMessage() {}

func (x *MeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeRequest.ProtoReflect.Descriptor instead.
func (*MeRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_user_proto_rawDescGZIP(), []int{0}
}

type MeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *v1.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *MeResponse) Reset() {
	*x = MeResponse{}
	mi := &file_api_v1_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeResponse) ProtoMessage() {}

func (x *MeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeResponse.ProtoReflect.Descriptor instead.
func (*MeResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *MeResponse) GetUser() *v1.User {
	if x != nil {
		return x.User
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Avatar   string `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	mi := &file_api_v1_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UpdateRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *v1.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	mi := &file_api_v1_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateResponse) GetUser() *v1.User {
	if x != nil {
		return x.User
	}
	return nil
}

type BindPhoneWxMiniRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *BindPhoneWxMiniRequest) Reset() {
	*x = BindPhoneWxMiniRequest{}
	mi := &file_api_v1_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BindPhoneWxMiniRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindPhoneWxMiniRequest) ProtoMessage() {}

func (x *BindPhoneWxMiniRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindPhoneWxMiniRequest.ProtoReflect.Descriptor instead.
func (*BindPhoneWxMiniRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_user_proto_rawDescGZIP(), []int{4}
}

func (x *BindPhoneWxMiniRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type BindPhoneWxMiniResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BindPhoneWxMiniResponse) Reset() {
	*x = BindPhoneWxMiniResponse{}
	mi := &file_api_v1_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BindPhoneWxMiniResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindPhoneWxMiniResponse) ProtoMessage() {}

func (x *BindPhoneWxMiniResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindPhoneWxMiniResponse.ProtoReflect.Descriptor instead.
func (*BindPhoneWxMiniResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_user_proto_rawDescGZIP(), []int{5}
}

var File_api_v1_user_proto protoreflect.FileDescriptor

var file_api_v1_user_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x0b, 0x0a, 0x09, 0x4d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x0a,
	0x4d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x43, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x22, 0x35, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x2c, 0x0a, 0x16, 0x42,
	0x69, 0x6e, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x57, 0x78, 0x4d, 0x69, 0x6e, 0x69, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x19, 0x0a, 0x17, 0x42, 0x69, 0x6e,
	0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x57, 0x78, 0x4d, 0x69, 0x6e, 0x69, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa2, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x02, 0x4d, 0x65, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x12, 0x54, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x7a, 0x0a,
	0x0f, 0x42, 0x69, 0x6e, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x57, 0x78, 0x4d, 0x69, 0x6e, 0x69,
	0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x57, 0x78, 0x4d, 0x69, 0x6e, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x57, 0x78, 0x4d, 0x69, 0x6e, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x62, 0x69, 0x6e, 0x64, 0x2f, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x2f, 0x77, 0x78, 0x6d, 0x69, 0x6e, 0x69, 0x42, 0x7b, 0x0a, 0x0a, 0x63, 0x6f, 0x6d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x62, 0x78, 0x61, 0x72, 0x6b, 0x2f, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x06, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x41,
	0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_user_proto_rawDescOnce sync.Once
	file_api_v1_user_proto_rawDescData = file_api_v1_user_proto_rawDesc
)

func file_api_v1_user_proto_rawDescGZIP() []byte {
	file_api_v1_user_proto_rawDescOnce.Do(func() {
		file_api_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_user_proto_rawDescData)
	})
	return file_api_v1_user_proto_rawDescData
}

var file_api_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_v1_user_proto_goTypes = []any{
	(*MeRequest)(nil),               // 0: api.v1.MeRequest
	(*MeResponse)(nil),              // 1: api.v1.MeResponse
	(*UpdateRequest)(nil),           // 2: api.v1.UpdateRequest
	(*UpdateResponse)(nil),          // 3: api.v1.UpdateResponse
	(*BindPhoneWxMiniRequest)(nil),  // 4: api.v1.BindPhoneWxMiniRequest
	(*BindPhoneWxMiniResponse)(nil), // 5: api.v1.BindPhoneWxMiniResponse
	(*v1.User)(nil),                 // 6: shared.v1.User
}
var file_api_v1_user_proto_depIdxs = []int32{
	6, // 0: api.v1.MeResponse.user:type_name -> shared.v1.User
	6, // 1: api.v1.UpdateResponse.user:type_name -> shared.v1.User
	0, // 2: api.v1.UserService.Me:input_type -> api.v1.MeRequest
	2, // 3: api.v1.UserService.Update:input_type -> api.v1.UpdateRequest
	4, // 4: api.v1.UserService.BindPhoneWxMini:input_type -> api.v1.BindPhoneWxMiniRequest
	1, // 5: api.v1.UserService.Me:output_type -> api.v1.MeResponse
	3, // 6: api.v1.UserService.Update:output_type -> api.v1.UpdateResponse
	5, // 7: api.v1.UserService.BindPhoneWxMini:output_type -> api.v1.BindPhoneWxMiniResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1_user_proto_init() }
func file_api_v1_user_proto_init() {
	if File_api_v1_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_user_proto_goTypes,
		DependencyIndexes: file_api_v1_user_proto_depIdxs,
		MessageInfos:      file_api_v1_user_proto_msgTypes,
	}.Build()
	File_api_v1_user_proto = out.File
	file_api_v1_user_proto_rawDesc = nil
	file_api_v1_user_proto_goTypes = nil
	file_api_v1_user_proto_depIdxs = nil
}
