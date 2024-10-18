// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: dash/v1/admin.proto

package dashv1

import (
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

type Admin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Avatar   string   `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Username string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Nickname string   `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Roles    []string `protobuf:"bytes,5,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *Admin) Reset() {
	*x = Admin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dash_v1_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Admin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Admin) ProtoMessage() {}

func (x *Admin) ProtoReflect() protoreflect.Message {
	mi := &file_dash_v1_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Admin.ProtoReflect.Descriptor instead.
func (*Admin) Descriptor() ([]byte, []int) {
	return file_dash_v1_admin_proto_rawDescGZIP(), []int{0}
}

func (x *Admin) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Admin) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *Admin) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Admin) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *Admin) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

type AdminListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AdminListRequest) Reset() {
	*x = AdminListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dash_v1_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminListRequest) ProtoMessage() {}

func (x *AdminListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dash_v1_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminListRequest.ProtoReflect.Descriptor instead.
func (*AdminListRequest) Descriptor() ([]byte, []int) {
	return file_dash_v1_admin_proto_rawDescGZIP(), []int{1}
}

type AdminListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Admins []*Admin `protobuf:"bytes,1,rep,name=admins,proto3" json:"admins,omitempty"`
}

func (x *AdminListResponse) Reset() {
	*x = AdminListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dash_v1_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminListResponse) ProtoMessage() {}

func (x *AdminListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dash_v1_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminListResponse.ProtoReflect.Descriptor instead.
func (*AdminListResponse) Descriptor() ([]byte, []int) {
	return file_dash_v1_admin_proto_rawDescGZIP(), []int{2}
}

func (x *AdminListResponse) GetAdmins() []*Admin {
	if x != nil {
		return x.Admins
	}
	return nil
}

type AdminCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Avatar   string   `protobuf:"bytes,1,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Username string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Nickname string   `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Password string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Roles    []string `protobuf:"bytes,5,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *AdminCreateRequest) Reset() {
	*x = AdminCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dash_v1_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminCreateRequest) ProtoMessage() {}

func (x *AdminCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dash_v1_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminCreateRequest.ProtoReflect.Descriptor instead.
func (*AdminCreateRequest) Descriptor() ([]byte, []int) {
	return file_dash_v1_admin_proto_rawDescGZIP(), []int{3}
}

func (x *AdminCreateRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *AdminCreateRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AdminCreateRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *AdminCreateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AdminCreateRequest) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

type AdminCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Admin *Admin `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
}

func (x *AdminCreateResponse) Reset() {
	*x = AdminCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dash_v1_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminCreateResponse) ProtoMessage() {}

func (x *AdminCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dash_v1_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminCreateResponse.ProtoReflect.Descriptor instead.
func (*AdminCreateResponse) Descriptor() ([]byte, []int) {
	return file_dash_v1_admin_proto_rawDescGZIP(), []int{4}
}

func (x *AdminCreateResponse) GetAdmin() *Admin {
	if x != nil {
		return x.Admin
	}
	return nil
}

type AdminUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64    `protobuf:"varint,1,opt,name=id,proto3" json:"-" uri:"id"`
	Avatar   string   `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Username string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Nickname string   `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Password string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Roles    []string `protobuf:"bytes,6,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *AdminUpdateRequest) Reset() {
	*x = AdminUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dash_v1_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminUpdateRequest) ProtoMessage() {}

func (x *AdminUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dash_v1_admin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminUpdateRequest.ProtoReflect.Descriptor instead.
func (*AdminUpdateRequest) Descriptor() ([]byte, []int) {
	return file_dash_v1_admin_proto_rawDescGZIP(), []int{5}
}

func (x *AdminUpdateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AdminUpdateRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *AdminUpdateRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AdminUpdateRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *AdminUpdateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AdminUpdateRequest) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

type AdminUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Admin *Admin `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
}

func (x *AdminUpdateResponse) Reset() {
	*x = AdminUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dash_v1_admin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminUpdateResponse) ProtoMessage() {}

func (x *AdminUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dash_v1_admin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminUpdateResponse.ProtoReflect.Descriptor instead.
func (*AdminUpdateResponse) Descriptor() ([]byte, []int) {
	return file_dash_v1_admin_proto_rawDescGZIP(), []int{6}
}

func (x *AdminUpdateResponse) GetAdmin() *Admin {
	if x != nil {
		return x.Admin
	}
	return nil
}

type AdminDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"-" uri:"id"`
}

func (x *AdminDetailRequest) Reset() {
	*x = AdminDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dash_v1_admin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminDetailRequest) ProtoMessage() {}

func (x *AdminDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dash_v1_admin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminDetailRequest.ProtoReflect.Descriptor instead.
func (*AdminDetailRequest) Descriptor() ([]byte, []int) {
	return file_dash_v1_admin_proto_rawDescGZIP(), []int{7}
}

func (x *AdminDetailRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AdminDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Admin *Admin `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
}

func (x *AdminDetailResponse) Reset() {
	*x = AdminDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dash_v1_admin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminDetailResponse) ProtoMessage() {}

func (x *AdminDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dash_v1_admin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminDetailResponse.ProtoReflect.Descriptor instead.
func (*AdminDetailResponse) Descriptor() ([]byte, []int) {
	return file_dash_v1_admin_proto_rawDescGZIP(), []int{8}
}

func (x *AdminDetailResponse) GetAdmin() *Admin {
	if x != nil {
		return x.Admin
	}
	return nil
}

type AdminDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"-" uri:"id"`
}

func (x *AdminDeleteRequest) Reset() {
	*x = AdminDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dash_v1_admin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminDeleteRequest) ProtoMessage() {}

func (x *AdminDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dash_v1_admin_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminDeleteRequest.ProtoReflect.Descriptor instead.
func (*AdminDeleteRequest) Descriptor() ([]byte, []int) {
	return file_dash_v1_admin_proto_rawDescGZIP(), []int{9}
}

func (x *AdminDeleteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AdminDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AdminDeleteResponse) Reset() {
	*x = AdminDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dash_v1_admin_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminDeleteResponse) ProtoMessage() {}

func (x *AdminDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dash_v1_admin_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminDeleteResponse.ProtoReflect.Descriptor instead.
func (*AdminDeleteResponse) Descriptor() ([]byte, []int) {
	return file_dash_v1_admin_proto_rawDescGZIP(), []int{10}
}

var File_dash_v1_admin_proto protoreflect.FileDescriptor

var file_dash_v1_admin_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x61, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x61, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x05,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x3b, 0x0a, 0x11, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x52, 0x06, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x22, 0x96, 0x01, 0x0a,
	0x12, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x3b, 0x0a, 0x13, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x61,
	0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x05, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x22, 0xa6, 0x01, 0x0a, 0x12, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x3b, 0x0a, 0x13, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x24, 0x0a, 0x12, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b,
	0x0a, 0x13, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x24, 0x0a, 0x12, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x15, 0x0a, 0x13, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x94, 0x04, 0x0a, 0x0c, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x09, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x66, 0x0a, 0x0b, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x6b,
	0x0a, 0x0b, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e,
	0x64, 0x61, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x61, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b,
	0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x68, 0x0a, 0x0b, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x2e, 0x64, 0x61, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x68, 0x0a, 0x0b, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x2a, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42,
	0x83, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x42,
	0x0a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x62, 0x78, 0x61, 0x72, 0x6b,
	0x2f, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x73, 0x68,
	0x2f, 0x76, 0x31, 0x3b, 0x64, 0x61, 0x73, 0x68, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x58, 0x58,
	0xaa, 0x02, 0x07, 0x44, 0x61, 0x73, 0x68, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x44, 0x61, 0x73,
	0x68, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x44, 0x61, 0x73, 0x68, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x44, 0x61, 0x73,
	0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dash_v1_admin_proto_rawDescOnce sync.Once
	file_dash_v1_admin_proto_rawDescData = file_dash_v1_admin_proto_rawDesc
)

func file_dash_v1_admin_proto_rawDescGZIP() []byte {
	file_dash_v1_admin_proto_rawDescOnce.Do(func() {
		file_dash_v1_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_dash_v1_admin_proto_rawDescData)
	})
	return file_dash_v1_admin_proto_rawDescData
}

var file_dash_v1_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_dash_v1_admin_proto_goTypes = []interface{}{
	(*Admin)(nil),               // 0: dash.v1.Admin
	(*AdminListRequest)(nil),    // 1: dash.v1.AdminListRequest
	(*AdminListResponse)(nil),   // 2: dash.v1.AdminListResponse
	(*AdminCreateRequest)(nil),  // 3: dash.v1.AdminCreateRequest
	(*AdminCreateResponse)(nil), // 4: dash.v1.AdminCreateResponse
	(*AdminUpdateRequest)(nil),  // 5: dash.v1.AdminUpdateRequest
	(*AdminUpdateResponse)(nil), // 6: dash.v1.AdminUpdateResponse
	(*AdminDetailRequest)(nil),  // 7: dash.v1.AdminDetailRequest
	(*AdminDetailResponse)(nil), // 8: dash.v1.AdminDetailResponse
	(*AdminDeleteRequest)(nil),  // 9: dash.v1.AdminDeleteRequest
	(*AdminDeleteResponse)(nil), // 10: dash.v1.AdminDeleteResponse
}
var file_dash_v1_admin_proto_depIdxs = []int32{
	0,  // 0: dash.v1.AdminListResponse.admins:type_name -> dash.v1.Admin
	0,  // 1: dash.v1.AdminCreateResponse.admin:type_name -> dash.v1.Admin
	0,  // 2: dash.v1.AdminUpdateResponse.admin:type_name -> dash.v1.Admin
	0,  // 3: dash.v1.AdminDetailResponse.admin:type_name -> dash.v1.Admin
	1,  // 4: dash.v1.AdminService.AdminList:input_type -> dash.v1.AdminListRequest
	3,  // 5: dash.v1.AdminService.AdminCreate:input_type -> dash.v1.AdminCreateRequest
	5,  // 6: dash.v1.AdminService.AdminUpdate:input_type -> dash.v1.AdminUpdateRequest
	7,  // 7: dash.v1.AdminService.AdminDetail:input_type -> dash.v1.AdminDetailRequest
	9,  // 8: dash.v1.AdminService.AdminDelete:input_type -> dash.v1.AdminDeleteRequest
	2,  // 9: dash.v1.AdminService.AdminList:output_type -> dash.v1.AdminListResponse
	4,  // 10: dash.v1.AdminService.AdminCreate:output_type -> dash.v1.AdminCreateResponse
	6,  // 11: dash.v1.AdminService.AdminUpdate:output_type -> dash.v1.AdminUpdateResponse
	8,  // 12: dash.v1.AdminService.AdminDetail:output_type -> dash.v1.AdminDetailResponse
	10, // 13: dash.v1.AdminService.AdminDelete:output_type -> dash.v1.AdminDeleteResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_dash_v1_admin_proto_init() }
func file_dash_v1_admin_proto_init() {
	if File_dash_v1_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dash_v1_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Admin); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dash_v1_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dash_v1_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dash_v1_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminCreateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dash_v1_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminCreateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dash_v1_admin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminUpdateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dash_v1_admin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminUpdateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dash_v1_admin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminDetailRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dash_v1_admin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminDetailResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dash_v1_admin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminDeleteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dash_v1_admin_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminDeleteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dash_v1_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dash_v1_admin_proto_goTypes,
		DependencyIndexes: file_dash_v1_admin_proto_depIdxs,
		MessageInfos:      file_dash_v1_admin_proto_msgTypes,
	}.Build()
	File_dash_v1_admin_proto = out.File
	file_dash_v1_admin_proto_rawDesc = nil
	file_dash_v1_admin_proto_goTypes = nil
	file_dash_v1_admin_proto_depIdxs = nil
}
