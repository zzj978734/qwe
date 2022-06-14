// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: apps/department/pb/request.proto

package department

import (
	types "github.com/infraboard/keyauth/common/types"
	request "github.com/infraboard/mcube/http/request"
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

// CreateDepartmentRequest 创建部门请求
type CreateDepartmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 部门名称
	// @gotags: json:"name" validate:"required,lte=60"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" validate:"required,lte=60"`
	// 显示名称
	// @gotags: json:"display_name"
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name"`
	// 上级部门ID
	// @gotags: json:"parent_id" validate:"lte=200"
	ParentId string `protobuf:"bytes,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id" validate:"lte=200"`
	// 部门管理者account
	// @gotags: json:"manager" validate:"required,lte=200"
	Manager string `protobuf:"bytes,4,opt,name=manager,proto3" json:"manager" validate:"required,lte=200"`
	// 部门成员默认角色
	// @gotags: json:"default_role_id" validate:"lte=200"
	DefaultRoleId string `protobuf:"bytes,5,opt,name=default_role_id,json=defaultRoleId,proto3" json:"default_role_id" validate:"lte=200"`
	// 部门所属域
	// @gotags: json:"domain"
	Domain string `protobuf:"bytes,6,opt,name=domain,proto3" json:"domain"`
	// 创建人
	// @gotags: json:"create_by"
	CreateBy string `protobuf:"bytes,7,opt,name=create_by,json=createBy,proto3" json:"create_by"`
}

func (x *CreateDepartmentRequest) Reset() {
	*x = CreateDepartmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_department_pb_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDepartmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDepartmentRequest) ProtoMessage() {}

func (x *CreateDepartmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_department_pb_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDepartmentRequest.ProtoReflect.Descriptor instead.
func (*CreateDepartmentRequest) Descriptor() ([]byte, []int) {
	return file_apps_department_pb_request_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDepartmentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDepartmentRequest) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *CreateDepartmentRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *CreateDepartmentRequest) GetManager() string {
	if x != nil {
		return x.Manager
	}
	return ""
}

func (x *CreateDepartmentRequest) GetDefaultRoleId() string {
	if x != nil {
		return x.DefaultRoleId
	}
	return ""
}

func (x *CreateDepartmentRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CreateDepartmentRequest) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

// QueryDepartmentRequest todo
type QueryDepartmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// @gotags: json:"parent_id"
	ParentId string `protobuf:"bytes,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id"`
	// @gotags: json:"keywords"
	Keywords string `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords"`
	// @gotags: json:"skip_items"
	SkipItems bool `protobuf:"varint,4,opt,name=skip_items,json=skipItems,proto3" json:"skip_items"`
	// @gotags: json:"with_sub_count"
	WithSubCount bool `protobuf:"varint,5,opt,name=with_sub_count,json=withSubCount,proto3" json:"with_sub_count"`
	// @gotags: json:"with_user_count"
	WithUserCount bool `protobuf:"varint,6,opt,name=with_user_count,json=withUserCount,proto3" json:"with_user_count"`
	// @gotags: json:"with_role"
	WithRole bool `protobuf:"varint,7,opt,name=with_role,json=withRole,proto3" json:"with_role"`
	// @gotags: json:"domain" validate:"required"
	Domain string `protobuf:"bytes,8,opt,name=domain,proto3" json:"domain" validate:"required"`
}

func (x *QueryDepartmentRequest) Reset() {
	*x = QueryDepartmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_department_pb_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDepartmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDepartmentRequest) ProtoMessage() {}

func (x *QueryDepartmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_department_pb_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryDepartmentRequest.ProtoReflect.Descriptor instead.
func (*QueryDepartmentRequest) Descriptor() ([]byte, []int) {
	return file_apps_department_pb_request_proto_rawDescGZIP(), []int{1}
}

func (x *QueryDepartmentRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryDepartmentRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *QueryDepartmentRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *QueryDepartmentRequest) GetSkipItems() bool {
	if x != nil {
		return x.SkipItems
	}
	return false
}

func (x *QueryDepartmentRequest) GetWithSubCount() bool {
	if x != nil {
		return x.WithSubCount
	}
	return false
}

func (x *QueryDepartmentRequest) GetWithUserCount() bool {
	if x != nil {
		return x.WithUserCount
	}
	return false
}

func (x *QueryDepartmentRequest) GetWithRole() bool {
	if x != nil {
		return x.WithRole
	}
	return false
}

func (x *QueryDepartmentRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

// DescribeDeparmentRequest 详情查询
type DescribeDeparmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// @gotags: json:"name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	// @gotags: json:"with_sub_count"
	WithSubCount bool `protobuf:"varint,3,opt,name=with_sub_count,json=withSubCount,proto3" json:"with_sub_count"`
	// @gotags: json:"with_user_count"
	WithUserCount bool `protobuf:"varint,4,opt,name=with_user_count,json=withUserCount,proto3" json:"with_user_count"`
	// @gotags: json:"with_role"
	WithRole bool `protobuf:"varint,5,opt,name=with_role,json=withRole,proto3" json:"with_role"`
	// @gotags: json:"domain" validate:"required"
	Domain string `protobuf:"bytes,6,opt,name=domain,proto3" json:"domain" validate:"required"`
}

func (x *DescribeDeparmentRequest) Reset() {
	*x = DescribeDeparmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_department_pb_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeDeparmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeDeparmentRequest) ProtoMessage() {}

func (x *DescribeDeparmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_department_pb_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeDeparmentRequest.ProtoReflect.Descriptor instead.
func (*DescribeDeparmentRequest) Descriptor() ([]byte, []int) {
	return file_apps_department_pb_request_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeDeparmentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DescribeDeparmentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DescribeDeparmentRequest) GetWithSubCount() bool {
	if x != nil {
		return x.WithSubCount
	}
	return false
}

func (x *DescribeDeparmentRequest) GetWithUserCount() bool {
	if x != nil {
		return x.WithUserCount
	}
	return false
}

func (x *DescribeDeparmentRequest) GetWithRole() bool {
	if x != nil {
		return x.WithRole
	}
	return false
}

func (x *DescribeDeparmentRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

// UpdateDepartmentRequest todo
type UpdateDepartmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 更新模式
	// @gotags: json:"update_mode"
	UpdateMode types.UpdateMode `protobuf:"varint,1,opt,name=update_mode,json=updateMode,proto3,enum=infraboard.keyauth.types.UpdateMode" json:"update_mode"`
	// @gotags: json:"id"
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// @gotags: json:"data"
	Data *CreateDepartmentRequest `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *UpdateDepartmentRequest) Reset() {
	*x = UpdateDepartmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_department_pb_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDepartmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDepartmentRequest) ProtoMessage() {}

func (x *UpdateDepartmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_department_pb_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDepartmentRequest.ProtoReflect.Descriptor instead.
func (*UpdateDepartmentRequest) Descriptor() ([]byte, []int) {
	return file_apps_department_pb_request_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateDepartmentRequest) GetUpdateMode() types.UpdateMode {
	if x != nil {
		return x.UpdateMode
	}
	return types.UpdateMode(0)
}

func (x *UpdateDepartmentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateDepartmentRequest) GetData() *CreateDepartmentRequest {
	if x != nil {
		return x.Data
	}
	return nil
}

// DeleteDepartmentRequest todo
type DeleteDepartmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteDepartmentRequest) Reset() {
	*x = DeleteDepartmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_department_pb_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDepartmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDepartmentRequest) ProtoMessage() {}

func (x *DeleteDepartmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_department_pb_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDepartmentRequest.ProtoReflect.Descriptor instead.
func (*DeleteDepartmentRequest) Descriptor() ([]byte, []int) {
	return file_apps_department_pb_request_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteDepartmentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// JoinDepartmentRequest todo
type JoinDepartmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 申请人
	// @gotags: json:"account" validate:"required"
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account" validate:"required"`
	// 申请加入的部门
	// @gotags: json:"department_id" validate:"required"
	DepartmentId string `protobuf:"bytes,2,opt,name=department_id,json=departmentId,proto3" json:"department_id" validate:"required"`
	// 留言
	// @gotags: json:"message"
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message"`
	// 所属域
	// @gotags: json:"domain"
	Domain string `protobuf:"bytes,4,opt,name=domain,proto3" json:"domain"`
}

func (x *JoinDepartmentRequest) Reset() {
	*x = JoinDepartmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_department_pb_request_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinDepartmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinDepartmentRequest) ProtoMessage() {}

func (x *JoinDepartmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_department_pb_request_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinDepartmentRequest.ProtoReflect.Descriptor instead.
func (*JoinDepartmentRequest) Descriptor() ([]byte, []int) {
	return file_apps_department_pb_request_proto_rawDescGZIP(), []int{5}
}

func (x *JoinDepartmentRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *JoinDepartmentRequest) GetDepartmentId() string {
	if x != nil {
		return x.DepartmentId
	}
	return ""
}

func (x *JoinDepartmentRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *JoinDepartmentRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

// QueryApplicationFormRequet todo
type QueryApplicationFormRequet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// @gotags: json:"account"
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account"`
	// @gotags: json:"department_id"
	DepartmentId string `protobuf:"bytes,3,opt,name=department_id,json=departmentId,proto3" json:"department_id"`
	// @gotags: json:"status"
	Status ApplicationFormStatus `protobuf:"varint,4,opt,name=status,proto3,enum=infraboard.keyauth.department.ApplicationFormStatus" json:"status"`
	// @gotags: json:"skip_items"
	SkipItems bool `protobuf:"varint,5,opt,name=skip_items,json=skipItems,proto3" json:"skip_items"`
	// @gotags: json:"domain" validate:"required"
	Domain string `protobuf:"bytes,6,opt,name=domain,proto3" json:"domain" validate:"required"`
}

func (x *QueryApplicationFormRequet) Reset() {
	*x = QueryApplicationFormRequet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_department_pb_request_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryApplicationFormRequet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryApplicationFormRequet) ProtoMessage() {}

func (x *QueryApplicationFormRequet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_department_pb_request_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryApplicationFormRequet.ProtoReflect.Descriptor instead.
func (*QueryApplicationFormRequet) Descriptor() ([]byte, []int) {
	return file_apps_department_pb_request_proto_rawDescGZIP(), []int{6}
}

func (x *QueryApplicationFormRequet) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryApplicationFormRequet) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *QueryApplicationFormRequet) GetDepartmentId() string {
	if x != nil {
		return x.DepartmentId
	}
	return ""
}

func (x *QueryApplicationFormRequet) GetStatus() ApplicationFormStatus {
	if x != nil {
		return x.Status
	}
	return ApplicationFormStatus_NULL
}

func (x *QueryApplicationFormRequet) GetSkipItems() bool {
	if x != nil {
		return x.SkipItems
	}
	return false
}

func (x *QueryApplicationFormRequet) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

// DescribeApplicationFormRequet todo
type DescribeApplicationFormRequet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// @gotags: json:"domain"
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain"`
}

func (x *DescribeApplicationFormRequet) Reset() {
	*x = DescribeApplicationFormRequet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_department_pb_request_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeApplicationFormRequet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeApplicationFormRequet) ProtoMessage() {}

func (x *DescribeApplicationFormRequet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_department_pb_request_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeApplicationFormRequet.ProtoReflect.Descriptor instead.
func (*DescribeApplicationFormRequet) Descriptor() ([]byte, []int) {
	return file_apps_department_pb_request_proto_rawDescGZIP(), []int{7}
}

func (x *DescribeApplicationFormRequet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DescribeApplicationFormRequet) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

// DealApplicationFormRequest todo
type DealApplicationFormRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 状态
	// @gotags: json:"status"
	Status ApplicationFormStatus `protobuf:"varint,2,opt,name=status,proto3,enum=infraboard.keyauth.department.ApplicationFormStatus" json:"status"`
	// 备注
	// @gotags: json:"message"
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message"`
}

func (x *DealApplicationFormRequest) Reset() {
	*x = DealApplicationFormRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_department_pb_request_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DealApplicationFormRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DealApplicationFormRequest) ProtoMessage() {}

func (x *DealApplicationFormRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_department_pb_request_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DealApplicationFormRequest.ProtoReflect.Descriptor instead.
func (*DealApplicationFormRequest) Descriptor() ([]byte, []int) {
	return file_apps_department_pb_request_proto_rawDescGZIP(), []int{8}
}

func (x *DealApplicationFormRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DealApplicationFormRequest) GetStatus() ApplicationFormStatus {
	if x != nil {
		return x.Status
	}
	return ApplicationFormStatus_NULL
}

func (x *DealApplicationFormRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_apps_department_pb_request_proto protoreflect.FileDescriptor

var file_apps_department_pb_request_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1d, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b,
	0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70,
	0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x01, 0x0a, 0x17, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62,
	0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x79, 0x22, 0xab, 0x02, 0x0a, 0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61,
	0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x73, 0x6b, 0x69, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x24, 0x0a, 0x0e,
	0x77, 0x69, 0x74, 0x68, 0x5f, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x77, 0x69, 0x74, 0x68, 0x53, 0x75, 0x62, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x77, 0x69, 0x74,
	0x68, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x69,
	0x74, 0x68, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x77,
	0x69, 0x74, 0x68, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22,
	0xc1, 0x01, 0x0a, 0x18, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x44, 0x65, 0x70, 0x61,
	0x72, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x24, 0x0a, 0x0e, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x77, 0x69, 0x74, 0x68, 0x53, 0x75,
	0x62, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x77, 0x69, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x77, 0x69, 0x74, 0x68, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x22, 0xbc, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x45, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x29, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x88, 0x01,
	0x0a, 0x15, 0x4a, 0x6f, 0x69, 0x6e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x98, 0x02, 0x0a, 0x1a, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x4c,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x6b, 0x69, 0x70, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x73, 0x6b, 0x69, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x22, 0x47, 0x0a, 0x1d, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x94, 0x01, 0x0a,
	0x1a, 0x44, 0x65, 0x61, 0x6c, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4c, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_department_pb_request_proto_rawDescOnce sync.Once
	file_apps_department_pb_request_proto_rawDescData = file_apps_department_pb_request_proto_rawDesc
)

func file_apps_department_pb_request_proto_rawDescGZIP() []byte {
	file_apps_department_pb_request_proto_rawDescOnce.Do(func() {
		file_apps_department_pb_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_department_pb_request_proto_rawDescData)
	})
	return file_apps_department_pb_request_proto_rawDescData
}

var file_apps_department_pb_request_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_apps_department_pb_request_proto_goTypes = []interface{}{
	(*CreateDepartmentRequest)(nil),       // 0: infraboard.keyauth.department.CreateDepartmentRequest
	(*QueryDepartmentRequest)(nil),        // 1: infraboard.keyauth.department.QueryDepartmentRequest
	(*DescribeDeparmentRequest)(nil),      // 2: infraboard.keyauth.department.DescribeDeparmentRequest
	(*UpdateDepartmentRequest)(nil),       // 3: infraboard.keyauth.department.UpdateDepartmentRequest
	(*DeleteDepartmentRequest)(nil),       // 4: infraboard.keyauth.department.DeleteDepartmentRequest
	(*JoinDepartmentRequest)(nil),         // 5: infraboard.keyauth.department.JoinDepartmentRequest
	(*QueryApplicationFormRequet)(nil),    // 6: infraboard.keyauth.department.QueryApplicationFormRequet
	(*DescribeApplicationFormRequet)(nil), // 7: infraboard.keyauth.department.DescribeApplicationFormRequet
	(*DealApplicationFormRequest)(nil),    // 8: infraboard.keyauth.department.DealApplicationFormRequest
	(*request.PageRequest)(nil),           // 9: infraboard.mcube.page.PageRequest
	(types.UpdateMode)(0),                 // 10: infraboard.keyauth.types.UpdateMode
	(ApplicationFormStatus)(0),            // 11: infraboard.keyauth.department.ApplicationFormStatus
}
var file_apps_department_pb_request_proto_depIdxs = []int32{
	9,  // 0: infraboard.keyauth.department.QueryDepartmentRequest.page:type_name -> infraboard.mcube.page.PageRequest
	10, // 1: infraboard.keyauth.department.UpdateDepartmentRequest.update_mode:type_name -> infraboard.keyauth.types.UpdateMode
	0,  // 2: infraboard.keyauth.department.UpdateDepartmentRequest.data:type_name -> infraboard.keyauth.department.CreateDepartmentRequest
	9,  // 3: infraboard.keyauth.department.QueryApplicationFormRequet.page:type_name -> infraboard.mcube.page.PageRequest
	11, // 4: infraboard.keyauth.department.QueryApplicationFormRequet.status:type_name -> infraboard.keyauth.department.ApplicationFormStatus
	11, // 5: infraboard.keyauth.department.DealApplicationFormRequest.status:type_name -> infraboard.keyauth.department.ApplicationFormStatus
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_apps_department_pb_request_proto_init() }
func file_apps_department_pb_request_proto_init() {
	if File_apps_department_pb_request_proto != nil {
		return
	}
	file_apps_department_pb_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_department_pb_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDepartmentRequest); i {
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
		file_apps_department_pb_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDepartmentRequest); i {
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
		file_apps_department_pb_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeDeparmentRequest); i {
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
		file_apps_department_pb_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDepartmentRequest); i {
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
		file_apps_department_pb_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDepartmentRequest); i {
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
		file_apps_department_pb_request_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinDepartmentRequest); i {
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
		file_apps_department_pb_request_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryApplicationFormRequet); i {
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
		file_apps_department_pb_request_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeApplicationFormRequet); i {
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
		file_apps_department_pb_request_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DealApplicationFormRequest); i {
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
			RawDescriptor: file_apps_department_pb_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_department_pb_request_proto_goTypes,
		DependencyIndexes: file_apps_department_pb_request_proto_depIdxs,
		MessageInfos:      file_apps_department_pb_request_proto_msgTypes,
	}.Build()
	File_apps_department_pb_request_proto = out.File
	file_apps_department_pb_request_proto_rawDesc = nil
	file_apps_department_pb_request_proto_goTypes = nil
	file_apps_department_pb_request_proto_depIdxs = nil
}
