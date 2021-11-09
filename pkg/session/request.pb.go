// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: pkg/session/pb/request.proto

package session

import (
	token "github.com/infraboard/keyauth/pkg/token"
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

// LogoutRequest 登出请求
type LogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Account   string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_session_pb_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_session_pb_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_pkg_session_pb_request_proto_rawDescGZIP(), []int{0}
}

func (x *LogoutRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *LogoutRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

// DescribeSessionRequest todo
type DescribeSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Domain    string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Account   string `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
	Login     bool   `protobuf:"varint,4,opt,name=login,proto3" json:"login,omitempty"`
}

func (x *DescribeSessionRequest) Reset() {
	*x = DescribeSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_session_pb_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeSessionRequest) ProtoMessage() {}

func (x *DescribeSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_session_pb_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeSessionRequest.ProtoReflect.Descriptor instead.
func (*DescribeSessionRequest) Descriptor() ([]byte, []int) {
	return file_pkg_session_pb_request_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeSessionRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *DescribeSessionRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *DescribeSessionRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *DescribeSessionRequest) GetLogin() bool {
	if x != nil {
		return x.Login
	}
	return false
}

// QuerySessionRequest todo
type QuerySessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page           *page.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Account        string            `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	LoginIp        string            `protobuf:"bytes,3,opt,name=login_ip,json=loginIp,proto3" json:"login_ip,omitempty"`
	LoginCity      string            `protobuf:"bytes,4,opt,name=login_city,json=loginCity,proto3" json:"login_city,omitempty"`
	ApplicationId  string            `protobuf:"bytes,5,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	GrantType      token.GrantType   `protobuf:"varint,6,opt,name=grant_type,json=grantType,proto3,enum=infraboard.keyauth.token.GrantType" json:"grant_type,omitempty"`
	StartLoginTime int64             `protobuf:"varint,7,opt,name=start_login_time,json=startLoginTime,proto3" json:"start_login_time,omitempty"`
	EndLoginTime   int64             `protobuf:"varint,8,opt,name=end_login_time,json=endLoginTime,proto3" json:"end_login_time,omitempty"`
	Domain         string            `protobuf:"bytes,9,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *QuerySessionRequest) Reset() {
	*x = QuerySessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_session_pb_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySessionRequest) ProtoMessage() {}

func (x *QuerySessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_session_pb_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuerySessionRequest.ProtoReflect.Descriptor instead.
func (*QuerySessionRequest) Descriptor() ([]byte, []int) {
	return file_pkg_session_pb_request_proto_rawDescGZIP(), []int{2}
}

func (x *QuerySessionRequest) GetPage() *page.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QuerySessionRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *QuerySessionRequest) GetLoginIp() string {
	if x != nil {
		return x.LoginIp
	}
	return ""
}

func (x *QuerySessionRequest) GetLoginCity() string {
	if x != nil {
		return x.LoginCity
	}
	return ""
}

func (x *QuerySessionRequest) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *QuerySessionRequest) GetGrantType() token.GrantType {
	if x != nil {
		return x.GrantType
	}
	return token.GrantType(0)
}

func (x *QuerySessionRequest) GetStartLoginTime() int64 {
	if x != nil {
		return x.StartLoginTime
	}
	return 0
}

func (x *QuerySessionRequest) GetEndLoginTime() int64 {
	if x != nil {
		return x.EndLoginTime
	}
	return 0
}

func (x *QuerySessionRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

// QueryUserLastSessionRequest todo
type QueryUserLastSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: validate:"required"
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty" validate:"required"`
}

func (x *QueryUserLastSessionRequest) Reset() {
	*x = QueryUserLastSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_session_pb_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUserLastSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserLastSessionRequest) ProtoMessage() {}

func (x *QueryUserLastSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_session_pb_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserLastSessionRequest.ProtoReflect.Descriptor instead.
func (*QueryUserLastSessionRequest) Descriptor() ([]byte, []int) {
	return file_pkg_session_pb_request_proto_rawDescGZIP(), []int{3}
}

func (x *QueryUserLastSessionRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

var File_pkg_session_pb_request_proto protoreflect.FileDescriptor

var file_pkg_session_pb_request_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x62,
	0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x18, 0x70, 0x6b, 0x67, 0x2f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62,
	0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7f,
	0x0a, 0x16, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x22,
	0xf4, 0x02, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x5f, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x49, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x63, 0x69,
	0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x43,
	0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x0a, 0x67, 0x72,
	0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x09, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28,
	0x0a, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x65, 0x6e, 0x64, 0x5f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x65, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x37, 0x0a, 0x1b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_session_pb_request_proto_rawDescOnce sync.Once
	file_pkg_session_pb_request_proto_rawDescData = file_pkg_session_pb_request_proto_rawDesc
)

func file_pkg_session_pb_request_proto_rawDescGZIP() []byte {
	file_pkg_session_pb_request_proto_rawDescOnce.Do(func() {
		file_pkg_session_pb_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_session_pb_request_proto_rawDescData)
	})
	return file_pkg_session_pb_request_proto_rawDescData
}

var file_pkg_session_pb_request_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_session_pb_request_proto_goTypes = []interface{}{
	(*LogoutRequest)(nil),               // 0: infraboard.keyauth.session.LogoutRequest
	(*DescribeSessionRequest)(nil),      // 1: infraboard.keyauth.session.DescribeSessionRequest
	(*QuerySessionRequest)(nil),         // 2: infraboard.keyauth.session.QuerySessionRequest
	(*QueryUserLastSessionRequest)(nil), // 3: infraboard.keyauth.session.QueryUserLastSessionRequest
	(*page.PageRequest)(nil),            // 4: infraboard.mcube.page.PageRequest
	(token.GrantType)(0),                // 5: infraboard.keyauth.token.GrantType
}
var file_pkg_session_pb_request_proto_depIdxs = []int32{
	4, // 0: infraboard.keyauth.session.QuerySessionRequest.page:type_name -> infraboard.mcube.page.PageRequest
	5, // 1: infraboard.keyauth.session.QuerySessionRequest.grant_type:type_name -> infraboard.keyauth.token.GrantType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_session_pb_request_proto_init() }
func file_pkg_session_pb_request_proto_init() {
	if File_pkg_session_pb_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_session_pb_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutRequest); i {
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
		file_pkg_session_pb_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeSessionRequest); i {
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
		file_pkg_session_pb_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerySessionRequest); i {
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
		file_pkg_session_pb_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUserLastSessionRequest); i {
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
			RawDescriptor: file_pkg_session_pb_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_session_pb_request_proto_goTypes,
		DependencyIndexes: file_pkg_session_pb_request_proto_depIdxs,
		MessageInfos:      file_pkg_session_pb_request_proto_msgTypes,
	}.Build()
	File_pkg_session_pb_request_proto = out.File
	file_pkg_session_pb_request_proto_rawDesc = nil
	file_pkg_session_pb_request_proto_goTypes = nil
	file_pkg_session_pb_request_proto_depIdxs = nil
}
