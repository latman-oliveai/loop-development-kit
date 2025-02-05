// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.14.0
// source: keyboard.proto

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

type KeyboardHotkey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// bit 0  = altL
	// bit 1  = altR
	// bit 2  = alt
	// bit 3  = ctrlL
	// bit 4  = ctrlR
	// bit 5  = ctrl
	// bit 6  = metaL
	// bit 7  = metaR
	// bit 8  = meta
	// bit 9  = shiftL
	// bit 10 = shiftR
	// bit 11 = shift
	Modifiers int32 `protobuf:"varint,2,opt,name=modifiers,proto3" json:"modifiers,omitempty"`
}

func (x *KeyboardHotkey) Reset() {
	*x = KeyboardHotkey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keyboard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyboardHotkey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyboardHotkey) ProtoMessage() {}

func (x *KeyboardHotkey) ProtoReflect() protoreflect.Message {
	mi := &file_keyboard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyboardHotkey.ProtoReflect.Descriptor instead.
func (*KeyboardHotkey) Descriptor() ([]byte, []int) {
	return file_keyboard_proto_rawDescGZIP(), []int{0}
}

func (x *KeyboardHotkey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyboardHotkey) GetModifiers() int32 {
	if x != nil {
		return x.Modifiers
	}
	return 0
}

type KeyboardHotkeyStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session        `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	Hotkey  *KeyboardHotkey `protobuf:"bytes,2,opt,name=hotkey,proto3" json:"hotkey,omitempty"`
}

func (x *KeyboardHotkeyStreamRequest) Reset() {
	*x = KeyboardHotkeyStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keyboard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyboardHotkeyStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyboardHotkeyStreamRequest) ProtoMessage() {}

func (x *KeyboardHotkeyStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_keyboard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyboardHotkeyStreamRequest.ProtoReflect.Descriptor instead.
func (*KeyboardHotkeyStreamRequest) Descriptor() ([]byte, []int) {
	return file_keyboard_proto_rawDescGZIP(), []int{1}
}

func (x *KeyboardHotkeyStreamRequest) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

func (x *KeyboardHotkeyStreamRequest) GetHotkey() *KeyboardHotkey {
	if x != nil {
		return x.Hotkey
	}
	return nil
}

type KeyboardHotkeyStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scanned bool   `protobuf:"varint,1,opt,name=scanned,proto3" json:"scanned,omitempty"`
	Error   string `protobuf:"bytes,15,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *KeyboardHotkeyStreamResponse) Reset() {
	*x = KeyboardHotkeyStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keyboard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyboardHotkeyStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyboardHotkeyStreamResponse) ProtoMessage() {}

func (x *KeyboardHotkeyStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_keyboard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyboardHotkeyStreamResponse.ProtoReflect.Descriptor instead.
func (*KeyboardHotkeyStreamResponse) Descriptor() ([]byte, []int) {
	return file_keyboard_proto_rawDescGZIP(), []int{2}
}

func (x *KeyboardHotkeyStreamResponse) GetScanned() bool {
	if x != nil {
		return x.Scanned
	}
	return false
}

func (x *KeyboardHotkeyStreamResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type KeyboardTextStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *KeyboardTextStreamRequest) Reset() {
	*x = KeyboardTextStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keyboard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyboardTextStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyboardTextStreamRequest) ProtoMessage() {}

func (x *KeyboardTextStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_keyboard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyboardTextStreamRequest.ProtoReflect.Descriptor instead.
func (*KeyboardTextStreamRequest) Descriptor() ([]byte, []int) {
	return file_keyboard_proto_rawDescGZIP(), []int{3}
}

func (x *KeyboardTextStreamRequest) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

type KeyboardTextStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text  string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Error string `protobuf:"bytes,15,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *KeyboardTextStreamResponse) Reset() {
	*x = KeyboardTextStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keyboard_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyboardTextStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyboardTextStreamResponse) ProtoMessage() {}

func (x *KeyboardTextStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_keyboard_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyboardTextStreamResponse.ProtoReflect.Descriptor instead.
func (*KeyboardTextStreamResponse) Descriptor() ([]byte, []int) {
	return file_keyboard_proto_rawDescGZIP(), []int{4}
}

func (x *KeyboardTextStreamResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *KeyboardTextStreamResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type KeyboardCharacterStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *KeyboardCharacterStreamRequest) Reset() {
	*x = KeyboardCharacterStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keyboard_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyboardCharacterStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyboardCharacterStreamRequest) ProtoMessage() {}

func (x *KeyboardCharacterStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_keyboard_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyboardCharacterStreamRequest.ProtoReflect.Descriptor instead.
func (*KeyboardCharacterStreamRequest) Descriptor() ([]byte, []int) {
	return file_keyboard_proto_rawDescGZIP(), []int{5}
}

func (x *KeyboardCharacterStreamRequest) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

type KeyboardCharacterStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text  string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Error string `protobuf:"bytes,15,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *KeyboardCharacterStreamResponse) Reset() {
	*x = KeyboardCharacterStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keyboard_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyboardCharacterStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyboardCharacterStreamResponse) ProtoMessage() {}

func (x *KeyboardCharacterStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_keyboard_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyboardCharacterStreamResponse.ProtoReflect.Descriptor instead.
func (*KeyboardCharacterStreamResponse) Descriptor() ([]byte, []int) {
	return file_keyboard_proto_rawDescGZIP(), []int{6}
}

func (x *KeyboardCharacterStreamResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *KeyboardCharacterStreamResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_keyboard_proto protoreflect.FileDescriptor

var file_keyboard_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x0e, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x48, 0x6f, 0x74, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x22, 0x76, 0x0a, 0x1b, 0x4b, 0x65, 0x79, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x48, 0x6f, 0x74, 0x6b, 0x65, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x2d, 0x0a, 0x06, 0x68, 0x6f, 0x74, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x48, 0x6f, 0x74, 0x6b, 0x65, 0x79, 0x52, 0x06, 0x68, 0x6f, 0x74, 0x6b, 0x65, 0x79,
	0x22, 0x4e, 0x0a, 0x1c, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x48, 0x6f, 0x74, 0x6b,
	0x65, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x45, 0x0a, 0x19, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x54, 0x65, 0x78, 0x74,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a,
	0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x46, 0x0a, 0x1a, 0x4b, 0x65, 0x79, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x54, 0x65, 0x78, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x4a, 0x0a, 0x1e, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x4b, 0x0a, 0x1f, 0x4b,
	0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xb6, 0x02, 0x0a, 0x08, 0x4b, 0x65, 0x79,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x61, 0x0a, 0x14, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x48, 0x6f, 0x74, 0x6b, 0x65, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x22, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x48, 0x6f,
	0x74, 0x6b, 0x65, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x48, 0x6f, 0x74, 0x6b, 0x65, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x5b, 0x0a, 0x12, 0x4b, 0x65, 0x79, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x54, 0x65, 0x78, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x20,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x54,
	0x65, 0x78, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x54, 0x65, 0x78, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x6a, 0x0a, 0x17, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_keyboard_proto_rawDescOnce sync.Once
	file_keyboard_proto_rawDescData = file_keyboard_proto_rawDesc
)

func file_keyboard_proto_rawDescGZIP() []byte {
	file_keyboard_proto_rawDescOnce.Do(func() {
		file_keyboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_keyboard_proto_rawDescData)
	})
	return file_keyboard_proto_rawDescData
}

var file_keyboard_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_keyboard_proto_goTypes = []interface{}{
	(*KeyboardHotkey)(nil),                  // 0: proto.KeyboardHotkey
	(*KeyboardHotkeyStreamRequest)(nil),     // 1: proto.KeyboardHotkeyStreamRequest
	(*KeyboardHotkeyStreamResponse)(nil),    // 2: proto.KeyboardHotkeyStreamResponse
	(*KeyboardTextStreamRequest)(nil),       // 3: proto.KeyboardTextStreamRequest
	(*KeyboardTextStreamResponse)(nil),      // 4: proto.KeyboardTextStreamResponse
	(*KeyboardCharacterStreamRequest)(nil),  // 5: proto.KeyboardCharacterStreamRequest
	(*KeyboardCharacterStreamResponse)(nil), // 6: proto.KeyboardCharacterStreamResponse
	(*Session)(nil),                         // 7: proto.Session
}
var file_keyboard_proto_depIdxs = []int32{
	7, // 0: proto.KeyboardHotkeyStreamRequest.session:type_name -> proto.Session
	0, // 1: proto.KeyboardHotkeyStreamRequest.hotkey:type_name -> proto.KeyboardHotkey
	7, // 2: proto.KeyboardTextStreamRequest.session:type_name -> proto.Session
	7, // 3: proto.KeyboardCharacterStreamRequest.session:type_name -> proto.Session
	1, // 4: proto.Keyboard.KeyboardHotkeyStream:input_type -> proto.KeyboardHotkeyStreamRequest
	3, // 5: proto.Keyboard.KeyboardTextStream:input_type -> proto.KeyboardTextStreamRequest
	5, // 6: proto.Keyboard.KeyboardCharacterStream:input_type -> proto.KeyboardCharacterStreamRequest
	2, // 7: proto.Keyboard.KeyboardHotkeyStream:output_type -> proto.KeyboardHotkeyStreamResponse
	4, // 8: proto.Keyboard.KeyboardTextStream:output_type -> proto.KeyboardTextStreamResponse
	6, // 9: proto.Keyboard.KeyboardCharacterStream:output_type -> proto.KeyboardCharacterStreamResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_keyboard_proto_init() }
func file_keyboard_proto_init() {
	if File_keyboard_proto != nil {
		return
	}
	file_session_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_keyboard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyboardHotkey); i {
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
		file_keyboard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyboardHotkeyStreamRequest); i {
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
		file_keyboard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyboardHotkeyStreamResponse); i {
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
		file_keyboard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyboardTextStreamRequest); i {
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
		file_keyboard_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyboardTextStreamResponse); i {
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
		file_keyboard_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyboardCharacterStreamRequest); i {
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
		file_keyboard_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyboardCharacterStreamResponse); i {
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
			RawDescriptor: file_keyboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_keyboard_proto_goTypes,
		DependencyIndexes: file_keyboard_proto_depIdxs,
		MessageInfos:      file_keyboard_proto_msgTypes,
	}.Build()
	File_keyboard_proto = out.File
	file_keyboard_proto_rawDesc = nil
	file_keyboard_proto_goTypes = nil
	file_keyboard_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KeyboardClient is the client API for Keyboard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeyboardClient interface {
	// register a hotkey and receive streamed messages when it is pressed
	KeyboardHotkeyStream(ctx context.Context, in *KeyboardHotkeyStreamRequest, opts ...grpc.CallOption) (Keyboard_KeyboardHotkeyStreamClient, error)
	// stream chunks of text when the user finishes typing them
	KeyboardTextStream(ctx context.Context, in *KeyboardTextStreamRequest, opts ...grpc.CallOption) (Keyboard_KeyboardTextStreamClient, error)
	// stream text as it is typed
	KeyboardCharacterStream(ctx context.Context, in *KeyboardCharacterStreamRequest, opts ...grpc.CallOption) (Keyboard_KeyboardCharacterStreamClient, error)
}

type keyboardClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyboardClient(cc grpc.ClientConnInterface) KeyboardClient {
	return &keyboardClient{cc}
}

func (c *keyboardClient) KeyboardHotkeyStream(ctx context.Context, in *KeyboardHotkeyStreamRequest, opts ...grpc.CallOption) (Keyboard_KeyboardHotkeyStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Keyboard_serviceDesc.Streams[0], "/proto.Keyboard/KeyboardHotkeyStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &keyboardKeyboardHotkeyStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Keyboard_KeyboardHotkeyStreamClient interface {
	Recv() (*KeyboardHotkeyStreamResponse, error)
	grpc.ClientStream
}

type keyboardKeyboardHotkeyStreamClient struct {
	grpc.ClientStream
}

func (x *keyboardKeyboardHotkeyStreamClient) Recv() (*KeyboardHotkeyStreamResponse, error) {
	m := new(KeyboardHotkeyStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *keyboardClient) KeyboardTextStream(ctx context.Context, in *KeyboardTextStreamRequest, opts ...grpc.CallOption) (Keyboard_KeyboardTextStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Keyboard_serviceDesc.Streams[1], "/proto.Keyboard/KeyboardTextStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &keyboardKeyboardTextStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Keyboard_KeyboardTextStreamClient interface {
	Recv() (*KeyboardTextStreamResponse, error)
	grpc.ClientStream
}

type keyboardKeyboardTextStreamClient struct {
	grpc.ClientStream
}

func (x *keyboardKeyboardTextStreamClient) Recv() (*KeyboardTextStreamResponse, error) {
	m := new(KeyboardTextStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *keyboardClient) KeyboardCharacterStream(ctx context.Context, in *KeyboardCharacterStreamRequest, opts ...grpc.CallOption) (Keyboard_KeyboardCharacterStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Keyboard_serviceDesc.Streams[2], "/proto.Keyboard/KeyboardCharacterStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &keyboardKeyboardCharacterStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Keyboard_KeyboardCharacterStreamClient interface {
	Recv() (*KeyboardCharacterStreamResponse, error)
	grpc.ClientStream
}

type keyboardKeyboardCharacterStreamClient struct {
	grpc.ClientStream
}

func (x *keyboardKeyboardCharacterStreamClient) Recv() (*KeyboardCharacterStreamResponse, error) {
	m := new(KeyboardCharacterStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KeyboardServer is the server API for Keyboard service.
type KeyboardServer interface {
	// register a hotkey and receive streamed messages when it is pressed
	KeyboardHotkeyStream(*KeyboardHotkeyStreamRequest, Keyboard_KeyboardHotkeyStreamServer) error
	// stream chunks of text when the user finishes typing them
	KeyboardTextStream(*KeyboardTextStreamRequest, Keyboard_KeyboardTextStreamServer) error
	// stream text as it is typed
	KeyboardCharacterStream(*KeyboardCharacterStreamRequest, Keyboard_KeyboardCharacterStreamServer) error
}

// UnimplementedKeyboardServer can be embedded to have forward compatible implementations.
type UnimplementedKeyboardServer struct {
}

func (*UnimplementedKeyboardServer) KeyboardHotkeyStream(*KeyboardHotkeyStreamRequest, Keyboard_KeyboardHotkeyStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method KeyboardHotkeyStream not implemented")
}
func (*UnimplementedKeyboardServer) KeyboardTextStream(*KeyboardTextStreamRequest, Keyboard_KeyboardTextStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method KeyboardTextStream not implemented")
}
func (*UnimplementedKeyboardServer) KeyboardCharacterStream(*KeyboardCharacterStreamRequest, Keyboard_KeyboardCharacterStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method KeyboardCharacterStream not implemented")
}

func RegisterKeyboardServer(s *grpc.Server, srv KeyboardServer) {
	s.RegisterService(&_Keyboard_serviceDesc, srv)
}

func _Keyboard_KeyboardHotkeyStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(KeyboardHotkeyStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KeyboardServer).KeyboardHotkeyStream(m, &keyboardKeyboardHotkeyStreamServer{stream})
}

type Keyboard_KeyboardHotkeyStreamServer interface {
	Send(*KeyboardHotkeyStreamResponse) error
	grpc.ServerStream
}

type keyboardKeyboardHotkeyStreamServer struct {
	grpc.ServerStream
}

func (x *keyboardKeyboardHotkeyStreamServer) Send(m *KeyboardHotkeyStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Keyboard_KeyboardTextStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(KeyboardTextStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KeyboardServer).KeyboardTextStream(m, &keyboardKeyboardTextStreamServer{stream})
}

type Keyboard_KeyboardTextStreamServer interface {
	Send(*KeyboardTextStreamResponse) error
	grpc.ServerStream
}

type keyboardKeyboardTextStreamServer struct {
	grpc.ServerStream
}

func (x *keyboardKeyboardTextStreamServer) Send(m *KeyboardTextStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Keyboard_KeyboardCharacterStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(KeyboardCharacterStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KeyboardServer).KeyboardCharacterStream(m, &keyboardKeyboardCharacterStreamServer{stream})
}

type Keyboard_KeyboardCharacterStreamServer interface {
	Send(*KeyboardCharacterStreamResponse) error
	grpc.ServerStream
}

type keyboardKeyboardCharacterStreamServer struct {
	grpc.ServerStream
}

func (x *keyboardKeyboardCharacterStreamServer) Send(m *KeyboardCharacterStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Keyboard_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Keyboard",
	HandlerType: (*KeyboardServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "KeyboardHotkeyStream",
			Handler:       _Keyboard_KeyboardHotkeyStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "KeyboardTextStream",
			Handler:       _Keyboard_KeyboardTextStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "KeyboardCharacterStream",
			Handler:       _Keyboard_KeyboardCharacterStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "keyboard.proto",
}
