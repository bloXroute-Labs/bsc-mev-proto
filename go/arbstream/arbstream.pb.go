// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: arbstream/protobuf/arbstream.proto

package arbstream

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

type Order struct {
	state           protoimpl.MessageState  `protogen:"open.v1"`
	Transactions    []*Transaction          `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	State           map[string]*StateChange `protobuf:"bytes,2,rep,name=state,proto3" json:"state,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"` // State changes
	IsArbTx         bool                    `protobuf:"varint,3,opt,name=isArbTx,proto3" json:"isArbTx,omitempty"`
	BackrunmeConfig []string                `protobuf:"bytes,4,rep,name=backrunmeConfig,proto3" json:"backrunmeConfig,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_arbstream_protobuf_arbstream_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_arbstream_protobuf_arbstream_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_arbstream_protobuf_arbstream_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *Order) GetState() map[string]*StateChange {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *Order) GetIsArbTx() bool {
	if x != nil {
		return x.IsArbTx
	}
	return false
}

func (x *Order) GetBackrunmeConfig() []string {
	if x != nil {
		return x.BackrunmeConfig
	}
	return nil
}

type StateChange struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Slots         map[string]string      `protobuf:"bytes,1,rep,name=slots,proto3" json:"slots,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StateChange) Reset() {
	*x = StateChange{}
	mi := &file_arbstream_protobuf_arbstream_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StateChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateChange) ProtoMessage() {}

func (x *StateChange) ProtoReflect() protoreflect.Message {
	mi := &file_arbstream_protobuf_arbstream_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateChange.ProtoReflect.Descriptor instead.
func (*StateChange) Descriptor() ([]byte, []int) {
	return file_arbstream_protobuf_arbstream_proto_rawDescGZIP(), []int{1}
}

func (x *StateChange) GetSlots() map[string]string {
	if x != nil {
		return x.Slots
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TxHash        string                 `protobuf:"bytes,1,opt,name=txHash,proto3" json:"txHash,omitempty"`
	TxContents    *TxContents            `protobuf:"bytes,2,opt,name=txContents,proto3" json:"txContents,omitempty"`
	LocalRegion   bool                   `protobuf:"varint,3,opt,name=localRegion,proto3" json:"localRegion,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	mi := &file_arbstream_protobuf_arbstream_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_arbstream_protobuf_arbstream_proto_msgTypes[2]
	if x != nil {
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
	return file_arbstream_protobuf_arbstream_proto_rawDescGZIP(), []int{2}
}

func (x *Transaction) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

func (x *Transaction) GetTxContents() *TxContents {
	if x != nil {
		return x.TxContents
	}
	return nil
}

func (x *Transaction) GetLocalRegion() bool {
	if x != nil {
		return x.LocalRegion
	}
	return false
}

type TxContents struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	From             string                 `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Gas              string                 `protobuf:"bytes,2,opt,name=gas,proto3" json:"gas,omitempty"`
	GasPrice         string                 `protobuf:"bytes,3,opt,name=gasPrice,proto3" json:"gasPrice,omitempty"`
	Hash             string                 `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
	Input            string                 `protobuf:"bytes,5,opt,name=input,proto3" json:"input,omitempty"`
	Nonce            string                 `protobuf:"bytes,6,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Value            string                 `protobuf:"bytes,7,opt,name=value,proto3" json:"value,omitempty"`
	Type             string                 `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	To               string                 `protobuf:"bytes,9,opt,name=to,proto3" json:"to,omitempty"`
	FunctionSelector string                 `protobuf:"bytes,10,opt,name=functionSelector,proto3" json:"functionSelector,omitempty"`
	Logs             []*Log                 `protobuf:"bytes,11,rep,name=logs,proto3" json:"logs,omitempty"`
	GasTipCap        string                 `protobuf:"bytes,12,opt,name=gasTipCap,proto3" json:"gasTipCap,omitempty"`
	GasFeeCap        string                 `protobuf:"bytes,13,opt,name=gasFeeCap,proto3" json:"gasFeeCap,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *TxContents) Reset() {
	*x = TxContents{}
	mi := &file_arbstream_protobuf_arbstream_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TxContents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxContents) ProtoMessage() {}

func (x *TxContents) ProtoReflect() protoreflect.Message {
	mi := &file_arbstream_protobuf_arbstream_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxContents.ProtoReflect.Descriptor instead.
func (*TxContents) Descriptor() ([]byte, []int) {
	return file_arbstream_protobuf_arbstream_proto_rawDescGZIP(), []int{3}
}

func (x *TxContents) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *TxContents) GetGas() string {
	if x != nil {
		return x.Gas
	}
	return ""
}

func (x *TxContents) GetGasPrice() string {
	if x != nil {
		return x.GasPrice
	}
	return ""
}

func (x *TxContents) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *TxContents) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *TxContents) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *TxContents) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *TxContents) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TxContents) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *TxContents) GetFunctionSelector() string {
	if x != nil {
		return x.FunctionSelector
	}
	return ""
}

func (x *TxContents) GetLogs() []*Log {
	if x != nil {
		return x.Logs
	}
	return nil
}

func (x *TxContents) GetGasTipCap() string {
	if x != nil {
		return x.GasTipCap
	}
	return ""
}

func (x *TxContents) GetGasFeeCap() string {
	if x != nil {
		return x.GasFeeCap
	}
	return ""
}

type Log struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Topics        []string               `protobuf:"bytes,2,rep,name=topics,proto3" json:"topics,omitempty"`
	Data          string                 `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Log) Reset() {
	*x = Log{}
	mi := &file_arbstream_protobuf_arbstream_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_arbstream_protobuf_arbstream_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_arbstream_protobuf_arbstream_proto_rawDescGZIP(), []int{4}
}

func (x *Log) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Log) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

func (x *Log) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type StreamRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	RequestData     string                 `protobuf:"bytes,1,opt,name=request_data,json=requestData,proto3" json:"request_data,omitempty"` // defaults to "all"
	IncludeNonArbTx bool                   `protobuf:"varint,2,opt,name=includeNonArbTx,proto3" json:"includeNonArbTx,omitempty"`           //defaults to false
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *StreamRequest) Reset() {
	*x = StreamRequest{}
	mi := &file_arbstream_protobuf_arbstream_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequest) ProtoMessage() {}

func (x *StreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_arbstream_protobuf_arbstream_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRequest.ProtoReflect.Descriptor instead.
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return file_arbstream_protobuf_arbstream_proto_rawDescGZIP(), []int{5}
}

func (x *StreamRequest) GetRequestData() string {
	if x != nil {
		return x.RequestData
	}
	return ""
}

func (x *StreamRequest) GetIncludeNonArbTx() bool {
	if x != nil {
		return x.IncludeNonArbTx
	}
	return false
}

var File_arbstream_protobuf_arbstream_proto protoreflect.FileDescriptor

var file_arbstream_protobuf_arbstream_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x72, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x72, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x72, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22,
	0x8c, 0x02, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x0c, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x61, 0x72, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x72, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x41, 0x72,
	0x62, 0x54, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x72, 0x62,
	0x54, 0x78, 0x12, 0x28, 0x0a, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x61, 0x63,
	0x6b, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x50, 0x0a, 0x0a,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x72,
	0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x80,
	0x01, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x37,
	0x0a, 0x05, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x61, 0x72, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x53, 0x6c, 0x6f, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x7e, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x35, 0x0a, 0x0a, 0x74, 0x78, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61,
	0x72, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x54, 0x78, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x0a, 0x74, 0x78, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x22, 0xd4, 0x02, 0x0a, 0x0a, 0x54, 0x78, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x67, 0x61, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x74, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x2a, 0x0a, 0x10,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73,
	0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x72, 0x62, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x67, 0x61, 0x73, 0x54, 0x69, 0x70, 0x43, 0x61, 0x70, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x67, 0x61, 0x73, 0x54, 0x69, 0x70, 0x43, 0x61, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x61,
	0x73, 0x46, 0x65, 0x65, 0x43, 0x61, 0x70, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67,
	0x61, 0x73, 0x46, 0x65, 0x65, 0x43, 0x61, 0x70, 0x22, 0x4b, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5c, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x4e, 0x6f, 0x6e, 0x41, 0x72, 0x62, 0x54, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x4e, 0x6f, 0x6e, 0x41, 0x72,
	0x62, 0x54, 0x78, 0x32, 0x50, 0x0a, 0x0f, 0x41, 0x72, 0x62, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x41, 0x72, 0x62, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x2e, 0x61, 0x72, 0x62, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x61, 0x72, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x30, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x72, 0x62, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_arbstream_protobuf_arbstream_proto_rawDescOnce sync.Once
	file_arbstream_protobuf_arbstream_proto_rawDescData = file_arbstream_protobuf_arbstream_proto_rawDesc
)

func file_arbstream_protobuf_arbstream_proto_rawDescGZIP() []byte {
	file_arbstream_protobuf_arbstream_proto_rawDescOnce.Do(func() {
		file_arbstream_protobuf_arbstream_proto_rawDescData = protoimpl.X.CompressGZIP(file_arbstream_protobuf_arbstream_proto_rawDescData)
	})
	return file_arbstream_protobuf_arbstream_proto_rawDescData
}

var file_arbstream_protobuf_arbstream_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_arbstream_protobuf_arbstream_proto_goTypes = []any{
	(*Order)(nil),         // 0: arbstream.Order
	(*StateChange)(nil),   // 1: arbstream.StateChange
	(*Transaction)(nil),   // 2: arbstream.Transaction
	(*TxContents)(nil),    // 3: arbstream.TxContents
	(*Log)(nil),           // 4: arbstream.Log
	(*StreamRequest)(nil), // 5: arbstream.StreamRequest
	nil,                   // 6: arbstream.Order.StateEntry
	nil,                   // 7: arbstream.StateChange.SlotsEntry
}
var file_arbstream_protobuf_arbstream_proto_depIdxs = []int32{
	2, // 0: arbstream.Order.transactions:type_name -> arbstream.Transaction
	6, // 1: arbstream.Order.state:type_name -> arbstream.Order.StateEntry
	7, // 2: arbstream.StateChange.slots:type_name -> arbstream.StateChange.SlotsEntry
	3, // 3: arbstream.Transaction.txContents:type_name -> arbstream.TxContents
	4, // 4: arbstream.TxContents.logs:type_name -> arbstream.Log
	1, // 5: arbstream.Order.StateEntry.value:type_name -> arbstream.StateChange
	5, // 6: arbstream.ArbStreamServer.StreamArbData:input_type -> arbstream.StreamRequest
	0, // 7: arbstream.ArbStreamServer.StreamArbData:output_type -> arbstream.Order
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_arbstream_protobuf_arbstream_proto_init() }
func file_arbstream_protobuf_arbstream_proto_init() {
	if File_arbstream_protobuf_arbstream_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_arbstream_protobuf_arbstream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_arbstream_protobuf_arbstream_proto_goTypes,
		DependencyIndexes: file_arbstream_protobuf_arbstream_proto_depIdxs,
		MessageInfos:      file_arbstream_protobuf_arbstream_proto_msgTypes,
	}.Build()
	File_arbstream_protobuf_arbstream_proto = out.File
	file_arbstream_protobuf_arbstream_proto_rawDesc = nil
	file_arbstream_protobuf_arbstream_proto_goTypes = nil
	file_arbstream_protobuf_arbstream_proto_depIdxs = nil
}
