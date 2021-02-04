//go:generate  mcube enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/policy/pb/request.proto

package policy

import (
	proto "github.com/golang/protobuf/proto"
	page "github.com/infraboard/mcube/pb/page"
	_ "github.com/infraboard/protoc-gen-go-ext/extension/tag"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// CreatePolicyRequest 创建策略的请求
type CreatePolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 范围
	NamespaceId string `protobuf:"bytes,1,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id" validate:"lte=120" bson:"namespace_id"`
	// 用户
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account" validate:"required,lte=120" bson:"account"`
	// 角色名称
	RoleId string `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id" bson:"role_id" validate:"required,lte=40"`
	// 范围控制
	Scope string `protobuf:"bytes,4,opt,name=scope,proto3" json:"scope" bson:"scope"`
	// 策略过期时间
	ExpiredTime int64 `protobuf:"varint,5,opt,name=expired_time,json=expiredTime,proto3" json:"expired_time" bson:"expired_time"`
	// 策略的类型
	Type PolicyType `protobuf:"varint,6,opt,name=type,proto3,enum=keyauth.policy.PolicyType" json:"type" bson:"type"`
}

func (x *CreatePolicyRequest) Reset() {
	*x = CreatePolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_policy_pb_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePolicyRequest) ProtoMessage() {}

func (x *CreatePolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_policy_pb_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePolicyRequest.ProtoReflect.Descriptor instead.
func (*CreatePolicyRequest) Descriptor() ([]byte, []int) {
	return file_pkg_policy_pb_request_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePolicyRequest) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

func (x *CreatePolicyRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *CreatePolicyRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *CreatePolicyRequest) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *CreatePolicyRequest) GetExpiredTime() int64 {
	if x != nil {
		return x.ExpiredTime
	}
	return 0
}

func (x *CreatePolicyRequest) GetType() PolicyType {
	if x != nil {
		return x.Type
	}
	return PolicyType_NULL
}

// QueryPolicyRequest 获取子账号列表
type QueryPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page          *page.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page" bson:"page"`
	Account       string            `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	RoleId        string            `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	NamespaceId   string            `protobuf:"bytes,4,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	Type          PolicyType        `protobuf:"varint,5,opt,name=type,proto3,enum=keyauth.policy.PolicyType" json:"type,omitempty"`
	WithRole      bool              `protobuf:"varint,6,opt,name=with_role,json=withRole,proto3" json:"with_role,omitempty"`
	WithNamespace bool              `protobuf:"varint,7,opt,name=with_namespace,json=withNamespace,proto3" json:"with_namespace,omitempty"`
}

func (x *QueryPolicyRequest) Reset() {
	*x = QueryPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_policy_pb_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPolicyRequest) ProtoMessage() {}

func (x *QueryPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_policy_pb_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPolicyRequest.ProtoReflect.Descriptor instead.
func (*QueryPolicyRequest) Descriptor() ([]byte, []int) {
	return file_pkg_policy_pb_request_proto_rawDescGZIP(), []int{1}
}

func (x *QueryPolicyRequest) GetPage() *page.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryPolicyRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *QueryPolicyRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *QueryPolicyRequest) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

func (x *QueryPolicyRequest) GetType() PolicyType {
	if x != nil {
		return x.Type
	}
	return PolicyType_NULL
}

func (x *QueryPolicyRequest) GetWithRole() bool {
	if x != nil {
		return x.WithRole
	}
	return false
}

func (x *QueryPolicyRequest) GetWithNamespace() bool {
	if x != nil {
		return x.WithNamespace
	}
	return false
}

// DescribePolicyRequest todo
type DescribePolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DescribePolicyRequest) Reset() {
	*x = DescribePolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_policy_pb_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribePolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePolicyRequest) ProtoMessage() {}

func (x *DescribePolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_policy_pb_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePolicyRequest.ProtoReflect.Descriptor instead.
func (*DescribePolicyRequest) Descriptor() ([]byte, []int) {
	return file_pkg_policy_pb_request_proto_rawDescGZIP(), []int{2}
}

func (x *DescribePolicyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// DeletePolicyRequest todo
type DeletePolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Account     string     `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	RoleId      string     `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	NamespaceId string     `protobuf:"bytes,4,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	Type        PolicyType `protobuf:"varint,5,opt,name=type,proto3,enum=keyauth.policy.PolicyType" json:"type,omitempty"`
}

func (x *DeletePolicyRequest) Reset() {
	*x = DeletePolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_policy_pb_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePolicyRequest) ProtoMessage() {}

func (x *DeletePolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_policy_pb_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePolicyRequest.ProtoReflect.Descriptor instead.
func (*DeletePolicyRequest) Descriptor() ([]byte, []int) {
	return file_pkg_policy_pb_request_proto_rawDescGZIP(), []int{3}
}

func (x *DeletePolicyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeletePolicyRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *DeletePolicyRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *DeletePolicyRequest) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

func (x *DeletePolicyRequest) GetType() PolicyType {
	if x != nil {
		return x.Type
	}
	return PolicyType_NULL
}

var File_pkg_policy_pb_request_proto protoreflect.FileDescriptor

var file_pkg_policy_pb_request_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x70, 0x62, 0x2f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6b,
	0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x18, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x65, 0x78, 0x74,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x74,
	0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f,
	0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x04, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x63, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x40, 0xc2, 0xde, 0x1f, 0x3c, 0x0a, 0x3a, 0x62, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22,
	0x6c, 0x74, 0x65, 0x3d, 0x31, 0x32, 0x30, 0x22, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x59, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3f, 0xc2, 0xde, 0x1f, 0x3b, 0x0a, 0x39, 0x62, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2c, 0x6c,
	0x74, 0x65, 0x3d, 0x31, 0x32, 0x30, 0x22, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x57, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x3e, 0xc2, 0xde, 0x1f, 0x3a, 0x0a, 0x38, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72,
	0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x6f,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a,
	0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2c, 0x6c, 0x74, 0x65, 0x3d, 0x34, 0x30,
	0x22, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x05, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1f, 0xc2, 0xde, 0x1f, 0x1b, 0x0a, 0x19,
	0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x22, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x12, 0x50, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x42, 0x2d, 0xc2, 0xde, 0x1f, 0x29, 0x0a, 0x27, 0x62, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x22, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x4d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1a, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x42, 0x1d, 0xc2, 0xde,
	0x1f, 0x19, 0x0a, 0x17, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x20,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0xef, 0x03, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x1d, 0xc2, 0xde, 0x1f, 0x19, 0x0a,
	0x17, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x61, 0x67, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x70, 0x61, 0x67, 0x65, 0x22, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x38,
	0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x1e, 0xc2, 0xde, 0x1f, 0x1a, 0x0a, 0x18, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xc2, 0xde, 0x1f, 0x1a, 0x0a,
	0x18, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x2c, 0x6f,
	0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x46, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0xc2, 0xde, 0x1f, 0x1f, 0x0a, 0x1d, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x0b, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x42, 0x1b, 0xc2, 0xde, 0x1f, 0x17, 0x0a, 0x15, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x74, 0x79, 0x70, 0x65, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x42, 0x20, 0xc2, 0xde, 0x1f, 0x1c, 0x0a,
	0x1a, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x72, 0x6f, 0x6c, 0x65,
	0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x08, 0x77, 0x69, 0x74,
	0x68, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x4c, 0x0a, 0x0e, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x42, 0x25, 0xc2,
	0xde, 0x1f, 0x21, 0x0a, 0x1f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x77, 0x69, 0x74, 0x68, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x52, 0x0d, 0x77, 0x69, 0x74, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x22, 0x38, 0x0a, 0x15, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xc2, 0xde, 0x1f, 0x0b, 0x0a, 0x09,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x22, 0xbe, 0x02,
	0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0f, 0xc2, 0xde, 0x1f, 0x0b, 0x0a, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69,
	0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xc2, 0xde, 0x1f, 0x1a, 0x0a, 0x18, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2c, 0x6f, 0x6d, 0x69,
	0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x37, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x1e, 0xc2, 0xde, 0x1f, 0x1a, 0x0a, 0x18, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72,
	0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x0c, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x23, 0xc2, 0xde, 0x1f, 0x1f, 0x0a, 0x1d, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x4b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1a, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x42, 0x1b, 0xc2, 0xde, 0x1f,
	0x17, 0x0a, 0x15, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x2c, 0x6f, 0x6d,
	0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x2a,
	0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_policy_pb_request_proto_rawDescOnce sync.Once
	file_pkg_policy_pb_request_proto_rawDescData = file_pkg_policy_pb_request_proto_rawDesc
)

func file_pkg_policy_pb_request_proto_rawDescGZIP() []byte {
	file_pkg_policy_pb_request_proto_rawDescOnce.Do(func() {
		file_pkg_policy_pb_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_policy_pb_request_proto_rawDescData)
	})
	return file_pkg_policy_pb_request_proto_rawDescData
}

var file_pkg_policy_pb_request_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_policy_pb_request_proto_goTypes = []interface{}{
	(*CreatePolicyRequest)(nil),   // 0: keyauth.policy.CreatePolicyRequest
	(*QueryPolicyRequest)(nil),    // 1: keyauth.policy.QueryPolicyRequest
	(*DescribePolicyRequest)(nil), // 2: keyauth.policy.DescribePolicyRequest
	(*DeletePolicyRequest)(nil),   // 3: keyauth.policy.DeletePolicyRequest
	(PolicyType)(0),               // 4: keyauth.policy.PolicyType
	(*page.PageRequest)(nil),      // 5: page.PageRequest
}
var file_pkg_policy_pb_request_proto_depIdxs = []int32{
	4, // 0: keyauth.policy.CreatePolicyRequest.type:type_name -> keyauth.policy.PolicyType
	5, // 1: keyauth.policy.QueryPolicyRequest.page:type_name -> page.PageRequest
	4, // 2: keyauth.policy.QueryPolicyRequest.type:type_name -> keyauth.policy.PolicyType
	4, // 3: keyauth.policy.DeletePolicyRequest.type:type_name -> keyauth.policy.PolicyType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_policy_pb_request_proto_init() }
func file_pkg_policy_pb_request_proto_init() {
	if File_pkg_policy_pb_request_proto != nil {
		return
	}
	file_pkg_policy_pb_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_policy_pb_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePolicyRequest); i {
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
		file_pkg_policy_pb_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPolicyRequest); i {
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
		file_pkg_policy_pb_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribePolicyRequest); i {
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
		file_pkg_policy_pb_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePolicyRequest); i {
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
			RawDescriptor: file_pkg_policy_pb_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_policy_pb_request_proto_goTypes,
		DependencyIndexes: file_pkg_policy_pb_request_proto_depIdxs,
		MessageInfos:      file_pkg_policy_pb_request_proto_msgTypes,
	}.Build()
	File_pkg_policy_pb_request_proto = out.File
	file_pkg_policy_pb_request_proto_rawDesc = nil
	file_pkg_policy_pb_request_proto_goTypes = nil
	file_pkg_policy_pb_request_proto_depIdxs = nil
}
