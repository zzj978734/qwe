// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: common/types/types.proto

package types

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

type UpdateMode int32

const (
	UpdateMode_PUT   UpdateMode = 0
	UpdateMode_PATCH UpdateMode = 1
)

// Enum value maps for UpdateMode.
var (
	UpdateMode_name = map[int32]string{
		0: "PUT",
		1: "PATCH",
	}
	UpdateMode_value = map[string]int32{
		"PUT":   0,
		"PATCH": 1,
	}
)

func (x UpdateMode) Enum() *UpdateMode {
	p := new(UpdateMode)
	*p = x
	return p
}

func (x UpdateMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateMode) Descriptor() protoreflect.EnumDescriptor {
	return file_common_types_types_proto_enumTypes[0].Descriptor()
}

func (UpdateMode) Type() protoreflect.EnumType {
	return &file_common_types_types_proto_enumTypes[0]
}

func (x UpdateMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateMode.Descriptor instead.
func (UpdateMode) EnumDescriptor() ([]byte, []int) {
	return file_common_types_types_proto_rawDescGZIP(), []int{0}
}

var File_common_types_types_proto protoreflect.FileDescriptor

var file_common_types_types_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2a, 0x20, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x55, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x50,
	0x41, 0x54, 0x43, 0x48, 0x10, 0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f,
	0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_types_types_proto_rawDescOnce sync.Once
	file_common_types_types_proto_rawDescData = file_common_types_types_proto_rawDesc
)

func file_common_types_types_proto_rawDescGZIP() []byte {
	file_common_types_types_proto_rawDescOnce.Do(func() {
		file_common_types_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_types_types_proto_rawDescData)
	})
	return file_common_types_types_proto_rawDescData
}

var file_common_types_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_types_types_proto_goTypes = []interface{}{
	(UpdateMode)(0), // 0: infraboard.keyauth.types.UpdateMode
}
var file_common_types_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_types_types_proto_init() }
func file_common_types_types_proto_init() {
	if File_common_types_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_types_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_types_types_proto_goTypes,
		DependencyIndexes: file_common_types_types_proto_depIdxs,
		EnumInfos:         file_common_types_types_proto_enumTypes,
	}.Build()
	File_common_types_types_proto = out.File
	file_common_types_types_proto_rawDesc = nil
	file_common_types_types_proto_goTypes = nil
	file_common_types_types_proto_depIdxs = nil
}
