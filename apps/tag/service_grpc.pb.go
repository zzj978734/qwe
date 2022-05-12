// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: apps/tag/pb/service.proto

package tag

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*TagKey, error)
	DescribeTag(ctx context.Context, in *DescribeTagRequest, opts ...grpc.CallOption) (*TagKey, error)
	DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*TagKey, error)
	QueryTagKey(ctx context.Context, in *QueryTagKeyRequest, opts ...grpc.CallOption) (*TagKeySet, error)
	QueryTagValue(ctx context.Context, in *QueryTagValueRequest, opts ...grpc.CallOption) (*TagValueSet, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*TagKey, error) {
	out := new(TagKey)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.tag.Service/CreateTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DescribeTag(ctx context.Context, in *DescribeTagRequest, opts ...grpc.CallOption) (*TagKey, error) {
	out := new(TagKey)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.tag.Service/DescribeTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*TagKey, error) {
	out := new(TagKey)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.tag.Service/DeleteTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryTagKey(ctx context.Context, in *QueryTagKeyRequest, opts ...grpc.CallOption) (*TagKeySet, error) {
	out := new(TagKeySet)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.tag.Service/QueryTagKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryTagValue(ctx context.Context, in *QueryTagValueRequest, opts ...grpc.CallOption) (*TagValueSet, error) {
	out := new(TagValueSet)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.tag.Service/QueryTagValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	CreateTag(context.Context, *CreateTagRequest) (*TagKey, error)
	DescribeTag(context.Context, *DescribeTagRequest) (*TagKey, error)
	DeleteTag(context.Context, *DeleteTagRequest) (*TagKey, error)
	QueryTagKey(context.Context, *QueryTagKeyRequest) (*TagKeySet, error)
	QueryTagValue(context.Context, *QueryTagValueRequest) (*TagValueSet, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) CreateTag(context.Context, *CreateTagRequest) (*TagKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTag not implemented")
}
func (UnimplementedServiceServer) DescribeTag(context.Context, *DescribeTagRequest) (*TagKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeTag not implemented")
}
func (UnimplementedServiceServer) DeleteTag(context.Context, *DeleteTagRequest) (*TagKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTag not implemented")
}
func (UnimplementedServiceServer) QueryTagKey(context.Context, *QueryTagKeyRequest) (*TagKeySet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTagKey not implemented")
}
func (UnimplementedServiceServer) QueryTagValue(context.Context, *QueryTagValueRequest) (*TagValueSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTagValue not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_CreateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.tag.Service/CreateTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateTag(ctx, req.(*CreateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DescribeTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DescribeTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.tag.Service/DescribeTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DescribeTag(ctx, req.(*DescribeTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.tag.Service/DeleteTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteTag(ctx, req.(*DeleteTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryTagKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTagKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryTagKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.tag.Service/QueryTagKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryTagKey(ctx, req.(*QueryTagKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryTagValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTagValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryTagValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.tag.Service/QueryTagValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryTagValue(ctx, req.(*QueryTagValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.keyauth.tag.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTag",
			Handler:    _Service_CreateTag_Handler,
		},
		{
			MethodName: "DescribeTag",
			Handler:    _Service_DescribeTag_Handler,
		},
		{
			MethodName: "DeleteTag",
			Handler:    _Service_DeleteTag_Handler,
		},
		{
			MethodName: "QueryTagKey",
			Handler:    _Service_QueryTagKey_Handler,
		},
		{
			MethodName: "QueryTagValue",
			Handler:    _Service_QueryTagValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/tag/pb/service.proto",
}