// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.14.0
// source: process.proto

package proto

import (
	context "context"
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

type ProcessAction int32

const (
	ProcessAction_PROCESS_ACTION_UNKNOWN ProcessAction = 0
	ProcessAction_PROCESS_ACTION_STARTED ProcessAction = 1
	ProcessAction_PROCESS_ACTION_STOPPED ProcessAction = 2
)

// Enum value maps for ProcessAction.
var (
	ProcessAction_name = map[int32]string{
		0: "PROCESS_ACTION_UNKNOWN",
		1: "PROCESS_ACTION_STARTED",
		2: "PROCESS_ACTION_STOPPED",
	}
	ProcessAction_value = map[string]int32{
		"PROCESS_ACTION_UNKNOWN": 0,
		"PROCESS_ACTION_STARTED": 1,
		"PROCESS_ACTION_STOPPED": 2,
	}
)

func (x ProcessAction) Enum() *ProcessAction {
	p := new(ProcessAction)
	*p = x
	return p
}

func (x ProcessAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProcessAction) Descriptor() protoreflect.EnumDescriptor {
	return file_process_proto_enumTypes[0].Descriptor()
}

func (ProcessAction) Type() protoreflect.EnumType {
	return &file_process_proto_enumTypes[0]
}

func (x ProcessAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProcessAction.Descriptor instead.
func (ProcessAction) EnumDescriptor() ([]byte, []int) {
	return file_process_proto_rawDescGZIP(), []int{0}
}

type ProcessInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid       int32  `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Command   string `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	Arguments string `protobuf:"bytes,3,opt,name=arguments,proto3" json:"arguments,omitempty"`
}

func (x *ProcessInfo) Reset() {
	*x = ProcessInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_process_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessInfo) ProtoMessage() {}

func (x *ProcessInfo) ProtoReflect() protoreflect.Message {
	mi := &file_process_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessInfo.ProtoReflect.Descriptor instead.
func (*ProcessInfo) Descriptor() ([]byte, []int) {
	return file_process_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessInfo) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *ProcessInfo) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *ProcessInfo) GetArguments() string {
	if x != nil {
		return x.Arguments
	}
	return ""
}

type ProcessStateStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *ProcessStateStreamRequest) Reset() {
	*x = ProcessStateStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_process_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessStateStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessStateStreamRequest) ProtoMessage() {}

func (x *ProcessStateStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_process_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessStateStreamRequest.ProtoReflect.Descriptor instead.
func (*ProcessStateStreamRequest) Descriptor() ([]byte, []int) {
	return file_process_proto_rawDescGZIP(), []int{1}
}

func (x *ProcessStateStreamRequest) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

type ProcessStateStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Process *ProcessInfo  `protobuf:"bytes,1,opt,name=process,proto3" json:"process,omitempty"`
	Action  ProcessAction `protobuf:"varint,2,opt,name=action,proto3,enum=proto.ProcessAction" json:"action,omitempty"`
	Error   string        `protobuf:"bytes,15,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ProcessStateStreamResponse) Reset() {
	*x = ProcessStateStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_process_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessStateStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessStateStreamResponse) ProtoMessage() {}

func (x *ProcessStateStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_process_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessStateStreamResponse.ProtoReflect.Descriptor instead.
func (*ProcessStateStreamResponse) Descriptor() ([]byte, []int) {
	return file_process_proto_rawDescGZIP(), []int{2}
}

func (x *ProcessStateStreamResponse) GetProcess() *ProcessInfo {
	if x != nil {
		return x.Process
	}
	return nil
}

func (x *ProcessStateStreamResponse) GetAction() ProcessAction {
	if x != nil {
		return x.Action
	}
	return ProcessAction_PROCESS_ACTION_UNKNOWN
}

func (x *ProcessStateStreamResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ProcessStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *ProcessStateRequest) Reset() {
	*x = ProcessStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_process_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessStateRequest) ProtoMessage() {}

func (x *ProcessStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_process_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessStateRequest.ProtoReflect.Descriptor instead.
func (*ProcessStateRequest) Descriptor() ([]byte, []int) {
	return file_process_proto_rawDescGZIP(), []int{3}
}

func (x *ProcessStateRequest) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

type ProcessStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Processes []*ProcessInfo `protobuf:"bytes,1,rep,name=processes,proto3" json:"processes,omitempty"`
}

func (x *ProcessStateResponse) Reset() {
	*x = ProcessStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_process_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessStateResponse) ProtoMessage() {}

func (x *ProcessStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_process_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessStateResponse.ProtoReflect.Descriptor instead.
func (*ProcessStateResponse) Descriptor() ([]byte, []int) {
	return file_process_proto_rawDescGZIP(), []int{4}
}

func (x *ProcessStateResponse) GetProcesses() []*ProcessInfo {
	if x != nil {
		return x.Processes
	}
	return nil
}

var File_process_proto protoreflect.FileDescriptor

var file_process_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x45,
	0x0a, 0x19, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x8e, 0x01, 0x0a, 0x1a, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x3f, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a,
	0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x48, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x30, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x2a, 0x63, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1a,
	0x0a, 0x16, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x52,
	0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x4f,
	0x50, 0x50, 0x45, 0x44, 0x10, 0x02, 0x32, 0xaf, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x5b, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12,
	0x47, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_process_proto_rawDescOnce sync.Once
	file_process_proto_rawDescData = file_process_proto_rawDesc
)

func file_process_proto_rawDescGZIP() []byte {
	file_process_proto_rawDescOnce.Do(func() {
		file_process_proto_rawDescData = protoimpl.X.CompressGZIP(file_process_proto_rawDescData)
	})
	return file_process_proto_rawDescData
}

var file_process_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_process_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_process_proto_goTypes = []interface{}{
	(ProcessAction)(0),                 // 0: proto.ProcessAction
	(*ProcessInfo)(nil),                // 1: proto.ProcessInfo
	(*ProcessStateStreamRequest)(nil),  // 2: proto.ProcessStateStreamRequest
	(*ProcessStateStreamResponse)(nil), // 3: proto.ProcessStateStreamResponse
	(*ProcessStateRequest)(nil),        // 4: proto.ProcessStateRequest
	(*ProcessStateResponse)(nil),       // 5: proto.ProcessStateResponse
	(*Session)(nil),                    // 6: proto.Session
}
var file_process_proto_depIdxs = []int32{
	6, // 0: proto.ProcessStateStreamRequest.session:type_name -> proto.Session
	1, // 1: proto.ProcessStateStreamResponse.process:type_name -> proto.ProcessInfo
	0, // 2: proto.ProcessStateStreamResponse.action:type_name -> proto.ProcessAction
	6, // 3: proto.ProcessStateRequest.session:type_name -> proto.Session
	1, // 4: proto.ProcessStateResponse.processes:type_name -> proto.ProcessInfo
	2, // 5: proto.Process.ProcessStateStream:input_type -> proto.ProcessStateStreamRequest
	4, // 6: proto.Process.ProcessState:input_type -> proto.ProcessStateRequest
	3, // 7: proto.Process.ProcessStateStream:output_type -> proto.ProcessStateStreamResponse
	5, // 8: proto.Process.ProcessState:output_type -> proto.ProcessStateResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_process_proto_init() }
func file_process_proto_init() {
	if File_process_proto != nil {
		return
	}
	file_session_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_process_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessInfo); i {
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
		file_process_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessStateStreamRequest); i {
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
		file_process_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessStateStreamResponse); i {
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
		file_process_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessStateRequest); i {
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
		file_process_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessStateResponse); i {
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
			RawDescriptor: file_process_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_process_proto_goTypes,
		DependencyIndexes: file_process_proto_depIdxs,
		EnumInfos:         file_process_proto_enumTypes,
		MessageInfos:      file_process_proto_msgTypes,
	}.Build()
	File_process_proto = out.File
	file_process_proto_rawDesc = nil
	file_process_proto_goTypes = nil
	file_process_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProcessClient is the client API for Process service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProcessClient interface {
	// stream updates to processes as they happen
	ProcessStateStream(ctx context.Context, in *ProcessStateStreamRequest, opts ...grpc.CallOption) (Process_ProcessStateStreamClient, error)
	// get a list of all processes
	ProcessState(ctx context.Context, in *ProcessStateRequest, opts ...grpc.CallOption) (*ProcessStateResponse, error)
}

type processClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessClient(cc grpc.ClientConnInterface) ProcessClient {
	return &processClient{cc}
}

func (c *processClient) ProcessStateStream(ctx context.Context, in *ProcessStateStreamRequest, opts ...grpc.CallOption) (Process_ProcessStateStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Process_serviceDesc.Streams[0], "/proto.Process/ProcessStateStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &processProcessStateStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Process_ProcessStateStreamClient interface {
	Recv() (*ProcessStateStreamResponse, error)
	grpc.ClientStream
}

type processProcessStateStreamClient struct {
	grpc.ClientStream
}

func (x *processProcessStateStreamClient) Recv() (*ProcessStateStreamResponse, error) {
	m := new(ProcessStateStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *processClient) ProcessState(ctx context.Context, in *ProcessStateRequest, opts ...grpc.CallOption) (*ProcessStateResponse, error) {
	out := new(ProcessStateResponse)
	err := c.cc.Invoke(ctx, "/proto.Process/ProcessState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessServer is the server API for Process service.
type ProcessServer interface {
	// stream updates to processes as they happen
	ProcessStateStream(*ProcessStateStreamRequest, Process_ProcessStateStreamServer) error
	// get a list of all processes
	ProcessState(context.Context, *ProcessStateRequest) (*ProcessStateResponse, error)
}

// UnimplementedProcessServer can be embedded to have forward compatible implementations.
type UnimplementedProcessServer struct {
}

func (*UnimplementedProcessServer) ProcessStateStream(*ProcessStateStreamRequest, Process_ProcessStateStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ProcessStateStream not implemented")
}
func (*UnimplementedProcessServer) ProcessState(context.Context, *ProcessStateRequest) (*ProcessStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessState not implemented")
}

func RegisterProcessServer(s *grpc.Server, srv ProcessServer) {
	s.RegisterService(&_Process_serviceDesc, srv)
}

func _Process_ProcessStateStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProcessStateStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProcessServer).ProcessStateStream(m, &processProcessStateStreamServer{stream})
}

type Process_ProcessStateStreamServer interface {
	Send(*ProcessStateStreamResponse) error
	grpc.ServerStream
}

type processProcessStateStreamServer struct {
	grpc.ServerStream
}

func (x *processProcessStateStreamServer) Send(m *ProcessStateStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Process_ProcessState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessServer).ProcessState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Process/ProcessState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessServer).ProcessState(ctx, req.(*ProcessStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Process_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Process",
	HandlerType: (*ProcessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessState",
			Handler:    _Process_ProcessState_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProcessStateStream",
			Handler:       _Process_ProcessStateStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "process.proto",
}
