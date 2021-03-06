// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: core.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// task req
type TaskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId   int64  `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	TaskType int32  `protobuf:"varint,2,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`
	TaskData []byte `protobuf:"bytes,3,opt,name=task_data,json=taskData,proto3" json:"task_data,omitempty"`
}

func (x *TaskReq) Reset() {
	*x = TaskReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskReq) ProtoMessage() {}

func (x *TaskReq) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskReq.ProtoReflect.Descriptor instead.
func (*TaskReq) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{0}
}

func (x *TaskReq) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *TaskReq) GetTaskType() int32 {
	if x != nil {
		return x.TaskType
	}
	return 0
}

func (x *TaskReq) GetTaskData() []byte {
	if x != nil {
		return x.TaskData
	}
	return nil
}

// task resp stream
type TaskResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resp []byte `protobuf:"bytes,3,opt,name=resp,proto3" json:"resp,omitempty"`
}

func (x *TaskResp) Reset() {
	*x = TaskResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResp) ProtoMessage() {}

func (x *TaskResp) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResp.ProtoReflect.Descriptor instead.
func (*TaskResp) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{1}
}

func (x *TaskResp) GetResp() []byte {
	if x != nil {
		return x.Resp
	}
	return nil
}

type RegistryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip        string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port      int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Weight    int32  `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Hostname  string `protobuf:"bytes,4,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Version   string `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	HostGroup string `protobuf:"bytes,6,opt,name=host_group,json=hostGroup,proto3" json:"host_group,omitempty"`
	Remark    string `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *RegistryReq) Reset() {
	*x = RegistryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryReq) ProtoMessage() {}

func (x *RegistryReq) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryReq.ProtoReflect.Descriptor instead.
func (*RegistryReq) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{2}
}

func (x *RegistryReq) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *RegistryReq) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *RegistryReq) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *RegistryReq) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *RegistryReq) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *RegistryReq) GetHostGroup() string {
	if x != nil {
		return x.HostGroup
	}
	return ""
}

func (x *RegistryReq) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type HeartbeatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip          string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port        int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	RunningTask []string `protobuf:"bytes,3,rep,name=running_task,json=runningTask,proto3" json:"running_task,omitempty"`
}

func (x *HeartbeatReq) Reset() {
	*x = HeartbeatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatReq) ProtoMessage() {}

func (x *HeartbeatReq) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatReq.ProtoReflect.Descriptor instead.
func (*HeartbeatReq) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{3}
}

func (x *HeartbeatReq) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *HeartbeatReq) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *HeartbeatReq) GetRunningTask() []string {
	if x != nil {
		return x.RunningTask
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{4}
}

var File_core_proto protoreflect.FileDescriptor

var file_core_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x07, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x1e, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x65, 0x73, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x72, 0x65, 0x73,
	0x70, 0x22, 0xb6, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x55, 0x0a, 0x0c, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73,
	0x6b, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x36, 0x0a, 0x04, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x28, 0x01,
	0x30, 0x01, 0x32, 0x6d, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12,
	0x2e, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x30, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_proto_rawDescOnce sync.Once
	file_core_proto_rawDescData = file_core_proto_rawDesc
)

func file_core_proto_rawDescGZIP() []byte {
	file_core_proto_rawDescOnce.Do(func() {
		file_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_proto_rawDescData)
	})
	return file_core_proto_rawDescData
}

var file_core_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_core_proto_goTypes = []interface{}{
	(*TaskReq)(nil),      // 0: proto.TaskReq
	(*TaskResp)(nil),     // 1: proto.TaskResp
	(*RegistryReq)(nil),  // 2: proto.RegistryReq
	(*HeartbeatReq)(nil), // 3: proto.HeartbeatReq
	(*Empty)(nil),        // 4: proto.Empty
}
var file_core_proto_depIdxs = []int32{
	0, // 0: proto.Task.RunTask:input_type -> proto.TaskReq
	2, // 1: proto.Heartbeat.Registry:input_type -> proto.RegistryReq
	3, // 2: proto.Heartbeat.Heartbeat:input_type -> proto.HeartbeatReq
	1, // 3: proto.Task.RunTask:output_type -> proto.TaskResp
	4, // 4: proto.Heartbeat.Registry:output_type -> proto.Empty
	4, // 5: proto.Heartbeat.Heartbeat:output_type -> proto.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_core_proto_init() }
func file_core_proto_init() {
	if File_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskReq); i {
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
		file_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskResp); i {
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
		file_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistryReq); i {
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
		file_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatReq); i {
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
		file_core_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_core_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_core_proto_goTypes,
		DependencyIndexes: file_core_proto_depIdxs,
		MessageInfos:      file_core_proto_msgTypes,
	}.Build()
	File_core_proto = out.File
	file_core_proto_rawDesc = nil
	file_core_proto_goTypes = nil
	file_core_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TaskClient is the client API for Task service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskClient interface {
	// run task return stream bytes
	RunTask(ctx context.Context, opts ...grpc.CallOption) (Task_RunTaskClient, error)
}

type taskClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskClient(cc grpc.ClientConnInterface) TaskClient {
	return &taskClient{cc}
}

func (c *taskClient) RunTask(ctx context.Context, opts ...grpc.CallOption) (Task_RunTaskClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Task_serviceDesc.Streams[0], "/proto.Task/RunTask", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskRunTaskClient{stream}
	return x, nil
}

type Task_RunTaskClient interface {
	Send(*TaskReq) error
	Recv() (*TaskResp, error)
	grpc.ClientStream
}

type taskRunTaskClient struct {
	grpc.ClientStream
}

func (x *taskRunTaskClient) Send(m *TaskReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *taskRunTaskClient) Recv() (*TaskResp, error) {
	m := new(TaskResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TaskServer is the server API for Task service.
type TaskServer interface {
	// run task return stream bytes
	RunTask(Task_RunTaskServer) error
}

// UnimplementedTaskServer can be embedded to have forward compatible implementations.
type UnimplementedTaskServer struct {
}

func (*UnimplementedTaskServer) RunTask(Task_RunTaskServer) error {
	return status.Errorf(codes.Unimplemented, "method RunTask not implemented")
}

func RegisterTaskServer(s *grpc.Server, srv TaskServer) {
	s.RegisterService(&_Task_serviceDesc, srv)
}

func _Task_RunTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TaskServer).RunTask(&taskRunTaskServer{stream})
}

type Task_RunTaskServer interface {
	Send(*TaskResp) error
	Recv() (*TaskReq, error)
	grpc.ServerStream
}

type taskRunTaskServer struct {
	grpc.ServerStream
}

func (x *taskRunTaskServer) Send(m *TaskResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *taskRunTaskServer) Recv() (*TaskReq, error) {
	m := new(TaskReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Task_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Task",
	HandlerType: (*TaskServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RunTask",
			Handler:       _Task_RunTask_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "core.proto",
}

// HeartbeatClient is the client API for Heartbeat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HeartbeatClient interface {
	// registry host
	Registry(ctx context.Context, in *RegistryReq, opts ...grpc.CallOption) (*Empty, error)
	// send heartbeat
	Heartbeat(ctx context.Context, in *HeartbeatReq, opts ...grpc.CallOption) (*Empty, error)
}

type heartbeatClient struct {
	cc grpc.ClientConnInterface
}

func NewHeartbeatClient(cc grpc.ClientConnInterface) HeartbeatClient {
	return &heartbeatClient{cc}
}

func (c *heartbeatClient) Registry(ctx context.Context, in *RegistryReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Heartbeat/Registry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heartbeatClient) Heartbeat(ctx context.Context, in *HeartbeatReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Heartbeat/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeartbeatServer is the server API for Heartbeat service.
type HeartbeatServer interface {
	// registry host
	Registry(context.Context, *RegistryReq) (*Empty, error)
	// send heartbeat
	Heartbeat(context.Context, *HeartbeatReq) (*Empty, error)
}

// UnimplementedHeartbeatServer can be embedded to have forward compatible implementations.
type UnimplementedHeartbeatServer struct {
}

func (*UnimplementedHeartbeatServer) Registry(context.Context, *RegistryReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Registry not implemented")
}
func (*UnimplementedHeartbeatServer) Heartbeat(context.Context, *HeartbeatReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}

func RegisterHeartbeatServer(s *grpc.Server, srv HeartbeatServer) {
	s.RegisterService(&_Heartbeat_serviceDesc, srv)
}

func _Heartbeat_Registry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartbeatServer).Registry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Heartbeat/Registry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartbeatServer).Registry(ctx, req.(*RegistryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Heartbeat_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartbeatServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Heartbeat/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartbeatServer).Heartbeat(ctx, req.(*HeartbeatReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Heartbeat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Heartbeat",
	HandlerType: (*HeartbeatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Registry",
			Handler:    _Heartbeat_Registry_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _Heartbeat_Heartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core.proto",
}
