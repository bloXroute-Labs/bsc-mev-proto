// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: compressed_bid_message.proto

package mevrelaypb

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

// CompressedRawBid is a compressed representation of a raw bid.
type CompressedRawBid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The block number of the bid.
	BlockNumber uint64 `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	// The hash of the parent block.
	ParentHash string `protobuf:"bytes,2,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	// The transactions in the bid.
	Txs []*CompressedTx `protobuf:"bytes,3,rep,name=txs,proto3" json:"txs,omitempty"`
	// The unRevertible hashes of the bid.
	UnRevertible []*CompressedHash `protobuf:"bytes,4,rep,name=unRevertible,proto3" json:"unRevertible,omitempty"`
	// The gas used in the bid.
	GasUsed uint64 `protobuf:"varint,5,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	// The gas fee in the bid.
	GasFee *BigInt `protobuf:"bytes,6,opt,name=gas_fee,json=gasFee,proto3" json:"gas_fee,omitempty"`
	// The builder fee in the bid.
	BuilderFee *BigInt `protobuf:"bytes,7,opt,name=builder_fee,json=builderFee,proto3" json:"builder_fee,omitempty"`
}

func (x *CompressedRawBid) Reset() {
	*x = CompressedRawBid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compressed_bid_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompressedRawBid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompressedRawBid) ProtoMessage() {}

func (x *CompressedRawBid) ProtoReflect() protoreflect.Message {
	mi := &file_compressed_bid_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompressedRawBid.ProtoReflect.Descriptor instead.
func (*CompressedRawBid) Descriptor() ([]byte, []int) {
	return file_compressed_bid_message_proto_rawDescGZIP(), []int{0}
}

func (x *CompressedRawBid) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *CompressedRawBid) GetParentHash() string {
	if x != nil {
		return x.ParentHash
	}
	return ""
}

func (x *CompressedRawBid) GetTxs() []*CompressedTx {
	if x != nil {
		return x.Txs
	}
	return nil
}

func (x *CompressedRawBid) GetUnRevertible() []*CompressedHash {
	if x != nil {
		return x.UnRevertible
	}
	return nil
}

func (x *CompressedRawBid) GetGasUsed() uint64 {
	if x != nil {
		return x.GasUsed
	}
	return 0
}

func (x *CompressedRawBid) GetGasFee() *BigInt {
	if x != nil {
		return x.GasFee
	}
	return nil
}

func (x *CompressedRawBid) GetBuilderFee() *BigInt {
	if x != nil {
		return x.BuilderFee
	}
	return nil
}

var File_compressed_bid_message_proto protoreflect.FileDescriptor

var file_compressed_bid_message_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x62, 0x69, 0x64,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x6c, 0x6f, 0x78, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x62, 0x73,
	0x63, 0x2e, 0x6d, 0x65, 0x76, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x14,
	0x62, 0x69, 0x67, 0x69, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f,
	0x74, 0x78, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8f, 0x03, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x52,
	0x61, 0x77, 0x42, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x3e, 0x0a, 0x03, 0x74, 0x78, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x6c, 0x6f,
	0x78, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x62, 0x73, 0x63, 0x2e, 0x6d, 0x65, 0x76, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x64, 0x54, 0x78, 0x52, 0x03, 0x74, 0x78, 0x73, 0x12, 0x52, 0x0a, 0x0c, 0x75, 0x6e, 0x52,
	0x65, 0x76, 0x65, 0x72, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x6c, 0x6f, 0x78, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e,
	0x62, 0x73, 0x63, 0x2e, 0x6d, 0x65, 0x76, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x48, 0x61, 0x73, 0x68, 0x52,
	0x0c, 0x75, 0x6e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x67, 0x61, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x67, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x3f, 0x0a, 0x07, 0x67, 0x61, 0x73, 0x5f,
	0x66, 0x65, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x62, 0x6c, 0x6f, 0x78, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x62, 0x73, 0x63, 0x2e, 0x6d, 0x65,
	0x76, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e,
	0x74, 0x52, 0x06, 0x67, 0x61, 0x73, 0x46, 0x65, 0x65, 0x12, 0x47, 0x0a, 0x0b, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x6c, 0x6f, 0x78, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x62,
	0x73, 0x63, 0x2e, 0x6d, 0x65, 0x76, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x46,
	0x65, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x6c, 0x6f, 0x58, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2d, 0x4c, 0x61, 0x62, 0x73, 0x2f,
	0x62, 0x73, 0x63, 0x2d, 0x6d, 0x65, 0x76, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x2f, 0x6d, 0x65, 0x76, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_compressed_bid_message_proto_rawDescOnce sync.Once
	file_compressed_bid_message_proto_rawDescData = file_compressed_bid_message_proto_rawDesc
)

func file_compressed_bid_message_proto_rawDescGZIP() []byte {
	file_compressed_bid_message_proto_rawDescOnce.Do(func() {
		file_compressed_bid_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_compressed_bid_message_proto_rawDescData)
	})
	return file_compressed_bid_message_proto_rawDescData
}

var file_compressed_bid_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_compressed_bid_message_proto_goTypes = []interface{}{
	(*CompressedRawBid)(nil), // 0: com.bloxroute.bsc.mev.relay.v1.CompressedRawBid
	(*CompressedTx)(nil),     // 1: com.bloxroute.bsc.mev.relay.v1.CompressedTx
	(*CompressedHash)(nil),   // 2: com.bloxroute.bsc.mev.relay.v1.CompressedHash
	(*BigInt)(nil),           // 3: com.bloxroute.bsc.mev.relay.v1.BigInt
}
var file_compressed_bid_message_proto_depIdxs = []int32{
	1, // 0: com.bloxroute.bsc.mev.relay.v1.CompressedRawBid.txs:type_name -> com.bloxroute.bsc.mev.relay.v1.CompressedTx
	2, // 1: com.bloxroute.bsc.mev.relay.v1.CompressedRawBid.unRevertible:type_name -> com.bloxroute.bsc.mev.relay.v1.CompressedHash
	3, // 2: com.bloxroute.bsc.mev.relay.v1.CompressedRawBid.gas_fee:type_name -> com.bloxroute.bsc.mev.relay.v1.BigInt
	3, // 3: com.bloxroute.bsc.mev.relay.v1.CompressedRawBid.builder_fee:type_name -> com.bloxroute.bsc.mev.relay.v1.BigInt
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_compressed_bid_message_proto_init() }
func file_compressed_bid_message_proto_init() {
	if File_compressed_bid_message_proto != nil {
		return
	}
	file_bigint_message_proto_init()
	file_compressed_hash_message_proto_init()
	file_compressed_tx_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_compressed_bid_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompressedRawBid); i {
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
			RawDescriptor: file_compressed_bid_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_compressed_bid_message_proto_goTypes,
		DependencyIndexes: file_compressed_bid_message_proto_depIdxs,
		MessageInfos:      file_compressed_bid_message_proto_msgTypes,
	}.Build()
	File_compressed_bid_message_proto = out.File
	file_compressed_bid_message_proto_rawDesc = nil
	file_compressed_bid_message_proto_goTypes = nil
	file_compressed_bid_message_proto_depIdxs = nil
}
