// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: pkg/role/pb/enum.proto

package role

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

// RoleType 角色
type RoleType int32

const (
	RoleType_NULL RoleType = 0
	// 内建角色, 系统初始时创建
	RoleType_BUILDIN RoleType = 1
	// 管理员创建的一些角色, 全局可用
	RoleType_GLOBAL RoleType = 2
	// 用户自定义的角色, 仅域内可见
	RoleType_CUSTOM RoleType = 3
)

// Enum value maps for RoleType.
var (
	RoleType_name = map[int32]string{
		0: "NULL",
		1: "BUILDIN",
		2: "GLOBAL",
		3: "CUSTOM",
	}
	RoleType_value = map[string]int32{
		"NULL":    0,
		"BUILDIN": 1,
		"GLOBAL":  2,
		"CUSTOM":  3,
	}
)

func (x RoleType) Enum() *RoleType {
	p := new(RoleType)
	*p = x
	return p
}

func (x RoleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoleType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_role_pb_enum_proto_enumTypes[0].Descriptor()
}

func (RoleType) Type() protoreflect.EnumType {
	return &file_pkg_role_pb_enum_proto_enumTypes[0]
}

func (x RoleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoleType.Descriptor instead.
func (RoleType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_role_pb_enum_proto_rawDescGZIP(), []int{0}
}

// EffectType 授权效力包括两种：允许（Allow）和拒绝（Deny）
type EffectType int32

const (
	// 允许访问
	EffectType_ALLOW EffectType = 0
	// 拒绝访问
	EffectType_DENY EffectType = 1
)

// Enum value maps for EffectType.
var (
	EffectType_name = map[int32]string{
		0: "ALLOW",
		1: "DENY",
	}
	EffectType_value = map[string]int32{
		"ALLOW": 0,
		"DENY":  1,
	}
)

func (x EffectType) Enum() *EffectType {
	p := new(EffectType)
	*p = x
	return p
}

func (x EffectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EffectType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_role_pb_enum_proto_enumTypes[1].Descriptor()
}

func (EffectType) Type() protoreflect.EnumType {
	return &file_pkg_role_pb_enum_proto_enumTypes[1]
}

func (x EffectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EffectType.Descriptor instead.
func (EffectType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_role_pb_enum_proto_rawDescGZIP(), []int{1}
}

var File_pkg_role_pb_enum_proto protoreflect.FileDescriptor

var file_pkg_role_pb_enum_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x2a, 0x39, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a,
	0x04, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x55, 0x49, 0x4c, 0x44,
	0x49, 0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x10, 0x02,
	0x12, 0x0a, 0x0a, 0x06, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x03, 0x2a, 0x21, 0x0a, 0x0a,
	0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c,
	0x4c, 0x4f, 0x57, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x4e, 0x59, 0x10, 0x01, 0x42,
	0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_role_pb_enum_proto_rawDescOnce sync.Once
	file_pkg_role_pb_enum_proto_rawDescData = file_pkg_role_pb_enum_proto_rawDesc
)

func file_pkg_role_pb_enum_proto_rawDescGZIP() []byte {
	file_pkg_role_pb_enum_proto_rawDescOnce.Do(func() {
		file_pkg_role_pb_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_role_pb_enum_proto_rawDescData)
	})
	return file_pkg_role_pb_enum_proto_rawDescData
}

var file_pkg_role_pb_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pkg_role_pb_enum_proto_goTypes = []interface{}{
	(RoleType)(0),   // 0: infraboard.keyauth.role.RoleType
	(EffectType)(0), // 1: infraboard.keyauth.role.EffectType
}
var file_pkg_role_pb_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_role_pb_enum_proto_init() }
func file_pkg_role_pb_enum_proto_init() {
	if File_pkg_role_pb_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_role_pb_enum_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_role_pb_enum_proto_goTypes,
		DependencyIndexes: file_pkg_role_pb_enum_proto_depIdxs,
		EnumInfos:         file_pkg_role_pb_enum_proto_enumTypes,
	}.Build()
	File_pkg_role_pb_enum_proto = out.File
	file_pkg_role_pb_enum_proto_rawDesc = nil
	file_pkg_role_pb_enum_proto_goTypes = nil
	file_pkg_role_pb_enum_proto_depIdxs = nil
}
