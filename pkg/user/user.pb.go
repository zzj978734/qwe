// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: pkg/user/pb/user.proto

package user

import (
	department "github.com/infraboard/keyauth/pkg/department"
	types "github.com/infraboard/keyauth/pkg/user/types"
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

type Password struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// hash过后的密码
	// @gotags: bson:"password" json:"password,omitempty"
	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty" bson:"password"`
	// 密码创建时间
	// @gotags: bson:"create_at" json:"create_at,omitempty"
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty" bson:"create_at"`
	// 密码更新时间
	// @gotags: bson:"update_at" json:"update_at,omitempty"
	UpdateAt int64 `protobuf:"varint,3,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty" bson:"update_at"`
	// 密码需要被重置
	// @gotags: bson:"need_reset" json:"need_reset"
	NeedReset bool `protobuf:"varint,4,opt,name=need_reset,json=needReset,proto3" json:"need_reset" bson:"need_reset"`
	// 需要重置的原因
	// @gotags: bson:"reset_reason" json:"reset_reason"
	ResetReason string `protobuf:"bytes,5,opt,name=reset_reason,json=resetReason,proto3" json:"reset_reason" bson:"reset_reason"`
	// 历史密码
	// @gotags: bson:"history" json:"history,omitempty"
	History []string `protobuf:"bytes,6,rep,name=history,proto3" json:"history,omitempty" bson:"history"`
	// 是否过期
	// @gotags: bson:"-" json:"is_expired"
	IsExpired bool `protobuf:"varint,7,opt,name=is_expired,json=isExpired,proto3" json:"is_expired" bson:"-"`
}

func (x *Password) Reset() {
	*x = Password{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_user_pb_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Password) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Password) ProtoMessage() {}

func (x *Password) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_user_pb_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Password.ProtoReflect.Descriptor instead.
func (*Password) Descriptor() ([]byte, []int) {
	return file_pkg_user_pb_user_proto_rawDescGZIP(), []int{0}
}

func (x *Password) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Password) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Password) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Password) GetNeedReset() bool {
	if x != nil {
		return x.NeedReset
	}
	return false
}

func (x *Password) GetResetReason() string {
	if x != nil {
		return x.ResetReason
	}
	return ""
}

func (x *Password) GetHistory() []string {
	if x != nil {
		return x.History
	}
	return nil
}

func (x *Password) GetIsExpired() bool {
	if x != nil {
		return x.IsExpired
	}
	return false
}

// Status 用户状态
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 是否冻结
	// @gotags: bson:"locked" json:"locked"
	Locked bool `protobuf:"varint,1,opt,name=locked,proto3" json:"locked" bson:"locked"`
	// 冻结时间
	// @gotags: bson:"locked_time" json:"locked_time,omitempty"
	LockedTime int64 `protobuf:"varint,2,opt,name=locked_time,json=lockedTime,proto3" json:"locked_time,omitempty" bson:"locked_time"`
	// 冻结原因
	// @gotags: bson:"locked_reson" json:"locked_reson,omitempty"
	LockedReson string `protobuf:"bytes,3,opt,name=locked_reson,json=lockedReson,proto3" json:"locked_reson,omitempty" bson:"locked_reson"`
	// 解冻时间
	// @gotags: bson:"unlock_time" json:"unlock_time,omitempty"
	UnlockTime int64 `protobuf:"varint,4,opt,name=unlock_time,json=unlockTime,proto3" json:"unlock_time,omitempty" bson:"unlock_time"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_user_pb_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_user_pb_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_pkg_user_pb_user_proto_rawDescGZIP(), []int{1}
}

func (x *Status) GetLocked() bool {
	if x != nil {
		return x.Locked
	}
	return false
}

func (x *Status) GetLockedTime() int64 {
	if x != nil {
		return x.LockedTime
	}
	return 0
}

func (x *Status) GetLockedReson() string {
	if x != nil {
		return x.LockedReson
	}
	return ""
}

func (x *Status) GetUnlockTime() int64 {
	if x != nil {
		return x.UnlockTime
	}
	return 0
}

// User info
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户所属部门
	// @gotags: bson:"department_id" json:"department_id" validate:"lte=200"
	DepartmentId string `protobuf:"bytes,1,opt,name=department_id,json=departmentId,proto3" json:"department_id" bson:"department_id" validate:"lte=200"`
	// 用户账号名称
	// @gotags: bson:"_id" json:"account" validate:"required,lte=60"
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account" bson:"_id" validate:"required,lte=60"`
	// 创建方式
	// @gotags: bson:"create_type" json:"create_type"
	CreateType CreateType `protobuf:"varint,3,opt,name=create_type,json=createType,proto3,enum=infraboard.keyauth.user.CreateType" json:"create_type" bson:"create_type"`
	// 用户创建的时间
	// @gotags: bson:"create_at" json:"create_at,omitempty"
	CreateAt int64 `protobuf:"varint,4,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty" bson:"create_at"`
	// 修改时间
	// @gotags: bson:"update_at" json:"update_at,omitempty"
	UpdateAt int64 `protobuf:"varint,5,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty" bson:"update_at"`
	// 如果是子账号和服务账号 都需要继承主用户Domain
	// @gotags: bson:"domain" json:"domain,omitempty"
	Domain string `protobuf:"bytes,6,opt,name=domain,proto3" json:"domain,omitempty" bson:"domain"`
	// 是否是主账号
	// @gotags: bson:"type"  json:"type"
	Type types.UserType `protobuf:"varint,7,opt,name=type,proto3,enum=infraboard.keyauth.user.UserType" json:"type" bson:"type"`
	// 数据
	// @gotags: bson:"profile" json:"profile"
	Profile *Profile `protobuf:"bytes,8,opt,name=profile,proto3" json:"profile" bson:"profile"`
	// 用户的角色(当携带Namesapce查询时会有)
	// @gotags: bson:"-" json:"roles,omitempty"
	Roles []string `protobuf:"bytes,9,rep,name=roles,proto3" json:"roles,omitempty" bson:"-"`
	// 用户多久未登录时(天), 冻结改用户, 防止僵尸用户的账号被利用
	// @gotags: bson:"expires_days" json:"expires_days"
	ExpiresDays int32 `protobuf:"varint,10,opt,name=expires_days,json=expiresDays,proto3" json:"expires_days" bson:"expires_days"`
	// 用户描述
	// @gotags: json:"description"
	Description string `protobuf:"bytes,11,opt,name=description,proto3" json:"description"`
	// 用户是否初始化
	// @gotags: bson:"is_initialized" json:"is_initialized"
	IsInitialized bool `protobuf:"varint,12,opt,name=is_initialized,json=isInitialized,proto3" json:"is_initialized" bson:"is_initialized"`
	// 密码相关信息
	// @gotags: bson:"password" json:"password"
	HashedPassword *Password `protobuf:"bytes,13,opt,name=hashed_password,json=hashedPassword,proto3" json:"password" bson:"password"`
	// 用户状态
	// @gotags: bson:"status" json:"status"
	Status *Status `protobuf:"bytes,14,opt,name=status,proto3" json:"status" bson:"status"`
	// 部门
	// @gotags: bson:"-" json:"department,omitempty"
	Department *department.Department `protobuf:"bytes,15,opt,name=department,proto3" json:"department,omitempty" bson:"-"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_user_pb_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_user_pb_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_pkg_user_pb_user_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetDepartmentId() string {
	if x != nil {
		return x.DepartmentId
	}
	return ""
}

func (x *User) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *User) GetCreateType() CreateType {
	if x != nil {
		return x.CreateType
	}
	return CreateType_USER_REGISTRY
}

func (x *User) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *User) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *User) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *User) GetType() types.UserType {
	if x != nil {
		return x.Type
	}
	return types.UserType(0)
}

func (x *User) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *User) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *User) GetExpiresDays() int32 {
	if x != nil {
		return x.ExpiresDays
	}
	return 0
}

func (x *User) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *User) GetIsInitialized() bool {
	if x != nil {
		return x.IsInitialized
	}
	return false
}

func (x *User) GetHashedPassword() *Password {
	if x != nil {
		return x.HashedPassword
	}
	return nil
}

func (x *User) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *User) GetDepartment() *department.Department {
	if x != nil {
		return x.Department
	}
	return nil
}

type Set struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// @gotags: bson:"items" json:"items"
	Items []*User `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *Set) Reset() {
	*x = Set{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_user_pb_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set) ProtoMessage() {}

func (x *Set) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_user_pb_user_proto_msgTypes[3]
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
	return file_pkg_user_pb_user_proto_rawDescGZIP(), []int{3}
}

func (x *Set) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Set) GetItems() []*User {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_pkg_user_pb_user_proto protoreflect.FileDescriptor

var file_pkg_user_pb_user_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x1a, 0x22, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x70, 0x62, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x70, 0x6b, 0x67, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xdb, 0x01, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6e, 0x65, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x22,
	0x85, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xa2, 0x05, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x44, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x35, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3a,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x64, 0x61, 0x79, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x44,
	0x61, 0x79, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69,
	0x73, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x4a, 0x0a, 0x0f,
	0x68, 0x61, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x0e, 0x68, 0x61, 0x73, 0x68, 0x65, 0x64,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x49, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x50, 0x0a, 0x03,
	0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x33, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x28,
	0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_user_pb_user_proto_rawDescOnce sync.Once
	file_pkg_user_pb_user_proto_rawDescData = file_pkg_user_pb_user_proto_rawDesc
)

func file_pkg_user_pb_user_proto_rawDescGZIP() []byte {
	file_pkg_user_pb_user_proto_rawDescOnce.Do(func() {
		file_pkg_user_pb_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_user_pb_user_proto_rawDescData)
	})
	return file_pkg_user_pb_user_proto_rawDescData
}

var file_pkg_user_pb_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_user_pb_user_proto_goTypes = []interface{}{
	(*Password)(nil),              // 0: infraboard.keyauth.user.Password
	(*Status)(nil),                // 1: infraboard.keyauth.user.Status
	(*User)(nil),                  // 2: infraboard.keyauth.user.User
	(*Set)(nil),                   // 3: infraboard.keyauth.user.Set
	(CreateType)(0),               // 4: infraboard.keyauth.user.CreateType
	(types.UserType)(0),           // 5: infraboard.keyauth.user.UserType
	(*Profile)(nil),               // 6: infraboard.keyauth.user.Profile
	(*department.Department)(nil), // 7: infraboard.keyauth.department.Department
}
var file_pkg_user_pb_user_proto_depIdxs = []int32{
	4, // 0: infraboard.keyauth.user.User.create_type:type_name -> infraboard.keyauth.user.CreateType
	5, // 1: infraboard.keyauth.user.User.type:type_name -> infraboard.keyauth.user.UserType
	6, // 2: infraboard.keyauth.user.User.profile:type_name -> infraboard.keyauth.user.Profile
	0, // 3: infraboard.keyauth.user.User.hashed_password:type_name -> infraboard.keyauth.user.Password
	1, // 4: infraboard.keyauth.user.User.status:type_name -> infraboard.keyauth.user.Status
	7, // 5: infraboard.keyauth.user.User.department:type_name -> infraboard.keyauth.department.Department
	2, // 6: infraboard.keyauth.user.Set.items:type_name -> infraboard.keyauth.user.User
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_pkg_user_pb_user_proto_init() }
func file_pkg_user_pb_user_proto_init() {
	if File_pkg_user_pb_user_proto != nil {
		return
	}
	file_pkg_user_pb_request_proto_init()
	file_pkg_user_pb_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_user_pb_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Password); i {
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
		file_pkg_user_pb_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_pkg_user_pb_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_pkg_user_pb_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_pkg_user_pb_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_user_pb_user_proto_goTypes,
		DependencyIndexes: file_pkg_user_pb_user_proto_depIdxs,
		MessageInfos:      file_pkg_user_pb_user_proto_msgTypes,
	}.Build()
	File_pkg_user_pb_user_proto = out.File
	file_pkg_user_pb_user_proto_rawDesc = nil
	file_pkg_user_pb_user_proto_goTypes = nil
	file_pkg_user_pb_user_proto_depIdxs = nil
}
