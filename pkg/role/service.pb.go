// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: pkg/role/pb/service.proto

package role

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

var File_pkg_role_pb_service_proto protoreflect.FileDescriptor

var file_pkg_role_pb_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x1a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70,
	0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xaa, 0x07, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b,
	0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x12, 0x54, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x29, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x12, 0x5b, 0x0a, 0x0c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x2c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65,
	0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x57, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c,
	0x65, 0x12, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b,
	0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x6a, 0x0a, 0x0f,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x2f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65,
	0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x6d, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x32,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x72, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x33,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x7c, 0x0a, 0x18, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x46,
	0x72, 0x6f, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x38, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b,
	0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x69, 0x0a, 0x10, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65,
	0x79, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_pkg_role_pb_service_proto_goTypes = []interface{}{
	(*CreateRoleRequest)(nil),               // 0: infraboard.keyauth.role.CreateRoleRequest
	(*QueryRoleRequest)(nil),                // 1: infraboard.keyauth.role.QueryRoleRequest
	(*DescribeRoleRequest)(nil),             // 2: infraboard.keyauth.role.DescribeRoleRequest
	(*DeleteRoleRequest)(nil),               // 3: infraboard.keyauth.role.DeleteRoleRequest
	(*QueryPermissionRequest)(nil),          // 4: infraboard.keyauth.role.QueryPermissionRequest
	(*DescribePermissionRequest)(nil),       // 5: infraboard.keyauth.role.DescribePermissionRequest
	(*AddPermissionToRoleRequest)(nil),      // 6: infraboard.keyauth.role.AddPermissionToRoleRequest
	(*RemovePermissionFromRoleRequest)(nil), // 7: infraboard.keyauth.role.RemovePermissionFromRoleRequest
	(*UpdatePermissionRequest)(nil),         // 8: infraboard.keyauth.role.UpdatePermissionRequest
	(*Role)(nil),                            // 9: infraboard.keyauth.role.Role
	(*Set)(nil),                             // 10: infraboard.keyauth.role.Set
	(*PermissionSet)(nil),                   // 11: infraboard.keyauth.role.PermissionSet
	(*Permission)(nil),                      // 12: infraboard.keyauth.role.Permission
}
var file_pkg_role_pb_service_proto_depIdxs = []int32{
	0,  // 0: infraboard.keyauth.role.RoleService.CreateRole:input_type -> infraboard.keyauth.role.CreateRoleRequest
	1,  // 1: infraboard.keyauth.role.RoleService.QueryRole:input_type -> infraboard.keyauth.role.QueryRoleRequest
	2,  // 2: infraboard.keyauth.role.RoleService.DescribeRole:input_type -> infraboard.keyauth.role.DescribeRoleRequest
	3,  // 3: infraboard.keyauth.role.RoleService.DeleteRole:input_type -> infraboard.keyauth.role.DeleteRoleRequest
	4,  // 4: infraboard.keyauth.role.RoleService.QueryPermission:input_type -> infraboard.keyauth.role.QueryPermissionRequest
	5,  // 5: infraboard.keyauth.role.RoleService.DescribePermission:input_type -> infraboard.keyauth.role.DescribePermissionRequest
	6,  // 6: infraboard.keyauth.role.RoleService.AddPermissionToRole:input_type -> infraboard.keyauth.role.AddPermissionToRoleRequest
	7,  // 7: infraboard.keyauth.role.RoleService.RemovePermissionFromRole:input_type -> infraboard.keyauth.role.RemovePermissionFromRoleRequest
	8,  // 8: infraboard.keyauth.role.RoleService.UpdatePermission:input_type -> infraboard.keyauth.role.UpdatePermissionRequest
	9,  // 9: infraboard.keyauth.role.RoleService.CreateRole:output_type -> infraboard.keyauth.role.Role
	10, // 10: infraboard.keyauth.role.RoleService.QueryRole:output_type -> infraboard.keyauth.role.Set
	9,  // 11: infraboard.keyauth.role.RoleService.DescribeRole:output_type -> infraboard.keyauth.role.Role
	9,  // 12: infraboard.keyauth.role.RoleService.DeleteRole:output_type -> infraboard.keyauth.role.Role
	11, // 13: infraboard.keyauth.role.RoleService.QueryPermission:output_type -> infraboard.keyauth.role.PermissionSet
	12, // 14: infraboard.keyauth.role.RoleService.DescribePermission:output_type -> infraboard.keyauth.role.Permission
	11, // 15: infraboard.keyauth.role.RoleService.AddPermissionToRole:output_type -> infraboard.keyauth.role.PermissionSet
	11, // 16: infraboard.keyauth.role.RoleService.RemovePermissionFromRole:output_type -> infraboard.keyauth.role.PermissionSet
	12, // 17: infraboard.keyauth.role.RoleService.UpdatePermission:output_type -> infraboard.keyauth.role.Permission
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_role_pb_service_proto_init() }
func file_pkg_role_pb_service_proto_init() {
	if File_pkg_role_pb_service_proto != nil {
		return
	}
	file_pkg_role_pb_request_proto_init()
	file_pkg_role_pb_role_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_role_pb_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_role_pb_service_proto_goTypes,
		DependencyIndexes: file_pkg_role_pb_service_proto_depIdxs,
	}.Build()
	File_pkg_role_pb_service_proto = out.File
	file_pkg_role_pb_service_proto_rawDesc = nil
	file_pkg_role_pb_service_proto_goTypes = nil
	file_pkg_role_pb_service_proto_depIdxs = nil
}
