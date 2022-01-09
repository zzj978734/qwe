// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: apps/user/pb/enum.proto

package user

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

type Gender int32

const (
	Gender_UNKNOWN Gender = 0
	Gender_MALE    Gender = 1
	Gender_FEMALE  Gender = 2
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "UNKNOWN",
		1: "MALE",
		2: "FEMALE",
	}
	Gender_value = map[string]int32{
		"UNKNOWN": 0,
		"MALE":    1,
		"FEMALE":  2,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_user_pb_enum_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_apps_user_pb_enum_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_apps_user_pb_enum_proto_rawDescGZIP(), []int{0}
}

type CreateType int32

const (
	// UserRegistry (user_registry) 用户自己注册的用户
	CreateType_USER_REGISTRY CreateType = 0
	// DOMAIN_CREATED (domain_admin) 域管理员创建的用户
	CreateType_DOMAIN_CREATED CreateType = 1
	// LDAPSync (ldap_sync) LDAP同步的用户
	CreateType_LDAP_SYNC CreateType = 2
	// 企业微信同步的用户;
	CreateType_WXWORK_SYNC CreateType = 3
)

// Enum value maps for CreateType.
var (
	CreateType_name = map[int32]string{
		0: "USER_REGISTRY",
		1: "DOMAIN_CREATED",
		2: "LDAP_SYNC",
		3: "WXWORK_SYNC",
	}
	CreateType_value = map[string]int32{
		"USER_REGISTRY":  0,
		"DOMAIN_CREATED": 1,
		"LDAP_SYNC":      2,
		"WXWORK_SYNC":    3,
	}
)

func (x CreateType) Enum() *CreateType {
	p := new(CreateType)
	*p = x
	return p
}

func (x CreateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateType) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_user_pb_enum_proto_enumTypes[1].Descriptor()
}

func (CreateType) Type() protoreflect.EnumType {
	return &file_apps_user_pb_enum_proto_enumTypes[1]
}

func (x CreateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreateType.Descriptor instead.
func (CreateType) EnumDescriptor() ([]byte, []int) {
	return file_apps_user_pb_enum_proto_rawDescGZIP(), []int{1}
}

var File_apps_user_pb_enum_proto protoreflect.FileDescriptor

var file_apps_user_pb_enum_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2a, 0x2b, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c,
	0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x2a,
	0x53, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a,
	0x0d, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x59, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x44, 0x41, 0x50, 0x5f, 0x53, 0x59, 0x4e,
	0x43, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x57, 0x58, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x53, 0x59,
	0x4e, 0x43, 0x10, 0x03, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65,
	0x79, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_user_pb_enum_proto_rawDescOnce sync.Once
	file_apps_user_pb_enum_proto_rawDescData = file_apps_user_pb_enum_proto_rawDesc
)

func file_apps_user_pb_enum_proto_rawDescGZIP() []byte {
	file_apps_user_pb_enum_proto_rawDescOnce.Do(func() {
		file_apps_user_pb_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_user_pb_enum_proto_rawDescData)
	})
	return file_apps_user_pb_enum_proto_rawDescData
}

var file_apps_user_pb_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_apps_user_pb_enum_proto_goTypes = []interface{}{
	(Gender)(0),     // 0: infraboard.keyauth.user.Gender
	(CreateType)(0), // 1: infraboard.keyauth.user.CreateType
}
var file_apps_user_pb_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apps_user_pb_enum_proto_init() }
func file_apps_user_pb_enum_proto_init() {
	if File_apps_user_pb_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_user_pb_enum_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_user_pb_enum_proto_goTypes,
		DependencyIndexes: file_apps_user_pb_enum_proto_depIdxs,
		EnumInfos:         file_apps_user_pb_enum_proto_enumTypes,
	}.Build()
	File_apps_user_pb_enum_proto = out.File
	file_apps_user_pb_enum_proto_rawDesc = nil
	file_apps_user_pb_enum_proto_goTypes = nil
	file_apps_user_pb_enum_proto_depIdxs = nil
}
