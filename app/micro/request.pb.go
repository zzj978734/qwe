// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: app/micro/pb/request.proto

package micro

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

// ValidateClientCredentialRequest 校验服务凭证
type ValidateClientCredentialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 服务客户端ID
	// @gotags: json:"client_id" validate:"required,lte=100"
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id" validate:"required,lte=100"`
	// 服务客户端凭证
	// @gotags: json:"client_secret" validate:"required,lte=100"
	ClientSecret string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret" validate:"required,lte=100"`
}

func (x *ValidateClientCredentialRequest) Reset() {
	*x = ValidateClientCredentialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_micro_pb_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateClientCredentialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateClientCredentialRequest) ProtoMessage() {}

func (x *ValidateClientCredentialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_micro_pb_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateClientCredentialRequest.ProtoReflect.Descriptor instead.
func (*ValidateClientCredentialRequest) Descriptor() ([]byte, []int) {
	return file_app_micro_pb_request_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateClientCredentialRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *ValidateClientCredentialRequest) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

// CreateMicroRequest 服务创建请求
type CreateMicroRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 服务类型
	// @gotags: bson:"type" json:"type"
	Type Type `protobuf:"varint,1,opt,name=type,proto3,enum=infraboard.keyauth.micro.Type" json:"type" bson:"type"`
	// 名称
	// @gotags: bson:"name" json:"name" validate:"required,lte=200"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" bson:"name" validate:"required,lte=200"`
	// 服务标签
	// @gotags: bson:"label" json:"label" validate:"lte=80"
	Label map[string]string `protobuf:"bytes,3,rep,name=label,proto3" json:"label" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"label" validate:"lte=80"`
	// 描述信息
	// @gotags: bson:"description" json:"description,omitempty"
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty" bson:"description"`
	// 服务所属域
	// @gotags: bson:"domain" json:"domain"
	Domain string `protobuf:"bytes,5,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// 创建者ID
	// @gotags: bson:"creater" json:"creater"
	Creater string `protobuf:"bytes,6,opt,name=creater,proto3" json:"creater" bson:"creater"`
}

func (x *CreateMicroRequest) Reset() {
	*x = CreateMicroRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_micro_pb_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMicroRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMicroRequest) ProtoMessage() {}

func (x *CreateMicroRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_micro_pb_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMicroRequest.ProtoReflect.Descriptor instead.
func (*CreateMicroRequest) Descriptor() ([]byte, []int) {
	return file_app_micro_pb_request_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMicroRequest) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_NULL
}

func (x *CreateMicroRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateMicroRequest) GetLabel() map[string]string {
	if x != nil {
		return x.Label
	}
	return nil
}

func (x *CreateMicroRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateMicroRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CreateMicroRequest) GetCreater() string {
	if x != nil {
		return x.Creater
	}
	return ""
}

// QueryMicroRequest 查询应用列表
type QueryMicroRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"page"
	Page *page.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 服务类型
	// @gotags: json:"type"
	Type Type `protobuf:"varint,2,opt,name=type,proto3,enum=infraboard.keyauth.micro.Type" json:"type"`
}

func (x *QueryMicroRequest) Reset() {
	*x = QueryMicroRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_micro_pb_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMicroRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMicroRequest) ProtoMessage() {}

func (x *QueryMicroRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_micro_pb_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMicroRequest.ProtoReflect.Descriptor instead.
func (*QueryMicroRequest) Descriptor() ([]byte, []int) {
	return file_app_micro_pb_request_proto_rawDescGZIP(), []int{2}
}

func (x *QueryMicroRequest) GetPage() *page.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryMicroRequest) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_NULL
}

// DescribeMicroRequest 查询应用详情
type DescribeMicroRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"client_id"
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id"`
	// @gotags: json:"name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	// @gotags: json:"id"
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id"`
}

func (x *DescribeMicroRequest) Reset() {
	*x = DescribeMicroRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_micro_pb_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeMicroRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeMicroRequest) ProtoMessage() {}

func (x *DescribeMicroRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_micro_pb_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeMicroRequest.ProtoReflect.Descriptor instead.
func (*DescribeMicroRequest) Descriptor() ([]byte, []int) {
	return file_app_micro_pb_request_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeMicroRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *DescribeMicroRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DescribeMicroRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// DeleteMicroRequest todo
type DeleteMicroRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DeleteMicroRequest) Reset() {
	*x = DeleteMicroRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_micro_pb_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMicroRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMicroRequest) ProtoMessage() {}

func (x *DeleteMicroRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_micro_pb_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMicroRequest.ProtoReflect.Descriptor instead.
func (*DeleteMicroRequest) Descriptor() ([]byte, []int) {
	return file_app_micro_pb_request_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteMicroRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_app_micro_pb_request_proto protoreflect.FileDescriptor

var file_app_micro_pb_request_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x70, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x1a, 0x17, 0x61, 0x70, 0x70, 0x2f, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f,
	0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x63, 0x0a, 0x1f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x22, 0xb9, 0x02, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x69, 0x63, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x4d, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x37, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x1a, 0x38, 0x0a, 0x0a, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x7f, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x57, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4d, 0x69, 0x63,
	0x72, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_app_micro_pb_request_proto_rawDescOnce sync.Once
	file_app_micro_pb_request_proto_rawDescData = file_app_micro_pb_request_proto_rawDesc
)

func file_app_micro_pb_request_proto_rawDescGZIP() []byte {
	file_app_micro_pb_request_proto_rawDescOnce.Do(func() {
		file_app_micro_pb_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_micro_pb_request_proto_rawDescData)
	})
	return file_app_micro_pb_request_proto_rawDescData
}

var file_app_micro_pb_request_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_app_micro_pb_request_proto_goTypes = []interface{}{
	(*ValidateClientCredentialRequest)(nil), // 0: infraboard.keyauth.micro.ValidateClientCredentialRequest
	(*CreateMicroRequest)(nil),              // 1: infraboard.keyauth.micro.CreateMicroRequest
	(*QueryMicroRequest)(nil),               // 2: infraboard.keyauth.micro.QueryMicroRequest
	(*DescribeMicroRequest)(nil),            // 3: infraboard.keyauth.micro.DescribeMicroRequest
	(*DeleteMicroRequest)(nil),              // 4: infraboard.keyauth.micro.DeleteMicroRequest
	nil,                                     // 5: infraboard.keyauth.micro.CreateMicroRequest.LabelEntry
	(Type)(0),                               // 6: infraboard.keyauth.micro.Type
	(*page.PageRequest)(nil),                // 7: infraboard.mcube.page.PageRequest
}
var file_app_micro_pb_request_proto_depIdxs = []int32{
	6, // 0: infraboard.keyauth.micro.CreateMicroRequest.type:type_name -> infraboard.keyauth.micro.Type
	5, // 1: infraboard.keyauth.micro.CreateMicroRequest.label:type_name -> infraboard.keyauth.micro.CreateMicroRequest.LabelEntry
	7, // 2: infraboard.keyauth.micro.QueryMicroRequest.page:type_name -> infraboard.mcube.page.PageRequest
	6, // 3: infraboard.keyauth.micro.QueryMicroRequest.type:type_name -> infraboard.keyauth.micro.Type
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_app_micro_pb_request_proto_init() }
func file_app_micro_pb_request_proto_init() {
	if File_app_micro_pb_request_proto != nil {
		return
	}
	file_app_micro_pb_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_app_micro_pb_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateClientCredentialRequest); i {
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
		file_app_micro_pb_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMicroRequest); i {
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
		file_app_micro_pb_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMicroRequest); i {
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
		file_app_micro_pb_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeMicroRequest); i {
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
		file_app_micro_pb_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMicroRequest); i {
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
			RawDescriptor: file_app_micro_pb_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_micro_pb_request_proto_goTypes,
		DependencyIndexes: file_app_micro_pb_request_proto_depIdxs,
		MessageInfos:      file_app_micro_pb_request_proto_msgTypes,
	}.Build()
	File_app_micro_pb_request_proto = out.File
	file_app_micro_pb_request_proto_rawDesc = nil
	file_app_micro_pb_request_proto_goTypes = nil
	file_app_micro_pb_request_proto_depIdxs = nil
}