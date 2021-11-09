// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package role

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

// RoleServiceClient is the client API for RoleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleServiceClient interface {
	CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*Role, error)
	QueryRole(ctx context.Context, in *QueryRoleRequest, opts ...grpc.CallOption) (*Set, error)
	DescribeRole(ctx context.Context, in *DescribeRoleRequest, opts ...grpc.CallOption) (*Role, error)
	DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*Role, error)
	QueryPermission(ctx context.Context, in *QueryPermissionRequest, opts ...grpc.CallOption) (*PermissionSet, error)
	DescribePermission(ctx context.Context, in *DescribePermissionRequest, opts ...grpc.CallOption) (*Permission, error)
	AddPermissionToRole(ctx context.Context, in *AddPermissionToRoleRequest, opts ...grpc.CallOption) (*PermissionSet, error)
	RemovePermissionFromRole(ctx context.Context, in *RemovePermissionFromRoleRequest, opts ...grpc.CallOption) (*PermissionSet, error)
	UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...grpc.CallOption) (*Permission, error)
}

type roleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleServiceClient(cc grpc.ClientConnInterface) RoleServiceClient {
	return &roleServiceClient{cc}
}

func (c *roleServiceClient) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.role.RoleService/CreateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) QueryRole(ctx context.Context, in *QueryRoleRequest, opts ...grpc.CallOption) (*Set, error) {
	out := new(Set)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.role.RoleService/QueryRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) DescribeRole(ctx context.Context, in *DescribeRoleRequest, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.role.RoleService/DescribeRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.role.RoleService/DeleteRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) QueryPermission(ctx context.Context, in *QueryPermissionRequest, opts ...grpc.CallOption) (*PermissionSet, error) {
	out := new(PermissionSet)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.role.RoleService/QueryPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) DescribePermission(ctx context.Context, in *DescribePermissionRequest, opts ...grpc.CallOption) (*Permission, error) {
	out := new(Permission)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.role.RoleService/DescribePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) AddPermissionToRole(ctx context.Context, in *AddPermissionToRoleRequest, opts ...grpc.CallOption) (*PermissionSet, error) {
	out := new(PermissionSet)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.role.RoleService/AddPermissionToRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) RemovePermissionFromRole(ctx context.Context, in *RemovePermissionFromRoleRequest, opts ...grpc.CallOption) (*PermissionSet, error) {
	out := new(PermissionSet)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.role.RoleService/RemovePermissionFromRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...grpc.CallOption) (*Permission, error) {
	out := new(Permission)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.role.RoleService/UpdatePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServiceServer is the server API for RoleService service.
// All implementations must embed UnimplementedRoleServiceServer
// for forward compatibility
type RoleServiceServer interface {
	CreateRole(context.Context, *CreateRoleRequest) (*Role, error)
	QueryRole(context.Context, *QueryRoleRequest) (*Set, error)
	DescribeRole(context.Context, *DescribeRoleRequest) (*Role, error)
	DeleteRole(context.Context, *DeleteRoleRequest) (*Role, error)
	QueryPermission(context.Context, *QueryPermissionRequest) (*PermissionSet, error)
	DescribePermission(context.Context, *DescribePermissionRequest) (*Permission, error)
	AddPermissionToRole(context.Context, *AddPermissionToRoleRequest) (*PermissionSet, error)
	RemovePermissionFromRole(context.Context, *RemovePermissionFromRoleRequest) (*PermissionSet, error)
	UpdatePermission(context.Context, *UpdatePermissionRequest) (*Permission, error)
	mustEmbedUnimplementedRoleServiceServer()
}

// UnimplementedRoleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRoleServiceServer struct {
}

func (UnimplementedRoleServiceServer) CreateRole(context.Context, *CreateRoleRequest) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedRoleServiceServer) QueryRole(context.Context, *QueryRoleRequest) (*Set, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryRole not implemented")
}
func (UnimplementedRoleServiceServer) DescribeRole(context.Context, *DescribeRoleRequest) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeRole not implemented")
}
func (UnimplementedRoleServiceServer) DeleteRole(context.Context, *DeleteRoleRequest) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedRoleServiceServer) QueryPermission(context.Context, *QueryPermissionRequest) (*PermissionSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPermission not implemented")
}
func (UnimplementedRoleServiceServer) DescribePermission(context.Context, *DescribePermissionRequest) (*Permission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePermission not implemented")
}
func (UnimplementedRoleServiceServer) AddPermissionToRole(context.Context, *AddPermissionToRoleRequest) (*PermissionSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPermissionToRole not implemented")
}
func (UnimplementedRoleServiceServer) RemovePermissionFromRole(context.Context, *RemovePermissionFromRoleRequest) (*PermissionSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePermissionFromRole not implemented")
}
func (UnimplementedRoleServiceServer) UpdatePermission(context.Context, *UpdatePermissionRequest) (*Permission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePermission not implemented")
}
func (UnimplementedRoleServiceServer) mustEmbedUnimplementedRoleServiceServer() {}

// UnsafeRoleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleServiceServer will
// result in compilation errors.
type UnsafeRoleServiceServer interface {
	mustEmbedUnimplementedRoleServiceServer()
}

func RegisterRoleServiceServer(s grpc.ServiceRegistrar, srv RoleServiceServer) {
	s.RegisterService(&RoleService_ServiceDesc, srv)
}

func _RoleService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.role.RoleService/CreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).CreateRole(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_QueryRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).QueryRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.role.RoleService/QueryRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).QueryRole(ctx, req.(*QueryRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_DescribeRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).DescribeRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.role.RoleService/DescribeRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).DescribeRole(ctx, req.(*DescribeRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.role.RoleService/DeleteRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).DeleteRole(ctx, req.(*DeleteRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_QueryPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).QueryPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.role.RoleService/QueryPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).QueryPermission(ctx, req.(*QueryPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_DescribePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).DescribePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.role.RoleService/DescribePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).DescribePermission(ctx, req.(*DescribePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_AddPermissionToRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPermissionToRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).AddPermissionToRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.role.RoleService/AddPermissionToRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).AddPermissionToRole(ctx, req.(*AddPermissionToRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_RemovePermissionFromRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePermissionFromRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).RemovePermissionFromRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.role.RoleService/RemovePermissionFromRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).RemovePermissionFromRole(ctx, req.(*RemovePermissionFromRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_UpdatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).UpdatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.role.RoleService/UpdatePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).UpdatePermission(ctx, req.(*UpdatePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoleService_ServiceDesc is the grpc.ServiceDesc for RoleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.keyauth.role.RoleService",
	HandlerType: (*RoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRole",
			Handler:    _RoleService_CreateRole_Handler,
		},
		{
			MethodName: "QueryRole",
			Handler:    _RoleService_QueryRole_Handler,
		},
		{
			MethodName: "DescribeRole",
			Handler:    _RoleService_DescribeRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _RoleService_DeleteRole_Handler,
		},
		{
			MethodName: "QueryPermission",
			Handler:    _RoleService_QueryPermission_Handler,
		},
		{
			MethodName: "DescribePermission",
			Handler:    _RoleService_DescribePermission_Handler,
		},
		{
			MethodName: "AddPermissionToRole",
			Handler:    _RoleService_AddPermissionToRole_Handler,
		},
		{
			MethodName: "RemovePermissionFromRole",
			Handler:    _RoleService_RemovePermissionFromRole_Handler,
		},
		{
			MethodName: "UpdatePermission",
			Handler:    _RoleService_UpdatePermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/role/pb/service.proto",
}
