// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: app/policy/pb/service.proto

package policy

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

var File_app_policy_pb_service_proto protoreflect.FileDescriptor

var file_app_policy_pb_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x70, 0x62, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x1b, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0x94, 0x03, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2e, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x5c, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x2d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x65, 0x74, 0x12, 0x65,
	0x0a, 0x0e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x30, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65,
	0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x61, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_app_policy_pb_service_proto_goTypes = []interface{}{
	(*CreatePolicyRequest)(nil),   // 0: infraboard.keyauth.policy.CreatePolicyRequest
	(*QueryPolicyRequest)(nil),    // 1: infraboard.keyauth.policy.QueryPolicyRequest
	(*DescribePolicyRequest)(nil), // 2: infraboard.keyauth.policy.DescribePolicyRequest
	(*DeletePolicyRequest)(nil),   // 3: infraboard.keyauth.policy.DeletePolicyRequest
	(*Policy)(nil),                // 4: infraboard.keyauth.policy.Policy
	(*Set)(nil),                   // 5: infraboard.keyauth.policy.Set
}
var file_app_policy_pb_service_proto_depIdxs = []int32{
	0, // 0: infraboard.keyauth.policy.Service.CreatePolicy:input_type -> infraboard.keyauth.policy.CreatePolicyRequest
	1, // 1: infraboard.keyauth.policy.Service.QueryPolicy:input_type -> infraboard.keyauth.policy.QueryPolicyRequest
	2, // 2: infraboard.keyauth.policy.Service.DescribePolicy:input_type -> infraboard.keyauth.policy.DescribePolicyRequest
	3, // 3: infraboard.keyauth.policy.Service.DeletePolicy:input_type -> infraboard.keyauth.policy.DeletePolicyRequest
	4, // 4: infraboard.keyauth.policy.Service.CreatePolicy:output_type -> infraboard.keyauth.policy.Policy
	5, // 5: infraboard.keyauth.policy.Service.QueryPolicy:output_type -> infraboard.keyauth.policy.Set
	4, // 6: infraboard.keyauth.policy.Service.DescribePolicy:output_type -> infraboard.keyauth.policy.Policy
	4, // 7: infraboard.keyauth.policy.Service.DeletePolicy:output_type -> infraboard.keyauth.policy.Policy
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_policy_pb_service_proto_init() }
func file_app_policy_pb_service_proto_init() {
	if File_app_policy_pb_service_proto != nil {
		return
	}
	file_app_policy_pb_request_proto_init()
	file_app_policy_pb_policy_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_policy_pb_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_policy_pb_service_proto_goTypes,
		DependencyIndexes: file_app_policy_pb_service_proto_depIdxs,
	}.Build()
	File_app_policy_pb_service_proto = out.File
	file_app_policy_pb_service_proto_rawDesc = nil
	file_app_policy_pb_service_proto_goTypes = nil
	file_app_policy_pb_service_proto_depIdxs = nil
}
