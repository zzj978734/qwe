// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: pkg/domain/pb/setting.proto

package domain

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

// PasswordSecurity 密码安全设置
type PasswordSecurity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 密码长度
	// @gotags: bson:"length" json:"length" validate:"required,min=8,max=64"
	Length int32 `protobuf:"varint,1,opt,name=length,proto3" json:"length" bson:"length" validate:"required,min=8,max=64"`
	// 包含数字
	// @gotags: bson:"include_number" json:"include_number"
	IncludeNumber bool `protobuf:"varint,2,opt,name=include_number,json=includeNumber,proto3" json:"include_number" bson:"include_number"`
	// 包含小写字母
	// @gotags: bson:"include_lower_letter" json:"include_lower_letter"
	IncludeLowerLetter bool `protobuf:"varint,3,opt,name=include_lower_letter,json=includeLowerLetter,proto3" json:"include_lower_letter" bson:"include_lower_letter"`
	// 包含大写字母
	// @gotags: bson:"include_upper_letter" json:"include_upper_letter"
	IncludeUpperLetter bool `protobuf:"varint,4,opt,name=include_upper_letter,json=includeUpperLetter,proto3" json:"include_upper_letter" bson:"include_upper_letter"`
	// 包含特殊字符
	// @gotags: bson:"include_symbols" json:"include_symbols"
	IncludeSymbols bool `protobuf:"varint,5,opt,name=include_symbols,json=includeSymbols,proto3" json:"include_symbols" bson:"include_symbols"`
	// 重复限制
	// @gotags: bson:"repeate_limite" json:"repeate_limite" validate:"required,min=1,max=24"
	RepeateLimite uint32 `protobuf:"varint,6,opt,name=repeate_limite,json=repeateLimite,proto3" json:"repeate_limite" bson:"repeate_limite" validate:"required,min=1,max=24"`
	// 密码过期时间, 密码过期后要求用户重置密码
	// @gotags: bson:"password_expired_days" json:"password_expired_days" validate:"required,min=0,max=365"
	PasswordExpiredDays uint32 `protobuf:"varint,7,opt,name=password_expired_days,json=passwordExpiredDays,proto3" json:"password_expired_days" bson:"password_expired_days" validate:"required,min=0,max=365"`
	// 密码过期前多少天开始提醒
	// @gotags: bson:"before_expired_remind_days" json:"before_expired_remind_days" validate:"required,min=0,max=365"
	BeforeExpiredRemindDays uint32 `protobuf:"varint,8,opt,name=before_expired_remind_days,json=beforeExpiredRemindDays,proto3" json:"before_expired_remind_days" bson:"before_expired_remind_days" validate:"required,min=0,max=365"`
}

func (x *PasswordSecurity) Reset() {
	*x = PasswordSecurity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_domain_pb_setting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordSecurity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordSecurity) ProtoMessage() {}

func (x *PasswordSecurity) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_domain_pb_setting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordSecurity.ProtoReflect.Descriptor instead.
func (*PasswordSecurity) Descriptor() ([]byte, []int) {
	return file_pkg_domain_pb_setting_proto_rawDescGZIP(), []int{0}
}

func (x *PasswordSecurity) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *PasswordSecurity) GetIncludeNumber() bool {
	if x != nil {
		return x.IncludeNumber
	}
	return false
}

func (x *PasswordSecurity) GetIncludeLowerLetter() bool {
	if x != nil {
		return x.IncludeLowerLetter
	}
	return false
}

func (x *PasswordSecurity) GetIncludeUpperLetter() bool {
	if x != nil {
		return x.IncludeUpperLetter
	}
	return false
}

func (x *PasswordSecurity) GetIncludeSymbols() bool {
	if x != nil {
		return x.IncludeSymbols
	}
	return false
}

func (x *PasswordSecurity) GetRepeateLimite() uint32 {
	if x != nil {
		return x.RepeateLimite
	}
	return 0
}

func (x *PasswordSecurity) GetPasswordExpiredDays() uint32 {
	if x != nil {
		return x.PasswordExpiredDays
	}
	return 0
}

func (x *PasswordSecurity) GetBeforeExpiredRemindDays() uint32 {
	if x != nil {
		return x.BeforeExpiredRemindDays
	}
	return 0
}

// ExceptionLockConfig todo
type ExceptionLockConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 异地登录
	// @gotags: bson:"other_place_login" json:"other_place_login"
	OtherPlaceLogin bool `protobuf:"varint,1,opt,name=other_place_login,json=otherPlaceLogin,proto3" json:"other_place_login" bson:"other_place_login"`
	// 未登录天数
	// @gotags: bson:"not_login_days" json:"not_login_days"
	NotLoginDays uint32 `protobuf:"varint,2,opt,name=not_login_days,json=notLoginDays,proto3" json:"not_login_days" bson:"not_login_days"`
}

func (x *ExceptionLockConfig) Reset() {
	*x = ExceptionLockConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_domain_pb_setting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExceptionLockConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExceptionLockConfig) ProtoMessage() {}

func (x *ExceptionLockConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_domain_pb_setting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExceptionLockConfig.ProtoReflect.Descriptor instead.
func (*ExceptionLockConfig) Descriptor() ([]byte, []int) {
	return file_pkg_domain_pb_setting_proto_rawDescGZIP(), []int{1}
}

func (x *ExceptionLockConfig) GetOtherPlaceLogin() bool {
	if x != nil {
		return x.OtherPlaceLogin
	}
	return false
}

func (x *ExceptionLockConfig) GetNotLoginDays() uint32 {
	if x != nil {
		return x.NotLoginDays
	}
	return 0
}

// IPLimiteConfig todo
type IPLimiteConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 黑名单还是白名单
	// @gotags: bson:"type" json:"type"
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type" bson:"type"`
	// ip列表
	// @gotags: bson:"ip" json:"ip"
	Ip []string `protobuf:"bytes,2,rep,name=ip,proto3" json:"ip" bson:"ip"`
}

func (x *IPLimiteConfig) Reset() {
	*x = IPLimiteConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_domain_pb_setting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPLimiteConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPLimiteConfig) ProtoMessage() {}

func (x *IPLimiteConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_domain_pb_setting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPLimiteConfig.ProtoReflect.Descriptor instead.
func (*IPLimiteConfig) Descriptor() ([]byte, []int) {
	return file_pkg_domain_pb_setting_proto_rawDescGZIP(), []int{2}
}

func (x *IPLimiteConfig) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *IPLimiteConfig) GetIp() []string {
	if x != nil {
		return x.Ip
	}
	return nil
}

// RetryLockConfig 重试锁配置
type RetryLockConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 重试限制
	// @gotags: bson:"retry_limite" json:"retry_limite"
	RetryLimite uint32 `protobuf:"varint,1,opt,name=retry_limite,json=retryLimite,proto3" json:"retry_limite" bson:"retry_limite"`
	// 锁定时长
	// @gotags: bson:"locked_minite" json:"locked_minite"
	LockedMinite uint32 `protobuf:"varint,2,opt,name=locked_minite,json=lockedMinite,proto3" json:"locked_minite" bson:"locked_minite"`
}

func (x *RetryLockConfig) Reset() {
	*x = RetryLockConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_domain_pb_setting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryLockConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryLockConfig) ProtoMessage() {}

func (x *RetryLockConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_domain_pb_setting_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryLockConfig.ProtoReflect.Descriptor instead.
func (*RetryLockConfig) Descriptor() ([]byte, []int) {
	return file_pkg_domain_pb_setting_proto_rawDescGZIP(), []int{3}
}

func (x *RetryLockConfig) GetRetryLimite() uint32 {
	if x != nil {
		return x.RetryLimite
	}
	return 0
}

func (x *RetryLockConfig) GetLockedMinite() uint32 {
	if x != nil {
		return x.LockedMinite
	}
	return 0
}

// LoginSecurity 登录安全
type LoginSecurity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 异常登录锁
	// @gotags: bson:"exception_lock" json:"exception_lock"
	ExceptionLock bool `protobuf:"varint,1,opt,name=exception_lock,json=exceptionLock,proto3" json:"exception_lock" bson:"exception_lock"`
	// 异常配置
	// @gotags: bson:"exception_lock_config" json:"exception_lock_config"
	ExceptionLockConfig *ExceptionLockConfig `protobuf:"bytes,2,opt,name=exception_lock_config,json=exceptionLockConfig,proto3" json:"exception_lock_config" bson:"exception_lock_config"`
	// 重试锁
	// @gotags: bson:"retry_lock" json:"retry_lock"
	RetryLock bool `protobuf:"varint,3,opt,name=retry_lock,json=retryLock,proto3" json:"retry_lock" bson:"retry_lock"`
	// 重试锁配置
	// @gotags: bson:"retry_lock_config" json:"retry_lock_config"
	RetryLockConfig *RetryLockConfig `protobuf:"bytes,4,opt,name=retry_lock_config,json=retryLockConfig,proto3" json:"retry_lock_config" bson:"retry_lock_config"`
	// IP限制
	// @gotags: bson:"ip_limite" json:"ip_limite"
	IpLimite bool `protobuf:"varint,5,opt,name=ip_limite,json=ipLimite,proto3" json:"ip_limite" bson:"ip_limite"`
	// IP限制配置
	// @gotags: bson:"ip_limite_config" json:"ip_limite_config"
	IpLimiteConfig *IPLimiteConfig `protobuf:"bytes,6,opt,name=ip_limite_config,json=ipLimiteConfig,proto3" json:"ip_limite_config" bson:"ip_limite_config"`
}

func (x *LoginSecurity) Reset() {
	*x = LoginSecurity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_domain_pb_setting_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginSecurity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginSecurity) ProtoMessage() {}

func (x *LoginSecurity) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_domain_pb_setting_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginSecurity.ProtoReflect.Descriptor instead.
func (*LoginSecurity) Descriptor() ([]byte, []int) {
	return file_pkg_domain_pb_setting_proto_rawDescGZIP(), []int{4}
}

func (x *LoginSecurity) GetExceptionLock() bool {
	if x != nil {
		return x.ExceptionLock
	}
	return false
}

func (x *LoginSecurity) GetExceptionLockConfig() *ExceptionLockConfig {
	if x != nil {
		return x.ExceptionLockConfig
	}
	return nil
}

func (x *LoginSecurity) GetRetryLock() bool {
	if x != nil {
		return x.RetryLock
	}
	return false
}

func (x *LoginSecurity) GetRetryLockConfig() *RetryLockConfig {
	if x != nil {
		return x.RetryLockConfig
	}
	return nil
}

func (x *LoginSecurity) GetIpLimite() bool {
	if x != nil {
		return x.IpLimite
	}
	return false
}

func (x *LoginSecurity) GetIpLimiteConfig() *IPLimiteConfig {
	if x != nil {
		return x.IpLimiteConfig
	}
	return nil
}

var File_pkg_domain_pb_setting_proto protoreflect.FileDescriptor

var file_pkg_domain_pb_setting_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x62, 0x2f,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0xf6, 0x02, 0x0a, 0x10, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69,
	0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x14,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x6c, 0x65,
	0x74, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x12, 0x30,
	0x0a, 0x14, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x75, 0x70, 0x70, 0x65, 0x72, 0x5f,
	0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x69, 0x6e,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x55, 0x70, 0x70, 0x65, 0x72, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72,
	0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x6e, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65,
	0x12, 0x32, 0x0a, 0x15, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x13, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64,
	0x44, 0x61, 0x79, 0x73, 0x12, 0x3b, 0x0a, 0x1a, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x5f, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x5f, 0x64, 0x61,
	0x79, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x17, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65,
	0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x44, 0x61, 0x79,
	0x73, 0x22, 0x67, 0x0a, 0x13, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f,
	0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2a, 0x0a, 0x11, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x6e, 0x6f, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6e, 0x6f,
	0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x79, 0x73, 0x22, 0x34, 0x0a, 0x0e, 0x49, 0x50,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70,
	0x22, 0x59, 0x0a, 0x0f, 0x52, 0x65, 0x74, 0x72, 0x79, 0x4c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x65, 0x74, 0x72, 0x79,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x5f, 0x6d, 0x69, 0x6e, 0x69, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x4d, 0x69, 0x6e, 0x69, 0x74, 0x65, 0x22, 0x83, 0x03, 0x0a, 0x0d,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x25, 0x0a,
	0x0e, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x62, 0x0a, 0x15, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x13, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f,
	0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x72,
	0x79, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x72, 0x65,
	0x74, 0x72, 0x79, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x56, 0x0a, 0x11, 0x72, 0x65, 0x74, 0x72, 0x79,
	0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52,
	0x65, 0x74, 0x72, 0x79, 0x4c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0f,
	0x72, 0x65, 0x74, 0x72, 0x79, 0x4c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x70, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x70, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x12, 0x53, 0x0a, 0x10,
	0x69, 0x70, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x49, 0x50, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x0e, 0x69, 0x70, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_domain_pb_setting_proto_rawDescOnce sync.Once
	file_pkg_domain_pb_setting_proto_rawDescData = file_pkg_domain_pb_setting_proto_rawDesc
)

func file_pkg_domain_pb_setting_proto_rawDescGZIP() []byte {
	file_pkg_domain_pb_setting_proto_rawDescOnce.Do(func() {
		file_pkg_domain_pb_setting_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_domain_pb_setting_proto_rawDescData)
	})
	return file_pkg_domain_pb_setting_proto_rawDescData
}

var file_pkg_domain_pb_setting_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_domain_pb_setting_proto_goTypes = []interface{}{
	(*PasswordSecurity)(nil),    // 0: infraboard.keyauth.domain.PasswordSecurity
	(*ExceptionLockConfig)(nil), // 1: infraboard.keyauth.domain.ExceptionLockConfig
	(*IPLimiteConfig)(nil),      // 2: infraboard.keyauth.domain.IPLimiteConfig
	(*RetryLockConfig)(nil),     // 3: infraboard.keyauth.domain.RetryLockConfig
	(*LoginSecurity)(nil),       // 4: infraboard.keyauth.domain.LoginSecurity
}
var file_pkg_domain_pb_setting_proto_depIdxs = []int32{
	1, // 0: infraboard.keyauth.domain.LoginSecurity.exception_lock_config:type_name -> infraboard.keyauth.domain.ExceptionLockConfig
	3, // 1: infraboard.keyauth.domain.LoginSecurity.retry_lock_config:type_name -> infraboard.keyauth.domain.RetryLockConfig
	2, // 2: infraboard.keyauth.domain.LoginSecurity.ip_limite_config:type_name -> infraboard.keyauth.domain.IPLimiteConfig
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_domain_pb_setting_proto_init() }
func file_pkg_domain_pb_setting_proto_init() {
	if File_pkg_domain_pb_setting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_domain_pb_setting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordSecurity); i {
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
		file_pkg_domain_pb_setting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExceptionLockConfig); i {
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
		file_pkg_domain_pb_setting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPLimiteConfig); i {
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
		file_pkg_domain_pb_setting_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetryLockConfig); i {
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
		file_pkg_domain_pb_setting_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginSecurity); i {
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
			RawDescriptor: file_pkg_domain_pb_setting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_domain_pb_setting_proto_goTypes,
		DependencyIndexes: file_pkg_domain_pb_setting_proto_depIdxs,
		MessageInfos:      file_pkg_domain_pb_setting_proto_msgTypes,
	}.Build()
	File_pkg_domain_pb_setting_proto = out.File
	file_pkg_domain_pb_setting_proto_rawDesc = nil
	file_pkg_domain_pb_setting_proto_goTypes = nil
	file_pkg_domain_pb_setting_proto_depIdxs = nil
}
