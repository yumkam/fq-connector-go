// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.0
// source: ydb/library/yql/providers/generic/connector/api/service/connector.proto

package service

import (
	protos "github.com/ydb-platform/fq-connector-go/libgo/service/protos"
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

var File_ydb_library_yql_providers_generic_connector_api_service_connector_proto protoreflect.FileDescriptor

var file_ydb_library_yql_providers_generic_connector_api_service_connector_proto_rawDesc = []byte{
	0x0a, 0x47, 0x79, 0x64, 0x62, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x79, 0x71,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x69, 0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x4e, 0x59, 0x71, 0x6c, 0x2e,
	0x4e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4e, 0x41, 0x70, 0x69, 0x1a,
	0x4e, 0x79, 0x64, 0x62, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x79, 0x71, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x69, 0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xa6, 0x03, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x63, 0x0a,
	0x0a, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x28, 0x2e, 0x4e, 0x59,
	0x71, 0x6c, 0x2e, 0x4e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4e, 0x41,
	0x70, 0x69, 0x2e, 0x54, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x4e, 0x59, 0x71, 0x6c, 0x2e, 0x4e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4e, 0x41, 0x70, 0x69, 0x2e, 0x54, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x30, 0x01, 0x12, 0x6a, 0x0a, 0x0d, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x2b, 0x2e, 0x4e, 0x59, 0x71, 0x6c, 0x2e, 0x4e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4e, 0x41, 0x70, 0x69, 0x2e, 0x54, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2c, 0x2e, 0x4e, 0x59, 0x71, 0x6c, 0x2e, 0x4e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x4e, 0x41, 0x70, 0x69, 0x2e, 0x54, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63,
	0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x12, 0x28, 0x2e, 0x4e,
	0x59, 0x71, 0x6c, 0x2e, 0x4e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4e,
	0x41, 0x70, 0x69, 0x2e, 0x54, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x4e, 0x59, 0x71, 0x6c, 0x2e, 0x4e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4e, 0x41, 0x70, 0x69, 0x2e, 0x54, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x30, 0x01, 0x12, 0x63, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x53, 0x70, 0x6c, 0x69, 0x74,
	0x73, 0x12, 0x28, 0x2e, 0x4e, 0x59, 0x71, 0x6c, 0x2e, 0x4e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x4e, 0x41, 0x70, 0x69, 0x2e, 0x54, 0x52, 0x65, 0x61, 0x64, 0x53, 0x70,
	0x6c, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x4e, 0x59,
	0x71, 0x6c, 0x2e, 0x4e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4e, 0x41,
	0x70, 0x69, 0x2e, 0x54, 0x52, 0x65, 0x61, 0x64, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x4c, 0x5a, 0x4a, 0x61, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x72, 0x75, 0x2f, 0x79, 0x64, 0x62,
	0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x79, 0x71, 0x6c, 0x2f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x6c, 0x69, 0x62, 0x67, 0x6f, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_ydb_library_yql_providers_generic_connector_api_service_connector_proto_goTypes = []interface{}{
	(*protos.TListTablesRequest)(nil),     // 0: NYql.NConnector.NApi.TListTablesRequest
	(*protos.TDescribeTableRequest)(nil),  // 1: NYql.NConnector.NApi.TDescribeTableRequest
	(*protos.TListSplitsRequest)(nil),     // 2: NYql.NConnector.NApi.TListSplitsRequest
	(*protos.TReadSplitsRequest)(nil),     // 3: NYql.NConnector.NApi.TReadSplitsRequest
	(*protos.TListTablesResponse)(nil),    // 4: NYql.NConnector.NApi.TListTablesResponse
	(*protos.TDescribeTableResponse)(nil), // 5: NYql.NConnector.NApi.TDescribeTableResponse
	(*protos.TListSplitsResponse)(nil),    // 6: NYql.NConnector.NApi.TListSplitsResponse
	(*protos.TReadSplitsResponse)(nil),    // 7: NYql.NConnector.NApi.TReadSplitsResponse
}
var file_ydb_library_yql_providers_generic_connector_api_service_connector_proto_depIdxs = []int32{
	0, // 0: NYql.NConnector.NApi.Connector.ListTables:input_type -> NYql.NConnector.NApi.TListTablesRequest
	1, // 1: NYql.NConnector.NApi.Connector.DescribeTable:input_type -> NYql.NConnector.NApi.TDescribeTableRequest
	2, // 2: NYql.NConnector.NApi.Connector.ListSplits:input_type -> NYql.NConnector.NApi.TListSplitsRequest
	3, // 3: NYql.NConnector.NApi.Connector.ReadSplits:input_type -> NYql.NConnector.NApi.TReadSplitsRequest
	4, // 4: NYql.NConnector.NApi.Connector.ListTables:output_type -> NYql.NConnector.NApi.TListTablesResponse
	5, // 5: NYql.NConnector.NApi.Connector.DescribeTable:output_type -> NYql.NConnector.NApi.TDescribeTableResponse
	6, // 6: NYql.NConnector.NApi.Connector.ListSplits:output_type -> NYql.NConnector.NApi.TListSplitsResponse
	7, // 7: NYql.NConnector.NApi.Connector.ReadSplits:output_type -> NYql.NConnector.NApi.TReadSplitsResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ydb_library_yql_providers_generic_connector_api_service_connector_proto_init() }
func file_ydb_library_yql_providers_generic_connector_api_service_connector_proto_init() {
	if File_ydb_library_yql_providers_generic_connector_api_service_connector_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ydb_library_yql_providers_generic_connector_api_service_connector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ydb_library_yql_providers_generic_connector_api_service_connector_proto_goTypes,
		DependencyIndexes: file_ydb_library_yql_providers_generic_connector_api_service_connector_proto_depIdxs,
	}.Build()
	File_ydb_library_yql_providers_generic_connector_api_service_connector_proto = out.File
	file_ydb_library_yql_providers_generic_connector_api_service_connector_proto_rawDesc = nil
	file_ydb_library_yql_providers_generic_connector_api_service_connector_proto_goTypes = nil
	file_ydb_library_yql_providers_generic_connector_api_service_connector_proto_depIdxs = nil
}
