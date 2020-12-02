// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: locker/messaging/protobuf/messageSchema.proto

package protobuf

import (
	proto "github.com/golang/protobuf/proto"
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

type MessageType int32

const (
	MessageType_META    MessageType = 0
	MessageType_PACKAGE MessageType = 1
	MessageType_ACK     MessageType = 2
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "META",
		1: "PACKAGE",
		2: "ACK",
	}
	MessageType_value = map[string]int32{
		"META":    0,
		"PACKAGE": 1,
		"ACK":     2,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_locker_messaging_protobuf_messageSchema_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_locker_messaging_protobuf_messageSchema_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_locker_messaging_protobuf_messageSchema_proto_rawDescGZIP(), []int{0}
}

type FileMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        []byte      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MsgType   MessageType `protobuf:"varint,2,opt,name=msgType,proto3,enum=protobuf.MessageType" json:"msgType,omitempty"`
	Namespace string      `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Project   string      `protobuf:"bytes,4,opt,name=project,proto3" json:"project,omitempty"`
	JobID     string      `protobuf:"bytes,5,opt,name=jobID,proto3" json:"jobID,omitempty"`
	Filename  string      `protobuf:"bytes,6,opt,name=filename,proto3" json:"filename,omitempty"`
	Hash      []byte      `protobuf:"bytes,7,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *FileMeta) Reset() {
	*x = FileMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_locker_messaging_protobuf_messageSchema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMeta) ProtoMessage() {}

func (x *FileMeta) ProtoReflect() protoreflect.Message {
	mi := &file_locker_messaging_protobuf_messageSchema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMeta.ProtoReflect.Descriptor instead.
func (*FileMeta) Descriptor() ([]byte, []int) {
	return file_locker_messaging_protobuf_messageSchema_proto_rawDescGZIP(), []int{0}
}

func (x *FileMeta) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *FileMeta) GetMsgType() MessageType {
	if x != nil {
		return x.MsgType
	}
	return MessageType_META
}

func (x *FileMeta) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *FileMeta) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *FileMeta) GetJobID() string {
	if x != nil {
		return x.JobID
	}
	return ""
}

func (x *FileMeta) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileMeta) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type FilePackage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           []byte      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MsgType      MessageType `protobuf:"varint,2,opt,name=msgType,proto3,enum=protobuf.MessageType" json:"msgType,omitempty"`
	Payload      []byte      `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	IsTerminated bool        `protobuf:"varint,4,opt,name=isTerminated,proto3" json:"isTerminated,omitempty"`
}

func (x *FilePackage) Reset() {
	*x = FilePackage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_locker_messaging_protobuf_messageSchema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilePackage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilePackage) ProtoMessage() {}

func (x *FilePackage) ProtoReflect() protoreflect.Message {
	mi := &file_locker_messaging_protobuf_messageSchema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilePackage.ProtoReflect.Descriptor instead.
func (*FilePackage) Descriptor() ([]byte, []int) {
	return file_locker_messaging_protobuf_messageSchema_proto_rawDescGZIP(), []int{1}
}

func (x *FilePackage) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *FilePackage) GetMsgType() MessageType {
	if x != nil {
		return x.MsgType
	}
	return MessageType_META
}

func (x *FilePackage) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *FilePackage) GetIsTerminated() bool {
	if x != nil {
		return x.IsTerminated
	}
	return false
}

type ServerAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            []byte      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MsgType       MessageType `protobuf:"varint,2,opt,name=msgType,proto3,enum=protobuf.MessageType" json:"msgType,omitempty"`
	ServerSuccess bool        `protobuf:"varint,3,opt,name=serverSuccess,proto3" json:"serverSuccess,omitempty"`
}

func (x *ServerAck) Reset() {
	*x = ServerAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_locker_messaging_protobuf_messageSchema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerAck) ProtoMessage() {}

func (x *ServerAck) ProtoReflect() protoreflect.Message {
	mi := &file_locker_messaging_protobuf_messageSchema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerAck.ProtoReflect.Descriptor instead.
func (*ServerAck) Descriptor() ([]byte, []int) {
	return file_locker_messaging_protobuf_messageSchema_proto_rawDescGZIP(), []int{2}
}

func (x *ServerAck) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ServerAck) GetMsgType() MessageType {
	if x != nil {
		return x.MsgType
	}
	return MessageType_META
}

func (x *ServerAck) GetServerSuccess() bool {
	if x != nil {
		return x.ServerSuccess
	}
	return false
}

type LockerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//	*LockerMessage_Meta
	//	*LockerMessage_Package
	//	*LockerMessage_Ack
	Message isLockerMessage_Message `protobuf_oneof:"message"`
}

func (x *LockerMessage) Reset() {
	*x = LockerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_locker_messaging_protobuf_messageSchema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockerMessage) ProtoMessage() {}

func (x *LockerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_locker_messaging_protobuf_messageSchema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockerMessage.ProtoReflect.Descriptor instead.
func (*LockerMessage) Descriptor() ([]byte, []int) {
	return file_locker_messaging_protobuf_messageSchema_proto_rawDescGZIP(), []int{3}
}

func (m *LockerMessage) GetMessage() isLockerMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *LockerMessage) GetMeta() *FileMeta {
	if x, ok := x.GetMessage().(*LockerMessage_Meta); ok {
		return x.Meta
	}
	return nil
}

func (x *LockerMessage) GetPackage() *FilePackage {
	if x, ok := x.GetMessage().(*LockerMessage_Package); ok {
		return x.Package
	}
	return nil
}

func (x *LockerMessage) GetAck() *ServerAck {
	if x, ok := x.GetMessage().(*LockerMessage_Ack); ok {
		return x.Ack
	}
	return nil
}

type isLockerMessage_Message interface {
	isLockerMessage_Message()
}

type LockerMessage_Meta struct {
	Meta *FileMeta `protobuf:"bytes,1,opt,name=meta,proto3,oneof"`
}

type LockerMessage_Package struct {
	Package *FilePackage `protobuf:"bytes,2,opt,name=package,proto3,oneof"`
}

type LockerMessage_Ack struct {
	Ack *ServerAck `protobuf:"bytes,3,opt,name=ack,proto3,oneof"`
}

func (*LockerMessage_Meta) isLockerMessage_Message() {}

func (*LockerMessage_Package) isLockerMessage_Message() {}

func (*LockerMessage_Ack) isLockerMessage_Message() {}

var File_locker_messaging_protobuf_messageSchema_proto protoreflect.FileDescriptor

var file_locker_messaging_protobuf_messageSchema_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0xc9, 0x01, 0x0a, 0x08, 0x46, 0x69,
	0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07,
	0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6a, 0x6f, 0x62, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x8c, 0x01, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6d,
	0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x64, 0x22, 0x72, 0x0a, 0x09, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x63,
	0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2f, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xa0, 0x01, 0x0a, 0x0d, 0x4c, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x07,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x03, 0x61, 0x63, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x03, 0x61, 0x63, 0x6b,
	0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x2d, 0x0a, 0x0b, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x45,
	0x54, 0x41, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x41, 0x43, 0x4b, 0x41, 0x47, 0x45, 0x10,
	0x01, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x43, 0x4b, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_locker_messaging_protobuf_messageSchema_proto_rawDescOnce sync.Once
	file_locker_messaging_protobuf_messageSchema_proto_rawDescData = file_locker_messaging_protobuf_messageSchema_proto_rawDesc
)

func file_locker_messaging_protobuf_messageSchema_proto_rawDescGZIP() []byte {
	file_locker_messaging_protobuf_messageSchema_proto_rawDescOnce.Do(func() {
		file_locker_messaging_protobuf_messageSchema_proto_rawDescData = protoimpl.X.CompressGZIP(file_locker_messaging_protobuf_messageSchema_proto_rawDescData)
	})
	return file_locker_messaging_protobuf_messageSchema_proto_rawDescData
}

var file_locker_messaging_protobuf_messageSchema_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_locker_messaging_protobuf_messageSchema_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_locker_messaging_protobuf_messageSchema_proto_goTypes = []interface{}{
	(MessageType)(0),      // 0: protobuf.MessageType
	(*FileMeta)(nil),      // 1: protobuf.FileMeta
	(*FilePackage)(nil),   // 2: protobuf.FilePackage
	(*ServerAck)(nil),     // 3: protobuf.ServerAck
	(*LockerMessage)(nil), // 4: protobuf.LockerMessage
}
var file_locker_messaging_protobuf_messageSchema_proto_depIdxs = []int32{
	0, // 0: protobuf.FileMeta.msgType:type_name -> protobuf.MessageType
	0, // 1: protobuf.FilePackage.msgType:type_name -> protobuf.MessageType
	0, // 2: protobuf.ServerAck.msgType:type_name -> protobuf.MessageType
	1, // 3: protobuf.LockerMessage.meta:type_name -> protobuf.FileMeta
	2, // 4: protobuf.LockerMessage.package:type_name -> protobuf.FilePackage
	3, // 5: protobuf.LockerMessage.ack:type_name -> protobuf.ServerAck
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_locker_messaging_protobuf_messageSchema_proto_init() }
func file_locker_messaging_protobuf_messageSchema_proto_init() {
	if File_locker_messaging_protobuf_messageSchema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_locker_messaging_protobuf_messageSchema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMeta); i {
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
		file_locker_messaging_protobuf_messageSchema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilePackage); i {
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
		file_locker_messaging_protobuf_messageSchema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerAck); i {
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
		file_locker_messaging_protobuf_messageSchema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockerMessage); i {
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
	file_locker_messaging_protobuf_messageSchema_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*LockerMessage_Meta)(nil),
		(*LockerMessage_Package)(nil),
		(*LockerMessage_Ack)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_locker_messaging_protobuf_messageSchema_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_locker_messaging_protobuf_messageSchema_proto_goTypes,
		DependencyIndexes: file_locker_messaging_protobuf_messageSchema_proto_depIdxs,
		EnumInfos:         file_locker_messaging_protobuf_messageSchema_proto_enumTypes,
		MessageInfos:      file_locker_messaging_protobuf_messageSchema_proto_msgTypes,
	}.Build()
	File_locker_messaging_protobuf_messageSchema_proto = out.File
	file_locker_messaging_protobuf_messageSchema_proto_rawDesc = nil
	file_locker_messaging_protobuf_messageSchema_proto_goTypes = nil
	file_locker_messaging_protobuf_messageSchema_proto_depIdxs = nil
}
