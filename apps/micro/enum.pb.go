// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: apps/micro/pb/enum.proto

package micro

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

type Type int32

const (
	Type_NULL Type = 0
	// Custom  自定义的服务
	Type_CUSTOM Type = 1
	// BuildIn  系统内建的服务
	Type_BUILD_IN Type = 2
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "NULL",
		1: "CUSTOM",
		2: "BUILD_IN",
	}
	Type_value = map[string]int32{
		"NULL":     0,
		"CUSTOM":   1,
		"BUILD_IN": 2,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_micro_pb_enum_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_apps_micro_pb_enum_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_apps_micro_pb_enum_proto_rawDescGZIP(), []int{0}
}

var File_apps_micro_pb_enum_proto protoreflect.FileDescriptor

var file_apps_micro_pb_enum_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x70, 0x62, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2a, 0x2a, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04,
	0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d,
	0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x49, 0x4e, 0x10, 0x02,
	0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_micro_pb_enum_proto_rawDescOnce sync.Once
	file_apps_micro_pb_enum_proto_rawDescData = file_apps_micro_pb_enum_proto_rawDesc
)

func file_apps_micro_pb_enum_proto_rawDescGZIP() []byte {
	file_apps_micro_pb_enum_proto_rawDescOnce.Do(func() {
		file_apps_micro_pb_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_micro_pb_enum_proto_rawDescData)
	})
	return file_apps_micro_pb_enum_proto_rawDescData
}

var file_apps_micro_pb_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_micro_pb_enum_proto_goTypes = []interface{}{
	(Type)(0), // 0: infraboard.keyauth.micro.Type
}
var file_apps_micro_pb_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apps_micro_pb_enum_proto_init() }
func file_apps_micro_pb_enum_proto_init() {
	if File_apps_micro_pb_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_micro_pb_enum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_micro_pb_enum_proto_goTypes,
		DependencyIndexes: file_apps_micro_pb_enum_proto_depIdxs,
		EnumInfos:         file_apps_micro_pb_enum_proto_enumTypes,
	}.Build()
	File_apps_micro_pb_enum_proto = out.File
	file_apps_micro_pb_enum_proto_rawDesc = nil
	file_apps_micro_pb_enum_proto_goTypes = nil
	file_apps_micro_pb_enum_proto_depIdxs = nil
}
