// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: dash/v1/task.proto

package dashv1

import (
	entpb "github.com/tbxark/sphere/api/entpb"
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

type TaskListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size int32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *TaskListRequest) Reset() {
	*x = TaskListRequest{}
	mi := &file_dash_v1_task_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListRequest) ProtoMessage() {}

func (x *TaskListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dash_v1_task_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListRequest.ProtoReflect.Descriptor instead.
func (*TaskListRequest) Descriptor() ([]byte, []int) {
	return file_dash_v1_task_proto_rawDescGZIP(), []int{0}
}

func (x *TaskListRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *TaskListRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type TaskListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*entpb.Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *TaskListResponse) Reset() {
	*x = TaskListResponse{}
	mi := &file_dash_v1_task_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListResponse) ProtoMessage() {}

func (x *TaskListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dash_v1_task_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListResponse.ProtoReflect.Descriptor instead.
func (*TaskListResponse) Descriptor() ([]byte, []int) {
	return file_dash_v1_task_proto_rawDescGZIP(), []int{1}
}

func (x *TaskListResponse) GetTasks() []*entpb.Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type TaskDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"-" uri:"id"`  
}

func (x *TaskDetailRequest) Reset() {
	*x = TaskDetailRequest{}
	mi := &file_dash_v1_task_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDetailRequest) ProtoMessage() {}

func (x *TaskDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dash_v1_task_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDetailRequest.ProtoReflect.Descriptor instead.
func (*TaskDetailRequest) Descriptor() ([]byte, []int) {
	return file_dash_v1_task_proto_rawDescGZIP(), []int{2}
}

func (x *TaskDetailRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type TaskDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *entpb.Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *TaskDetailResponse) Reset() {
	*x = TaskDetailResponse{}
	mi := &file_dash_v1_task_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDetailResponse) ProtoMessage() {}

func (x *TaskDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dash_v1_task_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDetailResponse.ProtoReflect.Descriptor instead.
func (*TaskDetailResponse) Descriptor() ([]byte, []int) {
	return file_dash_v1_task_proto_rawDescGZIP(), []int{3}
}

func (x *TaskDetailResponse) GetTask() *entpb.Task {
	if x != nil {
		return x.Task
	}
	return nil
}

var File_dash_v1_task_proto protoreflect.FileDescriptor

var file_dash_v1_task_proto_rawDesc = []byte{
	0x0a, 0x12, 0x64, 0x61, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x61, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x65, 0x6e, 0x74,
	0x70, 0x62, 0x2f, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39,
	0x0a, 0x0f, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x35, 0x0a, 0x10, 0x54, 0x61, 0x73,
	0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a,
	0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x22, 0x23, 0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x12, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x74,
	0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x65, 0x6e, 0x74, 0x70,
	0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x32, 0xcc, 0x01, 0x0a,
	0x0b, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x08,
	0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x64, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x1a, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x82, 0x01, 0x0a, 0x0b,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x54, 0x61, 0x73,
	0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x62, 0x78, 0x61, 0x72, 0x6b, 0x2f, 0x73, 0x70, 0x68, 0x65,
	0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x64,
	0x61, 0x73, 0x68, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x44, 0x61,
	0x73, 0x68, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x44, 0x61, 0x73, 0x68, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x13, 0x44, 0x61, 0x73, 0x68, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x44, 0x61, 0x73, 0x68, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dash_v1_task_proto_rawDescOnce sync.Once
	file_dash_v1_task_proto_rawDescData = file_dash_v1_task_proto_rawDesc
)

func file_dash_v1_task_proto_rawDescGZIP() []byte {
	file_dash_v1_task_proto_rawDescOnce.Do(func() {
		file_dash_v1_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_dash_v1_task_proto_rawDescData)
	})
	return file_dash_v1_task_proto_rawDescData
}

var file_dash_v1_task_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_dash_v1_task_proto_goTypes = []any{
	(*TaskListRequest)(nil),    // 0: dash.v1.TaskListRequest
	(*TaskListResponse)(nil),   // 1: dash.v1.TaskListResponse
	(*TaskDetailRequest)(nil),  // 2: dash.v1.TaskDetailRequest
	(*TaskDetailResponse)(nil), // 3: dash.v1.TaskDetailResponse
	(*entpb.Task)(nil),         // 4: entpb.Task
}
var file_dash_v1_task_proto_depIdxs = []int32{
	4, // 0: dash.v1.TaskListResponse.tasks:type_name -> entpb.Task
	4, // 1: dash.v1.TaskDetailResponse.task:type_name -> entpb.Task
	0, // 2: dash.v1.TaskService.TaskList:input_type -> dash.v1.TaskListRequest
	2, // 3: dash.v1.TaskService.TaskDetail:input_type -> dash.v1.TaskDetailRequest
	1, // 4: dash.v1.TaskService.TaskList:output_type -> dash.v1.TaskListResponse
	3, // 5: dash.v1.TaskService.TaskDetail:output_type -> dash.v1.TaskDetailResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_dash_v1_task_proto_init() }
func file_dash_v1_task_proto_init() {
	if File_dash_v1_task_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dash_v1_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dash_v1_task_proto_goTypes,
		DependencyIndexes: file_dash_v1_task_proto_depIdxs,
		MessageInfos:      file_dash_v1_task_proto_msgTypes,
	}.Build()
	File_dash_v1_task_proto = out.File
	file_dash_v1_task_proto_rawDesc = nil
	file_dash_v1_task_proto_goTypes = nil
	file_dash_v1_task_proto_depIdxs = nil
}