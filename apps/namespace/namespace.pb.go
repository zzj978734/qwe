// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: apps/namespace/pb/namespace.proto

package namespace

import (
	department "github.com/infraboard/keyauth/apps/department"
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

type Visible int32

const (
	// 默认空间是私有的
	Visible_PRIVATE Visible = 0
	// PUBLIC  公开的空间
	Visible_PUBLIC Visible = 1
)

// Enum value maps for Visible.
var (
	Visible_name = map[int32]string{
		0: "PRIVATE",
		1: "PUBLIC",
	}
	Visible_value = map[string]int32{
		"PRIVATE": 0,
		"PUBLIC":  1,
	}
)

func (x Visible) Enum() *Visible {
	p := new(Visible)
	*p = x
	return p
}

func (x Visible) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Visible) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_namespace_pb_namespace_proto_enumTypes[0].Descriptor()
}

func (Visible) Type() protoreflect.EnumType {
	return &file_apps_namespace_pb_namespace_proto_enumTypes[0]
}

func (x Visible) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Visible.Descriptor instead.
func (Visible) EnumDescriptor() ([]byte, []int) {
	return file_apps_namespace_pb_namespace_proto_rawDescGZIP(), []int{0}
}

// Namespace tenant resource container
type Namespace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 项目唯一ID
	// @gotags: bson:"_id" json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 所属域ID
	// @gotags: bson:"domain" json:"domain"
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// 创建人
	// @gotags: bson:"create_by" json:"create_by"
	CreateBy string `protobuf:"bytes,3,opt,name=create_by,json=createBy,proto3" json:"create_by" bson:"create_by"`
	// 创建时间
	// @gotags: bson:"create_at" json:"create_at"
	CreateAt int64 `protobuf:"varint,4,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 项目修改时间
	// @gotags: bson:"update_at" json:"update_at"
	UpdateAt int64 `protobuf:"varint,5,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 禁用项目, 该项目所有人暂时都无法访问
	// @gotags: bson:"enabled" json:"enabled"
	Enabled bool `protobuf:"varint,6,opt,name=enabled,proto3" json:"enabled" bson:"enabled"`
	// 所属部门
	// @gotags: bson:"department_id" json:"department_id"
	DepartmentId string `protobuf:"bytes,7,opt,name=department_id,json=departmentId,proto3" json:"department_id" bson:"department_id"`
	// 项目名称
	// @gotags: bson:"name" json:"name"'
	Name string `protobuf:"bytes,8,opt,name=name,proto3" json:"name" bson:"name"`
	// 项目描述图片
	// @gotags: bson:"picture" json:"picture"
	Picture string `protobuf:"bytes,9,opt,name=picture,proto3" json:"picture" bson:"picture"`
	// 项目所有者, PMO
	// @gotags: bson:"owner" json:"owner"
	Owner string `protobuf:"bytes,10,opt,name=owner,proto3" json:"owner" bson:"owner"`
	// 项目描述
	// @gotags: bson:"description" json:"description"
	Description string `protobuf:"bytes,11,opt,name=description,proto3" json:"description" bson:"description"`
	// 补充的部门
	// @gotags: bson:"-" json:"department,omitempty"
	Department *department.Department `protobuf:"bytes,12,opt,name=department,proto3" json:"department,omitempty" bson:"-"`
	// 空间可见性, 默认是私有空间
	// @gotags: bson:"-" bson:"visible" json:"visible"
	Visible Visible `protobuf:"varint,13,opt,name=visible,proto3,enum=infraboard.keyauth.namespace.Visible" json:"visible" bson:"-" bson:"visible"`
}

func (x *Namespace) Reset() {
	*x = Namespace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_namespace_pb_namespace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Namespace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Namespace) ProtoMessage() {}

func (x *Namespace) ProtoReflect() protoreflect.Message {
	mi := &file_apps_namespace_pb_namespace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Namespace.ProtoReflect.Descriptor instead.
func (*Namespace) Descriptor() ([]byte, []int) {
	return file_apps_namespace_pb_namespace_proto_rawDescGZIP(), []int{0}
}

func (x *Namespace) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Namespace) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Namespace) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *Namespace) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Namespace) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Namespace) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Namespace) GetDepartmentId() string {
	if x != nil {
		return x.DepartmentId
	}
	return ""
}

func (x *Namespace) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Namespace) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *Namespace) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Namespace) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Namespace) GetDepartment() *department.Department {
	if x != nil {
		return x.Department
	}
	return nil
}

func (x *Namespace) GetVisible() Visible {
	if x != nil {
		return x.Visible
	}
	return Visible_PRIVATE
}

type Set struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// @gotags: bson:"items" json:"items"
	Items []*Namespace `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *Set) Reset() {
	*x = Set{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_namespace_pb_namespace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set) ProtoMessage() {}

func (x *Set) ProtoReflect() protoreflect.Message {
	mi := &file_apps_namespace_pb_namespace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Set.ProtoReflect.Descriptor instead.
func (*Set) Descriptor() ([]byte, []int) {
	return file_apps_namespace_pb_namespace_proto_rawDescGZIP(), []int{1}
}

func (x *Set) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Set) GetItems() []*Namespace {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_apps_namespace_pb_namespace_proto protoreflect.FileDescriptor

var file_apps_namespace_pb_namespace_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x2f, 0x70, 0x62, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x1a, 0x23, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x03, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x3f, 0x0a, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x52, 0x07, 0x76, 0x69, 0x73,
	0x69, 0x62, 0x6c, 0x65, 0x22, 0x5a, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x3d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65,
	0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x2a, 0x22, 0x0a, 0x07, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50,
	0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c,
	0x49, 0x43, 0x10, 0x01, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65,
	0x79, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_namespace_pb_namespace_proto_rawDescOnce sync.Once
	file_apps_namespace_pb_namespace_proto_rawDescData = file_apps_namespace_pb_namespace_proto_rawDesc
)

func file_apps_namespace_pb_namespace_proto_rawDescGZIP() []byte {
	file_apps_namespace_pb_namespace_proto_rawDescOnce.Do(func() {
		file_apps_namespace_pb_namespace_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_namespace_pb_namespace_proto_rawDescData)
	})
	return file_apps_namespace_pb_namespace_proto_rawDescData
}

var file_apps_namespace_pb_namespace_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_namespace_pb_namespace_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_apps_namespace_pb_namespace_proto_goTypes = []interface{}{
	(Visible)(0),                  // 0: infraboard.keyauth.namespace.Visible
	(*Namespace)(nil),             // 1: infraboard.keyauth.namespace.Namespace
	(*Set)(nil),                   // 2: infraboard.keyauth.namespace.Set
	(*department.Department)(nil), // 3: infraboard.keyauth.department.Department
}
var file_apps_namespace_pb_namespace_proto_depIdxs = []int32{
	3, // 0: infraboard.keyauth.namespace.Namespace.department:type_name -> infraboard.keyauth.department.Department
	0, // 1: infraboard.keyauth.namespace.Namespace.visible:type_name -> infraboard.keyauth.namespace.Visible
	1, // 2: infraboard.keyauth.namespace.Set.items:type_name -> infraboard.keyauth.namespace.Namespace
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_apps_namespace_pb_namespace_proto_init() }
func file_apps_namespace_pb_namespace_proto_init() {
	if File_apps_namespace_pb_namespace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_namespace_pb_namespace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Namespace); i {
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
		file_apps_namespace_pb_namespace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Set); i {
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
			RawDescriptor: file_apps_namespace_pb_namespace_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_namespace_pb_namespace_proto_goTypes,
		DependencyIndexes: file_apps_namespace_pb_namespace_proto_depIdxs,
		EnumInfos:         file_apps_namespace_pb_namespace_proto_enumTypes,
		MessageInfos:      file_apps_namespace_pb_namespace_proto_msgTypes,
	}.Build()
	File_apps_namespace_pb_namespace_proto = out.File
	file_apps_namespace_pb_namespace_proto_rawDesc = nil
	file_apps_namespace_pb_namespace_proto_goTypes = nil
	file_apps_namespace_pb_namespace_proto_depIdxs = nil
}
