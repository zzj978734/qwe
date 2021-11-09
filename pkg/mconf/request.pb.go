// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: pkg/mconf/pb/request.proto

package mconf

import (
	page "github.com/infraboard/mcube/pb/page"
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

// CreateGroupRequest 服务组请求
type CreateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 组类型
	// @gotags: json:"type"
	Type Type `protobuf:"varint,1,opt,name=type,proto3,enum=infraboard.keyauth.mconf.Type" json:"type"`
	// 名称
	// @gotags: json:"name" validate:"required,lte=200"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" validate:"required,lte=200"`
	// 描述信息
	// @gotags: json:"description,omitempty"
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateGroupRequest) Reset() {
	*x = CreateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_mconf_pb_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupRequest) ProtoMessage() {}

func (x *CreateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_mconf_pb_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return file_pkg_mconf_pb_request_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGroupRequest) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_GLOBAL
}

func (x *CreateGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateGroupRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// ItemRequest 健值项
type ItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 建的名称
	// @gotags: bson:"_id" json:"key"
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key" bson:"_id"`
	// 关联的组
	// @gotags: bson:"group" json:"group"
	Group string `protobuf:"bytes,2,opt,name=group,proto3" json:"group" bson:"group"`
	// 值对应的值
	// @gotags: bson:"value" json:"value"
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value" bson:"value"`
	// 描述信息
	// @gotags: bson:"description" json:"description,omitempty"
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty" bson:"description"`
}

func (x *ItemRequest) Reset() {
	*x = ItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_mconf_pb_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemRequest) ProtoMessage() {}

func (x *ItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_mconf_pb_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemRequest.ProtoReflect.Descriptor instead.
func (*ItemRequest) Descriptor() ([]byte, []int) {
	return file_pkg_mconf_pb_request_proto_rawDescGZIP(), []int{1}
}

func (x *ItemRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ItemRequest) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *ItemRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ItemRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// QueryGroupRequest 查询组列表
type QueryGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"page"
	Page *page.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 组类型
	// @gotags: json:"type"
	Type Type `protobuf:"varint,2,opt,name=type,proto3,enum=infraboard.keyauth.mconf.Type" json:"type"`
}

func (x *QueryGroupRequest) Reset() {
	*x = QueryGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_mconf_pb_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryGroupRequest) ProtoMessage() {}

func (x *QueryGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_mconf_pb_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryGroupRequest.ProtoReflect.Descriptor instead.
func (*QueryGroupRequest) Descriptor() ([]byte, []int) {
	return file_pkg_mconf_pb_request_proto_rawDescGZIP(), []int{2}
}

func (x *QueryGroupRequest) GetPage() *page.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryGroupRequest) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_GLOBAL
}

// DescribeGroupRequest todo
type DescribeGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
}

func (x *DescribeGroupRequest) Reset() {
	*x = DescribeGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_mconf_pb_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeGroupRequest) ProtoMessage() {}

func (x *DescribeGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_mconf_pb_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeGroupRequest.ProtoReflect.Descriptor instead.
func (*DescribeGroupRequest) Descriptor() ([]byte, []int) {
	return file_pkg_mconf_pb_request_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// DeleteGroupRequest todo
type DeleteGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
}

func (x *DeleteGroupRequest) Reset() {
	*x = DeleteGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_mconf_pb_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupRequest) ProtoMessage() {}

func (x *DeleteGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_mconf_pb_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupRequest.ProtoReflect.Descriptor instead.
func (*DeleteGroupRequest) Descriptor() ([]byte, []int) {
	return file_pkg_mconf_pb_request_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// QueryItemRequest 查询值列表
type QueryItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"page"
	Page *page.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 组名称
	// @gotags: json:"group_name" validate:"required,lte=200"
	GroupName string `protobuf:"bytes,2,opt,name=group_name,json=groupName,proto3" json:"group_name" validate:"required,lte=200"`
}

func (x *QueryItemRequest) Reset() {
	*x = QueryItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_mconf_pb_request_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryItemRequest) ProtoMessage() {}

func (x *QueryItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_mconf_pb_request_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryItemRequest.ProtoReflect.Descriptor instead.
func (*QueryItemRequest) Descriptor() ([]byte, []int) {
	return file_pkg_mconf_pb_request_proto_rawDescGZIP(), []int{5}
}

func (x *QueryItemRequest) GetPage() *page.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryItemRequest) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

// AddItemToGroupRequest todo
type AddItemToGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"group_name"
	GroupName string `protobuf:"bytes,1,opt,name=group_name,json=groupName,proto3" json:"group_name"`
	// @gotags: bson:"items" json:"items"
	Items []*ItemRequest `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *AddItemToGroupRequest) Reset() {
	*x = AddItemToGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_mconf_pb_request_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItemToGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemToGroupRequest) ProtoMessage() {}

func (x *AddItemToGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_mconf_pb_request_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemToGroupRequest.ProtoReflect.Descriptor instead.
func (*AddItemToGroupRequest) Descriptor() ([]byte, []int) {
	return file_pkg_mconf_pb_request_proto_rawDescGZIP(), []int{6}
}

func (x *AddItemToGroupRequest) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *AddItemToGroupRequest) GetItems() []*ItemRequest {
	if x != nil {
		return x.Items
	}
	return nil
}

// RemoveItemFromGroupRequest todo
type RemoveItemFromGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"group_name"
	GroupName string `protobuf:"bytes,1,opt,name=group_name,json=groupName,proto3" json:"group_name"`
	// @gotags: json:"remove_all"
	RemoveAll bool `protobuf:"varint,2,opt,name=remove_all,json=removeAll,proto3" json:"remove_all"`
	// @gotags: bson:"items" json:"items"
	Items []string `protobuf:"bytes,3,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *RemoveItemFromGroupRequest) Reset() {
	*x = RemoveItemFromGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_mconf_pb_request_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveItemFromGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveItemFromGroupRequest) ProtoMessage() {}

func (x *RemoveItemFromGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_mconf_pb_request_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveItemFromGroupRequest.ProtoReflect.Descriptor instead.
func (*RemoveItemFromGroupRequest) Descriptor() ([]byte, []int) {
	return file_pkg_mconf_pb_request_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveItemFromGroupRequest) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *RemoveItemFromGroupRequest) GetRemoveAll() bool {
	if x != nil {
		return x.RemoveAll
	}
	return false
}

func (x *RemoveItemFromGroupRequest) GetItems() []string {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_pkg_mconf_pb_request_proto protoreflect.FileDescriptor

var file_pkg_mconf_pb_request_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x70, 0x62, 0x2f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x1a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x63, 0x6f, 0x6e,
	0x66, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f,
	0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x7e, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x6d, 0x0a, 0x0b, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7f,
	0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x2a, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x12, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x69, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x73, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x6f, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x63, 0x6f,
	0x6e, 0x66, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x70, 0x0a, 0x1a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x61, 0x6c, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x6c,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x63, 0x6f,
	0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_mconf_pb_request_proto_rawDescOnce sync.Once
	file_pkg_mconf_pb_request_proto_rawDescData = file_pkg_mconf_pb_request_proto_rawDesc
)

func file_pkg_mconf_pb_request_proto_rawDescGZIP() []byte {
	file_pkg_mconf_pb_request_proto_rawDescOnce.Do(func() {
		file_pkg_mconf_pb_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_mconf_pb_request_proto_rawDescData)
	})
	return file_pkg_mconf_pb_request_proto_rawDescData
}

var file_pkg_mconf_pb_request_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_mconf_pb_request_proto_goTypes = []interface{}{
	(*CreateGroupRequest)(nil),         // 0: infraboard.keyauth.mconf.CreateGroupRequest
	(*ItemRequest)(nil),                // 1: infraboard.keyauth.mconf.ItemRequest
	(*QueryGroupRequest)(nil),          // 2: infraboard.keyauth.mconf.QueryGroupRequest
	(*DescribeGroupRequest)(nil),       // 3: infraboard.keyauth.mconf.DescribeGroupRequest
	(*DeleteGroupRequest)(nil),         // 4: infraboard.keyauth.mconf.DeleteGroupRequest
	(*QueryItemRequest)(nil),           // 5: infraboard.keyauth.mconf.QueryItemRequest
	(*AddItemToGroupRequest)(nil),      // 6: infraboard.keyauth.mconf.AddItemToGroupRequest
	(*RemoveItemFromGroupRequest)(nil), // 7: infraboard.keyauth.mconf.RemoveItemFromGroupRequest
	(Type)(0),                          // 8: infraboard.keyauth.mconf.Type
	(*page.PageRequest)(nil),           // 9: infraboard.mcube.page.PageRequest
}
var file_pkg_mconf_pb_request_proto_depIdxs = []int32{
	8, // 0: infraboard.keyauth.mconf.CreateGroupRequest.type:type_name -> infraboard.keyauth.mconf.Type
	9, // 1: infraboard.keyauth.mconf.QueryGroupRequest.page:type_name -> infraboard.mcube.page.PageRequest
	8, // 2: infraboard.keyauth.mconf.QueryGroupRequest.type:type_name -> infraboard.keyauth.mconf.Type
	9, // 3: infraboard.keyauth.mconf.QueryItemRequest.page:type_name -> infraboard.mcube.page.PageRequest
	1, // 4: infraboard.keyauth.mconf.AddItemToGroupRequest.items:type_name -> infraboard.keyauth.mconf.ItemRequest
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_mconf_pb_request_proto_init() }
func file_pkg_mconf_pb_request_proto_init() {
	if File_pkg_mconf_pb_request_proto != nil {
		return
	}
	file_pkg_mconf_pb_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_mconf_pb_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupRequest); i {
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
		file_pkg_mconf_pb_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemRequest); i {
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
		file_pkg_mconf_pb_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryGroupRequest); i {
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
		file_pkg_mconf_pb_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeGroupRequest); i {
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
		file_pkg_mconf_pb_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroupRequest); i {
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
		file_pkg_mconf_pb_request_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryItemRequest); i {
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
		file_pkg_mconf_pb_request_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddItemToGroupRequest); i {
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
		file_pkg_mconf_pb_request_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveItemFromGroupRequest); i {
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
			RawDescriptor: file_pkg_mconf_pb_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_mconf_pb_request_proto_goTypes,
		DependencyIndexes: file_pkg_mconf_pb_request_proto_depIdxs,
		MessageInfos:      file_pkg_mconf_pb_request_proto_msgTypes,
	}.Build()
	File_pkg_mconf_pb_request_proto = out.File
	file_pkg_mconf_pb_request_proto_rawDesc = nil
	file_pkg_mconf_pb_request_proto_goTypes = nil
	file_pkg_mconf_pb_request_proto_depIdxs = nil
}
