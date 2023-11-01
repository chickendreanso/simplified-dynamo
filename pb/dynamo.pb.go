// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: dynamo.proto

package dynamo

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_dynamo_proto_msgTypes[0]
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
	return file_dynamo_proto_rawDescGZIP(), []int{0}
}

type KeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key         string       `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value       string       `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	VectorClock *VectorClock `protobuf:"bytes,3,opt,name=vector_clock,json=vectorClock,proto3" json:"vector_clock,omitempty"`
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_dynamo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValue.ProtoReflect.Descriptor instead.
func (*KeyValue) Descriptor() ([]byte, []int) {
	return file_dynamo_proto_rawDescGZIP(), []int{1}
}

func (x *KeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *KeyValue) GetVectorClock() *VectorClock {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

type ClockStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClokcVal  int64                  `protobuf:"varint,1,opt,name=clokcVal,proto3" json:"clokcVal,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ClockStruct) Reset() {
	*x = ClockStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClockStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClockStruct) ProtoMessage() {}

func (x *ClockStruct) ProtoReflect() protoreflect.Message {
	mi := &file_dynamo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClockStruct.ProtoReflect.Descriptor instead.
func (*ClockStruct) Descriptor() ([]byte, []int) {
	return file_dynamo_proto_rawDescGZIP(), []int{2}
}

func (x *ClockStruct) GetClokcVal() int64 {
	if x != nil {
		return x.ClokcVal
	}
	return 0
}

func (x *ClockStruct) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type VectorClock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamps map[uint32]*ClockStruct `protobuf:"bytes,1,rep,name=timestamps,proto3" json:"timestamps,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Map to hold the vector clock data
}

func (x *VectorClock) Reset() {
	*x = VectorClock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VectorClock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VectorClock) ProtoMessage() {}

func (x *VectorClock) ProtoReflect() protoreflect.Message {
	mi := &file_dynamo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VectorClock.ProtoReflect.Descriptor instead.
func (*VectorClock) Descriptor() ([]byte, []int) {
	return file_dynamo_proto_rawDescGZIP(), []int{3}
}

func (x *VectorClock) GetTimestamps() map[uint32]*ClockStruct {
	if x != nil {
		return x.Timestamps
	}
	return nil
}

type BulkWriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyValue []*KeyValue `protobuf:"bytes,1,rep,name=keyValue,proto3" json:"keyValue,omitempty"`
}

func (x *BulkWriteRequest) Reset() {
	*x = BulkWriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkWriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkWriteRequest) ProtoMessage() {}

func (x *BulkWriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dynamo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkWriteRequest.ProtoReflect.Descriptor instead.
func (*BulkWriteRequest) Descriptor() ([]byte, []int) {
	return file_dynamo_proto_rawDescGZIP(), []int{4}
}

func (x *BulkWriteRequest) GetKeyValue() []*KeyValue {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

type HintedHandoffWriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyValue *KeyValue `protobuf:"bytes,1,opt,name=keyValue,proto3" json:"keyValue,omitempty"`
	Node     *Node     `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *HintedHandoffWriteRequest) Reset() {
	*x = HintedHandoffWriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HintedHandoffWriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HintedHandoffWriteRequest) ProtoMessage() {}

func (x *HintedHandoffWriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dynamo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HintedHandoffWriteRequest.ProtoReflect.Descriptor instead.
func (*HintedHandoffWriteRequest) Descriptor() ([]byte, []int) {
	return file_dynamo_proto_rawDescGZIP(), []int{5}
}

func (x *HintedHandoffWriteRequest) GetKeyValue() *KeyValue {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

func (x *HintedHandoffWriteRequest) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyValue  *KeyValue `protobuf:"bytes,1,opt,name=keyValue,proto3" json:"keyValue,omitempty"`
	IsReplica bool      `protobuf:"varint,2,opt,name=isReplica,proto3" json:"isReplica,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dynamo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_dynamo_proto_rawDescGZIP(), []int{6}
}

func (x *WriteRequest) GetKeyValue() *KeyValue {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

func (x *WriteRequest) GetIsReplica() bool {
	if x != nil {
		return x.IsReplica
	}
	return false
}

type WriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyValue []*KeyValue `protobuf:"bytes,1,rep,name=keyValue,proto3" json:"keyValue,omitempty"`
	Success  bool        `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Message  string      `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *WriteResponse) Reset() {
	*x = WriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteResponse) ProtoMessage() {}

func (x *WriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dynamo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteResponse.ProtoReflect.Descriptor instead.
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return file_dynamo_proto_rawDescGZIP(), []int{7}
}

func (x *WriteResponse) GetKeyValue() []*KeyValue {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

func (x *WriteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *WriteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key       string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	IsReplica bool   `protobuf:"varint,2,opt,name=isReplica,proto3" json:"isReplica,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dynamo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_dynamo_proto_rawDescGZIP(), []int{8}
}

func (x *ReadRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ReadRequest) GetIsReplica() bool {
	if x != nil {
		return x.IsReplica
	}
	return false
}

type ReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyValue []*KeyValue `protobuf:"bytes,1,rep,name=keyValue,proto3" json:"keyValue,omitempty"`
	Success  bool        `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Message  string      `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamo_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dynamo_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_dynamo_proto_rawDescGZIP(), []int{9}
}

func (x *ReadResponse) GetKeyValue() []*KeyValue {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

func (x *ReadResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ReadResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GossipMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"` // "repeated" for lists of items
}

func (x *GossipMessage) Reset() {
	*x = GossipMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamo_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GossipMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GossipMessage) ProtoMessage() {}

func (x *GossipMessage) ProtoReflect() protoreflect.Message {
	mi := &file_dynamo_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GossipMessage.ProtoReflect.Descriptor instead.
func (*GossipMessage) Descriptor() ([]byte, []int) {
	return file_dynamo_proto_rawDescGZIP(), []int{10}
}

func (x *GossipMessage) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type GossipAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success        bool            `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message        string          `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	MembershipList *MembershipList `protobuf:"bytes,3,opt,name=membershipList,proto3" json:"membershipList,omitempty"`
}

func (x *GossipAck) Reset() {
	*x = GossipAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamo_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GossipAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GossipAck) ProtoMessage() {}

func (x *GossipAck) ProtoReflect() protoreflect.Message {
	mi := &file_dynamo_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GossipAck.ProtoReflect.Descriptor instead.
func (*GossipAck) Descriptor() ([]byte, []int) {
	return file_dynamo_proto_rawDescGZIP(), []int{11}
}

func (x *GossipAck) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GossipAck) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GossipAck) GetMembershipList() *MembershipList {
	if x != nil {
		return x.MembershipList
	}
	return nil
}

type MembershipList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *MembershipList) Reset() {
	*x = MembershipList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamo_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MembershipList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MembershipList) ProtoMessage() {}

func (x *MembershipList) ProtoReflect() protoreflect.Message {
	mi := &file_dynamo_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MembershipList.ProtoReflect.Descriptor instead.
func (*MembershipList) Descriptor() ([]byte, []int) {
	return file_dynamo_proto_rawDescGZIP(), []int{12}
}

func (x *MembershipList) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start   uint32 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Port    uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynamo_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_dynamo_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_dynamo_proto_rawDescGZIP(), []int{13}
}

func (x *Node) GetStart() uint32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Node) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Node) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_dynamo_proto protoreflect.FileDescriptor

var file_dynamo_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x6a, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x36, 0x0a, 0x0c, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x79, 0x6e,
	0x61, 0x6d, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x63, 0x0a, 0x0b,
	0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x6c, 0x6f, 0x6b, 0x63, 0x56, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63,
	0x6c, 0x6f, 0x6b, 0x63, 0x56, 0x61, 0x6c, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0xa6, 0x01, 0x0a, 0x0b, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x43, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2e, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x73, 0x1a, 0x52, 0x0a, 0x0f, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x79, 0x6e,
	0x61, 0x6d, 0x6f, 0x2e, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x40, 0x0a, 0x10, 0x42, 0x75,
	0x6c, 0x6b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c,
	0x0a, 0x08, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x6b, 0x0a, 0x19,
	0x48, 0x69, 0x6e, 0x74, 0x65, 0x64, 0x48, 0x61, 0x6e, 0x64, 0x6f, 0x66, 0x66, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x6b, 0x65, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x79,
	0x6e, 0x61, 0x6d, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6b,
	0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x5a, 0x0a, 0x0c, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x6b, 0x65, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x79,
	0x6e, 0x61, 0x6d, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6b,
	0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x22, 0x71, 0x0a, 0x0d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3d, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x22, 0x70, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6b, 0x65, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x33, 0x0a, 0x0d, 0x47, 0x6f, 0x73,
	0x73, 0x69, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x7f,
	0x0a, 0x09, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x41, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x3e, 0x0a, 0x0e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f,
	0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x0e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x34, 0x0a, 0x0e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x22, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x4a, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x32, 0x97, 0x03, 0x0a, 0x0d, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x64,
	0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2e, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x04, 0x52,
	0x65, 0x61, 0x64, 0x12, 0x13, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x34, 0x0a, 0x06, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x12, 0x15, 0x2e, 0x64, 0x79, 0x6e,
	0x61, 0x6d, 0x6f, 0x2e, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x11, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2e, 0x47, 0x6f, 0x73, 0x73, 0x69,
	0x70, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0c, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x1a, 0x11, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2e, 0x47, 0x6f, 0x73, 0x73,
	0x69, 0x70, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0c, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x1a, 0x0d, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0d, 0x48, 0x69, 0x6e, 0x74, 0x65, 0x64, 0x48, 0x61,
	0x6e, 0x64, 0x6f, 0x66, 0x66, 0x12, 0x21, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2e, 0x48,
	0x69, 0x6e, 0x74, 0x65, 0x64, 0x48, 0x61, 0x6e, 0x64, 0x6f, 0x66, 0x66, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0b, 0x53, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x12, 0x18, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x6f, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2e, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x64,
	0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dynamo_proto_rawDescOnce sync.Once
	file_dynamo_proto_rawDescData = file_dynamo_proto_rawDesc
)

func file_dynamo_proto_rawDescGZIP() []byte {
	file_dynamo_proto_rawDescOnce.Do(func() {
		file_dynamo_proto_rawDescData = protoimpl.X.CompressGZIP(file_dynamo_proto_rawDescData)
	})
	return file_dynamo_proto_rawDescData
}

var file_dynamo_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_dynamo_proto_goTypes = []interface{}{
	(*Empty)(nil),                     // 0: dynamo.Empty
	(*KeyValue)(nil),                  // 1: dynamo.KeyValue
	(*ClockStruct)(nil),               // 2: dynamo.ClockStruct
	(*VectorClock)(nil),               // 3: dynamo.VectorClock
	(*BulkWriteRequest)(nil),          // 4: dynamo.BulkWriteRequest
	(*HintedHandoffWriteRequest)(nil), // 5: dynamo.HintedHandoffWriteRequest
	(*WriteRequest)(nil),              // 6: dynamo.WriteRequest
	(*WriteResponse)(nil),             // 7: dynamo.WriteResponse
	(*ReadRequest)(nil),               // 8: dynamo.ReadRequest
	(*ReadResponse)(nil),              // 9: dynamo.ReadResponse
	(*GossipMessage)(nil),             // 10: dynamo.GossipMessage
	(*GossipAck)(nil),                 // 11: dynamo.GossipAck
	(*MembershipList)(nil),            // 12: dynamo.MembershipList
	(*Node)(nil),                      // 13: dynamo.Node
	nil,                               // 14: dynamo.VectorClock.TimestampsEntry
	(*timestamppb.Timestamp)(nil),     // 15: google.protobuf.Timestamp
}
var file_dynamo_proto_depIdxs = []int32{
	3,  // 0: dynamo.KeyValue.vector_clock:type_name -> dynamo.VectorClock
	15, // 1: dynamo.ClockStruct.timestamp:type_name -> google.protobuf.Timestamp
	14, // 2: dynamo.VectorClock.timestamps:type_name -> dynamo.VectorClock.TimestampsEntry
	1,  // 3: dynamo.BulkWriteRequest.keyValue:type_name -> dynamo.KeyValue
	1,  // 4: dynamo.HintedHandoffWriteRequest.keyValue:type_name -> dynamo.KeyValue
	13, // 5: dynamo.HintedHandoffWriteRequest.node:type_name -> dynamo.Node
	1,  // 6: dynamo.WriteRequest.keyValue:type_name -> dynamo.KeyValue
	1,  // 7: dynamo.WriteResponse.keyValue:type_name -> dynamo.KeyValue
	1,  // 8: dynamo.ReadResponse.keyValue:type_name -> dynamo.KeyValue
	13, // 9: dynamo.GossipMessage.nodes:type_name -> dynamo.Node
	12, // 10: dynamo.GossipAck.membershipList:type_name -> dynamo.MembershipList
	13, // 11: dynamo.MembershipList.nodes:type_name -> dynamo.Node
	2,  // 12: dynamo.VectorClock.TimestampsEntry.value:type_name -> dynamo.ClockStruct
	6,  // 13: dynamo.KeyValueStore.Write:input_type -> dynamo.WriteRequest
	8,  // 14: dynamo.KeyValueStore.Read:input_type -> dynamo.ReadRequest
	10, // 15: dynamo.KeyValueStore.Gossip:input_type -> dynamo.GossipMessage
	13, // 16: dynamo.KeyValueStore.AddNewNode:input_type -> dynamo.Node
	13, // 17: dynamo.KeyValueStore.RemoveNode:input_type -> dynamo.Node
	5,  // 18: dynamo.KeyValueStore.HintedHandoff:input_type -> dynamo.HintedHandoffWriteRequest
	4,  // 19: dynamo.KeyValueStore.SendReplica:input_type -> dynamo.BulkWriteRequest
	7,  // 20: dynamo.KeyValueStore.Write:output_type -> dynamo.WriteResponse
	9,  // 21: dynamo.KeyValueStore.Read:output_type -> dynamo.ReadResponse
	11, // 22: dynamo.KeyValueStore.Gossip:output_type -> dynamo.GossipAck
	11, // 23: dynamo.KeyValueStore.AddNewNode:output_type -> dynamo.GossipAck
	0,  // 24: dynamo.KeyValueStore.RemoveNode:output_type -> dynamo.Empty
	0,  // 25: dynamo.KeyValueStore.HintedHandoff:output_type -> dynamo.Empty
	7,  // 26: dynamo.KeyValueStore.SendReplica:output_type -> dynamo.WriteResponse
	20, // [20:27] is the sub-list for method output_type
	13, // [13:20] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_dynamo_proto_init() }
func file_dynamo_proto_init() {
	if File_dynamo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dynamo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_dynamo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValue); i {
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
		file_dynamo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClockStruct); i {
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
		file_dynamo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VectorClock); i {
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
		file_dynamo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkWriteRequest); i {
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
		file_dynamo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HintedHandoffWriteRequest); i {
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
		file_dynamo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRequest); i {
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
		file_dynamo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteResponse); i {
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
		file_dynamo_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_dynamo_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadResponse); i {
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
		file_dynamo_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GossipMessage); i {
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
		file_dynamo_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GossipAck); i {
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
		file_dynamo_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MembershipList); i {
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
		file_dynamo_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
			RawDescriptor: file_dynamo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dynamo_proto_goTypes,
		DependencyIndexes: file_dynamo_proto_depIdxs,
		MessageInfos:      file_dynamo_proto_msgTypes,
	}.Build()
	File_dynamo_proto = out.File
	file_dynamo_proto_rawDesc = nil
	file_dynamo_proto_goTypes = nil
	file_dynamo_proto_depIdxs = nil
}
