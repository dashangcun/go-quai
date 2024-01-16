// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: p2p/pb/quai_messages.proto

package pb

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

// GossipSub messages for broadcasting blocks and transactions
type GossipBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block *Block `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *GossipBlock) Reset() {
	*x = GossipBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_pb_quai_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GossipBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GossipBlock) ProtoMessage() {}

func (x *GossipBlock) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_pb_quai_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GossipBlock.ProtoReflect.Descriptor instead.
func (*GossipBlock) Descriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{0}
}

func (x *GossipBlock) GetBlock() *Block {
	if x != nil {
		return x.Block
	}
	return nil
}

type GossipTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *GossipTransaction) Reset() {
	*x = GossipTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_pb_quai_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GossipTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GossipTransaction) ProtoMessage() {}

func (x *GossipTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_pb_quai_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GossipTransaction.ProtoReflect.Descriptor instead.
func (*GossipTransaction) Descriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{1}
}

func (x *GossipTransaction) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

// QuaiRequestMessage is the main 'envelope' for QuaiProtocol request messages
type QuaiRequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *Location `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Hash     *Hash     `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// Types that are assignable to Request:
	//
	//	*QuaiRequestMessage_Block
	//	*QuaiRequestMessage_Header
	//	*QuaiRequestMessage_Transaction
	Request isQuaiRequestMessage_Request `protobuf_oneof:"request"`
}

func (x *QuaiRequestMessage) Reset() {
	*x = QuaiRequestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_pb_quai_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuaiRequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuaiRequestMessage) ProtoMessage() {}

func (x *QuaiRequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_pb_quai_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuaiRequestMessage.ProtoReflect.Descriptor instead.
func (*QuaiRequestMessage) Descriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{2}
}

func (x *QuaiRequestMessage) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *QuaiRequestMessage) GetHash() *Hash {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (m *QuaiRequestMessage) GetRequest() isQuaiRequestMessage_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *QuaiRequestMessage) GetBlock() *Block {
	if x, ok := x.GetRequest().(*QuaiRequestMessage_Block); ok {
		return x.Block
	}
	return nil
}

func (x *QuaiRequestMessage) GetHeader() *Header {
	if x, ok := x.GetRequest().(*QuaiRequestMessage_Header); ok {
		return x.Header
	}
	return nil
}

func (x *QuaiRequestMessage) GetTransaction() *Transaction {
	if x, ok := x.GetRequest().(*QuaiRequestMessage_Transaction); ok {
		return x.Transaction
	}
	return nil
}

type isQuaiRequestMessage_Request interface {
	isQuaiRequestMessage_Request()
}

type QuaiRequestMessage_Block struct {
	Block *Block `protobuf:"bytes,3,opt,name=block,proto3,oneof"`
}

type QuaiRequestMessage_Header struct {
	Header *Header `protobuf:"bytes,4,opt,name=header,proto3,oneof"`
}

type QuaiRequestMessage_Transaction struct {
	Transaction *Transaction `protobuf:"bytes,5,opt,name=transaction,proto3,oneof"`
}

func (*QuaiRequestMessage_Block) isQuaiRequestMessage_Request() {}

func (*QuaiRequestMessage_Header) isQuaiRequestMessage_Request() {}

func (*QuaiRequestMessage_Transaction) isQuaiRequestMessage_Request() {}

// QuaiResponseMessage is the main 'envelope' for QuaiProtocol response messages
type QuaiResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//
	//	*QuaiResponseMessage_Block
	//	*QuaiResponseMessage_Header
	//	*QuaiResponseMessage_Transaction
	Response isQuaiResponseMessage_Response `protobuf_oneof:"response"`
}

func (x *QuaiResponseMessage) Reset() {
	*x = QuaiResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_pb_quai_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuaiResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuaiResponseMessage) ProtoMessage() {}

func (x *QuaiResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_pb_quai_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuaiResponseMessage.ProtoReflect.Descriptor instead.
func (*QuaiResponseMessage) Descriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{3}
}

func (m *QuaiResponseMessage) GetResponse() isQuaiResponseMessage_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *QuaiResponseMessage) GetBlock() *Block {
	if x, ok := x.GetResponse().(*QuaiResponseMessage_Block); ok {
		return x.Block
	}
	return nil
}

func (x *QuaiResponseMessage) GetHeader() *Header {
	if x, ok := x.GetResponse().(*QuaiResponseMessage_Header); ok {
		return x.Header
	}
	return nil
}

func (x *QuaiResponseMessage) GetTransaction() *Transaction {
	if x, ok := x.GetResponse().(*QuaiResponseMessage_Transaction); ok {
		return x.Transaction
	}
	return nil
}

type isQuaiResponseMessage_Response interface {
	isQuaiResponseMessage_Response()
}

type QuaiResponseMessage_Block struct {
	Block *Block `protobuf:"bytes,1,opt,name=block,proto3,oneof"`
}

type QuaiResponseMessage_Header struct {
	Header *Header `protobuf:"bytes,2,opt,name=header,proto3,oneof"`
}

type QuaiResponseMessage_Transaction struct {
	Transaction *Transaction `protobuf:"bytes,3,opt,name=transaction,proto3,oneof"`
}

func (*QuaiResponseMessage_Block) isQuaiResponseMessage_Response() {}

func (*QuaiResponseMessage_Header) isQuaiResponseMessage_Response() {}

func (*QuaiResponseMessage_Transaction) isQuaiResponseMessage_Response() {}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location []byte `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_pb_quai_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_pb_quai_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{4}
}

func (x *Location) GetLocation() []byte {
	if x != nil {
		return x.Location
	}
	return nil
}

// Hash structure
type Hash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"` // Hash is an array of 32 bytes
}

func (x *Hash) Reset() {
	*x = Hash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_pb_quai_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hash) ProtoMessage() {}

func (x *Hash) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_pb_quai_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hash.ProtoReflect.Descriptor instead.
func (*Hash) Descriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{5}
}

func (x *Hash) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

// Block structure
type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Txs    []*Transaction `protobuf:"bytes,2,rep,name=txs,proto3" json:"txs,omitempty"`
	Uncles []*Header      `protobuf:"bytes,3,rep,name=uncles,proto3" json:"uncles,omitempty"`
	Etxs   []*Transaction `protobuf:"bytes,4,rep,name=etxs,proto3" json:"etxs,omitempty"`
	// TODO: add submanifest
	// TODO: add receipts
	// TODO: add hasher
	NodeCtx uint32 `protobuf:"varint,7,opt,name=nodeCtx,proto3" json:"nodeCtx,omitempty"` // Additional fields...
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_pb_quai_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_pb_quai_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{6}
}

func (x *Block) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Block) GetTxs() []*Transaction {
	if x != nil {
		return x.Txs
	}
	return nil
}

func (x *Block) GetUncles() []*Header {
	if x != nil {
		return x.Uncles
	}
	return nil
}

func (x *Block) GetEtxs() []*Transaction {
	if x != nil {
		return x.Etxs
	}
	return nil
}

func (x *Block) GetNodeCtx() uint32 {
	if x != nil {
		return x.NodeCtx
	}
	return 0
}

// Header structure
type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash       *Hash  `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	ParentHash *Hash  `protobuf:"bytes,2,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	GasLimit   uint64 `protobuf:"varint,3,opt,name=gasLimit,proto3" json:"gasLimit,omitempty"`
	GasUsed    uint64 `protobuf:"varint,4,opt,name=gasUsed,proto3" json:"gasUsed,omitempty"` // Additional fields...
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_pb_quai_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_pb_quai_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{7}
}

func (x *Header) GetHash() *Hash {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *Header) GetParentHash() *Hash {
	if x != nil {
		return x.ParentHash
	}
	return nil
}

func (x *Header) GetGasLimit() uint64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

func (x *Header) GetGasUsed() uint64 {
	if x != nil {
		return x.GasUsed
	}
	return 0
}

// Transaction structure
type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash     *Hash  `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	From     string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To       string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Nonce    uint64 `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Value    string `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	GasPrice string `protobuf:"bytes,6,opt,name=gasPrice,proto3" json:"gasPrice,omitempty"`
	Gas      uint64 `protobuf:"varint,7,opt,name=gas,proto3" json:"gas,omitempty"`
	Input    []byte `protobuf:"bytes,8,opt,name=input,proto3" json:"input,omitempty"` // Additional fields...
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_pb_quai_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_pb_quai_messages_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_p2p_pb_quai_messages_proto_rawDescGZIP(), []int{8}
}

func (x *Transaction) GetHash() *Hash {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *Transaction) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Transaction) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Transaction) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *Transaction) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Transaction) GetGasPrice() string {
	if x != nil {
		return x.GasPrice
	}
	return ""
}

func (x *Transaction) GetGas() uint64 {
	if x != nil {
		return x.Gas
	}
	return 0
}

func (x *Transaction) GetInput() []byte {
	if x != nil {
		return x.Input
	}
	return nil
}

var File_p2p_pb_quai_messages_proto protoreflect.FileDescriptor

var file_p2p_pb_quai_messages_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x32, 0x70, 0x2f, 0x70, 0x62, 0x2f, 0x71, 0x75, 0x61, 0x69, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x71, 0x75,
	0x61, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x38, 0x0a, 0x0b, 0x47, 0x6f,
	0x73, 0x73, 0x69, 0x70, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x29, 0x0a, 0x05, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x71, 0x75, 0x61, 0x69, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x50, 0x0a, 0x11, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x71, 0x75, 0x61, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x97, 0x02, 0x0a, 0x12, 0x51, 0x75, 0x61, 0x69, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x71, 0x75, 0x61, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x26, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x71, 0x75, 0x61, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x48,
	0x61, 0x73, 0x68, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x2b, 0x0a, 0x05, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x71, 0x75, 0x61, 0x69, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x00, 0x52,
	0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2e, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x71, 0x75, 0x61, 0x69, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x71, 0x75,
	0x61, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0xbd, 0x01, 0x0a, 0x13, 0x51, 0x75, 0x61, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x71, 0x75, 0x61, 0x69, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x05,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2e, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x71, 0x75, 0x61, 0x69, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x71, 0x75, 0x61,
	0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x26, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1a, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x22, 0xd9, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2c,
	0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x71, 0x75, 0x61, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x03,
	0x74, 0x78, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x71, 0x75, 0x61, 0x69,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x74, 0x78, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x75, 0x6e, 0x63,
	0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x71, 0x75, 0x61, 0x69,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x75, 0x6e, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x04, 0x65, 0x74, 0x78, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x71, 0x75, 0x61, 0x69, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x04, 0x65, 0x74, 0x78, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x74,
	0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x74, 0x78,
	0x22, 0x9a, 0x01, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x71, 0x75, 0x61, 0x69,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x52, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x12, 0x32, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x71, 0x75, 0x61, 0x69, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x52, 0x0a, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x67, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x22, 0xc9, 0x01,
	0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x71, 0x75,
	0x61, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x52,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x61, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x67, 0x61, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x6e, 0x74,
	0x2d, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x71,
	0x75, 0x61, 0x69, 0x2f, 0x70, 0x32, 0x70, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_p2p_pb_quai_messages_proto_rawDescOnce sync.Once
	file_p2p_pb_quai_messages_proto_rawDescData = file_p2p_pb_quai_messages_proto_rawDesc
)

func file_p2p_pb_quai_messages_proto_rawDescGZIP() []byte {
	file_p2p_pb_quai_messages_proto_rawDescOnce.Do(func() {
		file_p2p_pb_quai_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_p2p_pb_quai_messages_proto_rawDescData)
	})
	return file_p2p_pb_quai_messages_proto_rawDescData
}

var file_p2p_pb_quai_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_p2p_pb_quai_messages_proto_goTypes = []interface{}{
	(*GossipBlock)(nil),         // 0: quaiprotocol.GossipBlock
	(*GossipTransaction)(nil),   // 1: quaiprotocol.GossipTransaction
	(*QuaiRequestMessage)(nil),  // 2: quaiprotocol.QuaiRequestMessage
	(*QuaiResponseMessage)(nil), // 3: quaiprotocol.QuaiResponseMessage
	(*Location)(nil),            // 4: quaiprotocol.Location
	(*Hash)(nil),                // 5: quaiprotocol.Hash
	(*Block)(nil),               // 6: quaiprotocol.Block
	(*Header)(nil),              // 7: quaiprotocol.Header
	(*Transaction)(nil),         // 8: quaiprotocol.Transaction
}
var file_p2p_pb_quai_messages_proto_depIdxs = []int32{
	6,  // 0: quaiprotocol.GossipBlock.block:type_name -> quaiprotocol.Block
	8,  // 1: quaiprotocol.GossipTransaction.transaction:type_name -> quaiprotocol.Transaction
	4,  // 2: quaiprotocol.QuaiRequestMessage.location:type_name -> quaiprotocol.Location
	5,  // 3: quaiprotocol.QuaiRequestMessage.hash:type_name -> quaiprotocol.Hash
	6,  // 4: quaiprotocol.QuaiRequestMessage.block:type_name -> quaiprotocol.Block
	7,  // 5: quaiprotocol.QuaiRequestMessage.header:type_name -> quaiprotocol.Header
	8,  // 6: quaiprotocol.QuaiRequestMessage.transaction:type_name -> quaiprotocol.Transaction
	6,  // 7: quaiprotocol.QuaiResponseMessage.block:type_name -> quaiprotocol.Block
	7,  // 8: quaiprotocol.QuaiResponseMessage.header:type_name -> quaiprotocol.Header
	8,  // 9: quaiprotocol.QuaiResponseMessage.transaction:type_name -> quaiprotocol.Transaction
	7,  // 10: quaiprotocol.Block.header:type_name -> quaiprotocol.Header
	8,  // 11: quaiprotocol.Block.txs:type_name -> quaiprotocol.Transaction
	7,  // 12: quaiprotocol.Block.uncles:type_name -> quaiprotocol.Header
	8,  // 13: quaiprotocol.Block.etxs:type_name -> quaiprotocol.Transaction
	5,  // 14: quaiprotocol.Header.hash:type_name -> quaiprotocol.Hash
	5,  // 15: quaiprotocol.Header.parentHash:type_name -> quaiprotocol.Hash
	5,  // 16: quaiprotocol.Transaction.hash:type_name -> quaiprotocol.Hash
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_p2p_pb_quai_messages_proto_init() }
func file_p2p_pb_quai_messages_proto_init() {
	if File_p2p_pb_quai_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_p2p_pb_quai_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GossipBlock); i {
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
		file_p2p_pb_quai_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GossipTransaction); i {
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
		file_p2p_pb_quai_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuaiRequestMessage); i {
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
		file_p2p_pb_quai_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuaiResponseMessage); i {
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
		file_p2p_pb_quai_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_p2p_pb_quai_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hash); i {
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
		file_p2p_pb_quai_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_p2p_pb_quai_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_p2p_pb_quai_messages_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
	file_p2p_pb_quai_messages_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*QuaiRequestMessage_Block)(nil),
		(*QuaiRequestMessage_Header)(nil),
		(*QuaiRequestMessage_Transaction)(nil),
	}
	file_p2p_pb_quai_messages_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*QuaiResponseMessage_Block)(nil),
		(*QuaiResponseMessage_Header)(nil),
		(*QuaiResponseMessage_Transaction)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_p2p_pb_quai_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_p2p_pb_quai_messages_proto_goTypes,
		DependencyIndexes: file_p2p_pb_quai_messages_proto_depIdxs,
		MessageInfos:      file_p2p_pb_quai_messages_proto_msgTypes,
	}.Build()
	File_p2p_pb_quai_messages_proto = out.File
	file_p2p_pb_quai_messages_proto_rawDesc = nil
	file_p2p_pb_quai_messages_proto_goTypes = nil
	file_p2p_pb_quai_messages_proto_depIdxs = nil
}
