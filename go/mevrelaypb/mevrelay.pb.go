// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: mevrelay.proto

package mevrelaypb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_mevrelay_proto protoreflect.FileDescriptor

var file_mevrelay_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x76, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x6c, 0x6f, 0x78, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e,
	0x62, 0x73, 0x63, 0x2e, 0x6d, 0x65, 0x76, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x89, 0x02, 0x0a, 0x08,
	0x4d, 0x65, 0x76, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x7d, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x62,
	0x6c, 0x6f, 0x78, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x62, 0x73, 0x63, 0x2e, 0x6d, 0x65, 0x76,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x6c, 0x6f, 0x78, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x62, 0x73,
	0x63, 0x2e, 0x6d, 0x65, 0x76, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x7e, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x62,
	0x6c, 0x6f, 0x78, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x62, 0x73, 0x63, 0x2e, 0x6d, 0x65, 0x76,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x6c, 0x6f, 0x78, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x62,
	0x73, 0x63, 0x2e, 0x6d, 0x65, 0x76, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6c, 0x6f, 0x58, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2d,
	0x4c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x73, 0x63, 0x2d, 0x6d, 0x65, 0x76, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x65, 0x76, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_mevrelay_proto_goTypes = []interface{}{
	(*ProposeBlockRequest)(nil),   // 0: com.bloxroute.bsc.mev.relay.v1.ProposeBlockRequest
	(*ConfigUpdatesRequest)(nil),  // 1: com.bloxroute.bsc.mev.relay.v1.ConfigUpdatesRequest
	(*ProposeBlockResponse)(nil),  // 2: com.bloxroute.bsc.mev.relay.v1.ProposeBlockResponse
	(*ConfigUpdatesResponse)(nil), // 3: com.bloxroute.bsc.mev.relay.v1.ConfigUpdatesResponse
}
var file_mevrelay_proto_depIdxs = []int32{
	0, // 0: com.bloxroute.bsc.mev.relay.v1.MevRelay.ProposeBlock:input_type -> com.bloxroute.bsc.mev.relay.v1.ProposeBlockRequest
	1, // 1: com.bloxroute.bsc.mev.relay.v1.MevRelay.ConfigUpdates:input_type -> com.bloxroute.bsc.mev.relay.v1.ConfigUpdatesRequest
	2, // 2: com.bloxroute.bsc.mev.relay.v1.MevRelay.ProposeBlock:output_type -> com.bloxroute.bsc.mev.relay.v1.ProposeBlockResponse
	3, // 3: com.bloxroute.bsc.mev.relay.v1.MevRelay.ConfigUpdates:output_type -> com.bloxroute.bsc.mev.relay.v1.ConfigUpdatesResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mevrelay_proto_init() }
func file_mevrelay_proto_init() {
	if File_mevrelay_proto != nil {
		return
	}
	file_config_updates_message_proto_init()
	file_propose_block_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mevrelay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mevrelay_proto_goTypes,
		DependencyIndexes: file_mevrelay_proto_depIdxs,
	}.Build()
	File_mevrelay_proto = out.File
	file_mevrelay_proto_rawDesc = nil
	file_mevrelay_proto_goTypes = nil
	file_mevrelay_proto_depIdxs = nil
}
