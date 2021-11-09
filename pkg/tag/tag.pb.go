// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: pkg/tag/pb/tag.proto

package tag

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

// Tag is 标签
type TagKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tag ID
	// @gotags: bson:"_id" json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 创建时间
	// @gotags: bson:"create_at" json:"create_at"
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 更新时间
	// @gotags: bson:"update_at" json:"update_at"
	UpdateAt int64 `protobuf:"varint,3,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// Tag所属的域
	// @gotags: bson:"domain" json:"domain"
	Domain string `protobuf:"bytes,4,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// 创建人
	// @gotags: bson:"creater" json:"creater"
	Creater string `protobuf:"bytes,5,opt,name=creater,proto3" json:"creater" bson:"creater"`
	// tag生效范围
	// @gotags: bson:"scope_type" json:"scope_type"
	ScopeType ScopeType `protobuf:"varint,6,opt,name=scope_type,json=scopeType,proto3,enum=infraboard.keyauth.tag.ScopeType" json:"scope_type" bson:"scope_type"`
	// tag属于哪个namespace
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,7,opt,name=namespace,proto3" json:"namespace" bson:"namespace"`
	// 建名称
	// @gotags: bson:"key_name" json:"key_name"
	KeyName string `protobuf:"bytes,8,opt,name=key_name,json=keyName,proto3" json:"key_name" bson:"key_name"`
	// 建标识
	// @gotags: bson:"key_label" json:"key_label"
	KeyLabel string `protobuf:"bytes,9,opt,name=key_label,json=keyLabel,proto3" json:"key_label" bson:"key_label"`
	// 建描述
	// @gotags: bson:"key_desc" json:"key_desc"
	KeyDesc string `protobuf:"bytes,10,opt,name=key_desc,json=keyDesc,proto3" json:"key_desc" bson:"key_desc"`
	// 值来源
	// @gotags: bson:"value_from" json:"value_from"
	ValueFrom ValueFrom `protobuf:"varint,11,opt,name=value_from,json=valueFrom,proto3,enum=infraboard.keyauth.tag.ValueFrom" json:"value_from" bson:"value_from"`
	// http 获取Tag 值的参数
	// @gotags: bson:"http_from_option" json:"http_from_option,omitempty"
	HttpFromOption *HTTPFromOption `protobuf:"bytes,12,opt,name=http_from_option,json=httpFromOption,proto3" json:"http_from_option,omitempty" bson:"http_from_option"`
	// @gotags: bson:"-" json:"values,omitempty"
	Values []*TagValue `protobuf:"bytes,13,rep,name=values,proto3" json:"values,omitempty" bson:"-"`
}

func (x *TagKey) Reset() {
	*x = TagKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_tag_pb_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagKey) ProtoMessage() {}

func (x *TagKey) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_tag_pb_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagKey.ProtoReflect.Descriptor instead.
func (*TagKey) Descriptor() ([]byte, []int) {
	return file_pkg_tag_pb_tag_proto_rawDescGZIP(), []int{0}
}

func (x *TagKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TagKey) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *TagKey) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *TagKey) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *TagKey) GetCreater() string {
	if x != nil {
		return x.Creater
	}
	return ""
}

func (x *TagKey) GetScopeType() ScopeType {
	if x != nil {
		return x.ScopeType
	}
	return ScopeType_NAMESPACE
}

func (x *TagKey) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *TagKey) GetKeyName() string {
	if x != nil {
		return x.KeyName
	}
	return ""
}

func (x *TagKey) GetKeyLabel() string {
	if x != nil {
		return x.KeyLabel
	}
	return ""
}

func (x *TagKey) GetKeyDesc() string {
	if x != nil {
		return x.KeyDesc
	}
	return ""
}

func (x *TagKey) GetValueFrom() ValueFrom {
	if x != nil {
		return x.ValueFrom
	}
	return ValueFrom_MANUAL
}

func (x *TagKey) GetHttpFromOption() *HTTPFromOption {
	if x != nil {
		return x.HttpFromOption
	}
	return nil
}

func (x *TagKey) GetValues() []*TagValue {
	if x != nil {
		return x.Values
	}
	return nil
}

type TagValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tag Value ID
	// @gotags: bson:"_id" json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 创建时间
	// @gotags: bson:"create_at" json:"create_at"
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 更新时间
	// @gotags: bson:"update_at" json:"update_at"
	UpdateAt int64 `protobuf:"varint,3,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 创建人
	// @gotags: bson:"creater" json:"creater"
	Creater string `protobuf:"bytes,4,opt,name=creater,proto3" json:"creater" bson:"creater"`
	// 关联的Tag key
	// @gotags: bson:"key_id" json:"key_id"
	KeyId string `protobuf:"bytes,5,opt,name=key_id,json=keyId,proto3" json:"key_id" bson:"key_id"`
	// String 类型的值
	// @gotags: bson:"value" json:"value"
	Value *ValueOption `protobuf:"bytes,6,opt,name=value,proto3" json:"value" bson:"value"`
}

func (x *TagValue) Reset() {
	*x = TagValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_tag_pb_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagValue) ProtoMessage() {}

func (x *TagValue) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_tag_pb_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagValue.ProtoReflect.Descriptor instead.
func (*TagValue) Descriptor() ([]byte, []int) {
	return file_pkg_tag_pb_tag_proto_rawDescGZIP(), []int{1}
}

func (x *TagValue) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TagValue) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *TagValue) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *TagValue) GetCreater() string {
	if x != nil {
		return x.Creater
	}
	return ""
}

func (x *TagValue) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

func (x *TagValue) GetValue() *ValueOption {
	if x != nil {
		return x.Value
	}
	return nil
}

type TagKeySet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// @gotags: bson:"items" json:"items"
	Items []*TagKey `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *TagKeySet) Reset() {
	*x = TagKeySet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_tag_pb_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagKeySet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagKeySet) ProtoMessage() {}

func (x *TagKeySet) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_tag_pb_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagKeySet.ProtoReflect.Descriptor instead.
func (*TagKeySet) Descriptor() ([]byte, []int) {
	return file_pkg_tag_pb_tag_proto_rawDescGZIP(), []int{2}
}

func (x *TagKeySet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *TagKeySet) GetItems() []*TagKey {
	if x != nil {
		return x.Items
	}
	return nil
}

type TagValueSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// @gotags: bson:"items" json:"items"
	Items []*TagValue `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *TagValueSet) Reset() {
	*x = TagValueSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_tag_pb_tag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagValueSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagValueSet) ProtoMessage() {}

func (x *TagValueSet) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_tag_pb_tag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagValueSet.ProtoReflect.Descriptor instead.
func (*TagValueSet) Descriptor() ([]byte, []int) {
	return file_pkg_tag_pb_tag_proto_rawDescGZIP(), []int{3}
}

func (x *TagValueSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *TagValueSet) GetItems() []*TagValue {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_pkg_tag_pb_tag_proto protoreflect.FileDescriptor

var file_pkg_tag_pb_tag_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x61, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x61, 0x67, 0x1a, 0x15,
	0x70, 0x6b, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x70,
	0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x85, 0x04, 0x0a, 0x06, 0x54, 0x61, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x0a, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x74, 0x61, 0x67, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x19,
	0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6b, 0x65, 0x79, 0x44, 0x65, 0x73, 0x63, 0x12, 0x40, 0x0a, 0x0a, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x72, 0x6f, 0x6d,
	0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x50, 0x0a, 0x10, 0x68,
	0x74, 0x74, 0x70, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x48,
	0x54, 0x54, 0x50, 0x46, 0x72, 0x6f, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x68,
	0x74, 0x74, 0x70, 0x46, 0x72, 0x6f, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x54, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0xc0, 0x01, 0x0a, 0x08, 0x54, 0x61, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x12,
	0x39, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x57, 0x0a, 0x09, 0x54, 0x61,
	0x67, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x34, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x54, 0x61, 0x67, 0x4b, 0x65, 0x79, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x5b, 0x0a, 0x0b, 0x54, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x36, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x61, 0x67,
	0x2e, 0x54, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_tag_pb_tag_proto_rawDescOnce sync.Once
	file_pkg_tag_pb_tag_proto_rawDescData = file_pkg_tag_pb_tag_proto_rawDesc
)

func file_pkg_tag_pb_tag_proto_rawDescGZIP() []byte {
	file_pkg_tag_pb_tag_proto_rawDescOnce.Do(func() {
		file_pkg_tag_pb_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_tag_pb_tag_proto_rawDescData)
	})
	return file_pkg_tag_pb_tag_proto_rawDescData
}

var file_pkg_tag_pb_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_tag_pb_tag_proto_goTypes = []interface{}{
	(*TagKey)(nil),         // 0: infraboard.keyauth.tag.TagKey
	(*TagValue)(nil),       // 1: infraboard.keyauth.tag.TagValue
	(*TagKeySet)(nil),      // 2: infraboard.keyauth.tag.TagKeySet
	(*TagValueSet)(nil),    // 3: infraboard.keyauth.tag.TagValueSet
	(ScopeType)(0),         // 4: infraboard.keyauth.tag.ScopeType
	(ValueFrom)(0),         // 5: infraboard.keyauth.tag.ValueFrom
	(*HTTPFromOption)(nil), // 6: infraboard.keyauth.tag.HTTPFromOption
	(*ValueOption)(nil),    // 7: infraboard.keyauth.tag.ValueOption
}
var file_pkg_tag_pb_tag_proto_depIdxs = []int32{
	4, // 0: infraboard.keyauth.tag.TagKey.scope_type:type_name -> infraboard.keyauth.tag.ScopeType
	5, // 1: infraboard.keyauth.tag.TagKey.value_from:type_name -> infraboard.keyauth.tag.ValueFrom
	6, // 2: infraboard.keyauth.tag.TagKey.http_from_option:type_name -> infraboard.keyauth.tag.HTTPFromOption
	1, // 3: infraboard.keyauth.tag.TagKey.values:type_name -> infraboard.keyauth.tag.TagValue
	7, // 4: infraboard.keyauth.tag.TagValue.value:type_name -> infraboard.keyauth.tag.ValueOption
	0, // 5: infraboard.keyauth.tag.TagKeySet.items:type_name -> infraboard.keyauth.tag.TagKey
	1, // 6: infraboard.keyauth.tag.TagValueSet.items:type_name -> infraboard.keyauth.tag.TagValue
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_pkg_tag_pb_tag_proto_init() }
func file_pkg_tag_pb_tag_proto_init() {
	if File_pkg_tag_pb_tag_proto != nil {
		return
	}
	file_pkg_tag_pb_enum_proto_init()
	file_pkg_tag_pb_request_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_tag_pb_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagKey); i {
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
		file_pkg_tag_pb_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagValue); i {
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
		file_pkg_tag_pb_tag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagKeySet); i {
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
		file_pkg_tag_pb_tag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagValueSet); i {
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
			RawDescriptor: file_pkg_tag_pb_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_tag_pb_tag_proto_goTypes,
		DependencyIndexes: file_pkg_tag_pb_tag_proto_depIdxs,
		MessageInfos:      file_pkg_tag_pb_tag_proto_msgTypes,
	}.Build()
	File_pkg_tag_pb_tag_proto = out.File
	file_pkg_tag_pb_tag_proto_rawDesc = nil
	file_pkg_tag_pb_tag_proto_goTypes = nil
	file_pkg_tag_pb_tag_proto_depIdxs = nil
}
