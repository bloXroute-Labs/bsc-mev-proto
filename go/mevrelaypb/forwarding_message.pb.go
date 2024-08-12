// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: forwarding_message.proto

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

// ForwardRequest is the request to forward a bundle to a validator.
type ForwardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is a unique identifier for the request.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The account ID of the builder.
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// validator is the validator to forward the bundle to.
	Validator string `protobuf:"bytes,3,opt,name=validator,proto3" json:"validator,omitempty"`
	// bundle_hashes is the hashes of the transactions in the bundle.
	BundleHashes *Hashes `protobuf:"bytes,4,opt,name=bundle_hashes,json=bundleHashes,proto3" json:"bundle_hashes,omitempty"`
	// The payment transaction for the bid.
	PayBidTx []byte `protobuf:"bytes,5,opt,name=pay_bid_tx,json=payBidTx,proto3" json:"pay_bid_tx,omitempty"`
	// The gas used for the payment transaction.
	PayBidTxGasUsed *BigInt `protobuf:"bytes,6,opt,name=pay_bid_tx_gas_used,json=payBidTxGasUsed,proto3" json:"pay_bid_tx_gas_used,omitempty"`
	// The raw bid.
	RawBid *CompressedRawBid `protobuf:"bytes,7,opt,name=raw_bid,json=rawBid,proto3" json:"raw_bid,omitempty"`
}

func (x *ForwardRequest) Reset() {
	*x = ForwardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forwarding_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForwardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForwardRequest) ProtoMessage() {}

func (x *ForwardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_forwarding_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForwardRequest.ProtoReflect.Descriptor instead.
func (*ForwardRequest) Descriptor() ([]byte, []int) {
	return file_forwarding_message_proto_rawDescGZIP(), []int{0}
}

func (x *ForwardRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ForwardRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *ForwardRequest) GetValidator() string {
	if x != nil {
		return x.Validator
	}
	return ""
}

func (x *ForwardRequest) GetBundleHashes() *Hashes {
	if x != nil {
		return x.BundleHashes
	}
	return nil
}

func (x *ForwardRequest) GetPayBidTx() []byte {
	if x != nil {
		return x.PayBidTx
	}
	return nil
}

func (x *ForwardRequest) GetPayBidTxGasUsed() *BigInt {
	if x != nil {
		return x.PayBidTxGasUsed
	}
	return nil
}

func (x *ForwardRequest) GetRawBid() *CompressedRawBid {
	if x != nil {
		return x.RawBid
	}
	return nil
}

// ForwardResponse is the response to a ForwardRequest.
type ForwardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *ProposeBlockResponse `protobuf:"bytes,1,opt,name=Response,proto3" json:"Response,omitempty"`
}

func (x *ForwardResponse) Reset() {
	*x = ForwardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forwarding_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForwardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForwardResponse) ProtoMessage() {}

func (x *ForwardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_forwarding_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForwardResponse.ProtoReflect.Descriptor instead.
func (*ForwardResponse) Descriptor() ([]byte, []int) {
	return file_forwarding_message_proto_rawDescGZIP(), []int{1}
}

func (x *ForwardResponse) GetResponse() *ProposeBlockResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_forwarding_message_proto protoreflect.FileDescriptor

var file_forwarding_message_proto_rawDesc = []byte{
	0x0a, 0x18, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6f, 0x6d, 0x2e,
	0x62, 0x6c, 0x6f, 0x78, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x62, 0x73, 0x63, 0x2e, 0x6d, 0x65,
	0x76, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x62, 0x69, 0x67, 0x69,
	0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x62, 0x69, 0x64,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x68, 0x61, 0x73, 0x68, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe9, 0x02, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x4b, 0x0a, 0x0d, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x6c, 0x6f,
	0x78, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x62, 0x73, 0x63, 0x2e, 0x6d, 0x65, 0x76, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x52, 0x0c,
	0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x0a,
	0x70, 0x61, 0x79, 0x5f, 0x62, 0x69, 0x64, 0x5f, 0x74, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x70, 0x61, 0x79, 0x42, 0x69, 0x64, 0x54, 0x78, 0x12, 0x54, 0x0a, 0x13, 0x70, 0x61,
	0x79, 0x5f, 0x62, 0x69, 0x64, 0x5f, 0x74, 0x78, 0x5f, 0x67, 0x61, 0x73, 0x5f, 0x75, 0x73, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x6c,
	0x6f, 0x78, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x62, 0x73, 0x63, 0x2e, 0x6d, 0x65, 0x76, 0x2e,
	0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52,
	0x0f, 0x70, 0x61, 0x79, 0x42, 0x69, 0x64, 0x54, 0x78, 0x47, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64,
	0x12, 0x49, 0x0a, 0x07, 0x72, 0x61, 0x77, 0x5f, 0x62, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x6c, 0x6f, 0x78, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x2e, 0x62, 0x73, 0x63, 0x2e, 0x6d, 0x65, 0x76, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x52, 0x61, 0x77,
	0x42, 0x69, 0x64, 0x52, 0x06, 0x72, 0x61, 0x77, 0x42, 0x69, 0x64, 0x22, 0x63, 0x0a, 0x0f, 0x46,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x6c, 0x6f, 0x78, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x2e, 0x62, 0x73, 0x63, 0x2e, 0x6d, 0x65, 0x76, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62,
	0x6c, 0x6f, 0x58, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2d, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x73,
	0x63, 0x2d, 0x6d, 0x65, 0x76, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x6d,
	0x65, 0x76, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_forwarding_message_proto_rawDescOnce sync.Once
	file_forwarding_message_proto_rawDescData = file_forwarding_message_proto_rawDesc
)

func file_forwarding_message_proto_rawDescGZIP() []byte {
	file_forwarding_message_proto_rawDescOnce.Do(func() {
		file_forwarding_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_forwarding_message_proto_rawDescData)
	})
	return file_forwarding_message_proto_rawDescData
}

var file_forwarding_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_forwarding_message_proto_goTypes = []interface{}{
	(*ForwardRequest)(nil),       // 0: com.bloxroute.bsc.mev.relay.v1.ForwardRequest
	(*ForwardResponse)(nil),      // 1: com.bloxroute.bsc.mev.relay.v1.ForwardResponse
	(*Hashes)(nil),               // 2: com.bloxroute.bsc.mev.relay.v1.Hashes
	(*BigInt)(nil),               // 3: com.bloxroute.bsc.mev.relay.v1.BigInt
	(*CompressedRawBid)(nil),     // 4: com.bloxroute.bsc.mev.relay.v1.CompressedRawBid
	(*ProposeBlockResponse)(nil), // 5: com.bloxroute.bsc.mev.relay.v1.ProposeBlockResponse
}
var file_forwarding_message_proto_depIdxs = []int32{
	2, // 0: com.bloxroute.bsc.mev.relay.v1.ForwardRequest.bundle_hashes:type_name -> com.bloxroute.bsc.mev.relay.v1.Hashes
	3, // 1: com.bloxroute.bsc.mev.relay.v1.ForwardRequest.pay_bid_tx_gas_used:type_name -> com.bloxroute.bsc.mev.relay.v1.BigInt
	4, // 2: com.bloxroute.bsc.mev.relay.v1.ForwardRequest.raw_bid:type_name -> com.bloxroute.bsc.mev.relay.v1.CompressedRawBid
	5, // 3: com.bloxroute.bsc.mev.relay.v1.ForwardResponse.Response:type_name -> com.bloxroute.bsc.mev.relay.v1.ProposeBlockResponse
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_forwarding_message_proto_init() }
func file_forwarding_message_proto_init() {
	if File_forwarding_message_proto != nil {
		return
	}
	file_bigint_message_proto_init()
	file_compressed_bid_message_proto_init()
	file_hash_message_proto_init()
	file_propose_block_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_forwarding_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForwardRequest); i {
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
		file_forwarding_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForwardResponse); i {
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
			RawDescriptor: file_forwarding_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_forwarding_message_proto_goTypes,
		DependencyIndexes: file_forwarding_message_proto_depIdxs,
		MessageInfos:      file_forwarding_message_proto_msgTypes,
	}.Build()
	File_forwarding_message_proto = out.File
	file_forwarding_message_proto_rawDesc = nil
	file_forwarding_message_proto_goTypes = nil
	file_forwarding_message_proto_depIdxs = nil
}
