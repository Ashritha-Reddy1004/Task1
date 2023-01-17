// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto_pb/task.proto

package proto_pb

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	AddUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error)
	UpdateUser(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateRes, error)
	AddActivity(ctx context.Context, in *ActivityReq, opts ...grpc.CallOption) (*ActivityRes, error)
	IsValid(ctx context.Context, in *IsValidReq, opts ...grpc.CallOption) (*IsValidRes, error)
	IsDone(ctx context.Context, in *IsDoneReq, opts ...grpc.CallOption) (*IsDoneRes, error)
	RemoveUser(ctx context.Context, in *RemoveUserReq, opts ...grpc.CallOption) (*RemoveUserRes, error)
	PrintUser(ctx context.Context, in *PrintUserReq, opts ...grpc.CallOption) (*PrintUserRes, error)
	PrintActivity(ctx context.Context, in *PrintActivityReq, opts ...grpc.CallOption) (*PrintActivityRes, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) AddUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRes, error) {
	out := new(UserRes)
	err := c.cc.Invoke(ctx, "/proto_pb.UserService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateRes, error) {
	out := new(UpdateRes)
	err := c.cc.Invoke(ctx, "/proto_pb.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddActivity(ctx context.Context, in *ActivityReq, opts ...grpc.CallOption) (*ActivityRes, error) {
	out := new(ActivityRes)
	err := c.cc.Invoke(ctx, "/proto_pb.UserService/AddActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) IsValid(ctx context.Context, in *IsValidReq, opts ...grpc.CallOption) (*IsValidRes, error) {
	out := new(IsValidRes)
	err := c.cc.Invoke(ctx, "/proto_pb.UserService/IsValid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) IsDone(ctx context.Context, in *IsDoneReq, opts ...grpc.CallOption) (*IsDoneRes, error) {
	out := new(IsDoneRes)
	err := c.cc.Invoke(ctx, "/proto_pb.UserService/IsDone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RemoveUser(ctx context.Context, in *RemoveUserReq, opts ...grpc.CallOption) (*RemoveUserRes, error) {
	out := new(RemoveUserRes)
	err := c.cc.Invoke(ctx, "/proto_pb.UserService/RemoveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) PrintUser(ctx context.Context, in *PrintUserReq, opts ...grpc.CallOption) (*PrintUserRes, error) {
	out := new(PrintUserRes)
	err := c.cc.Invoke(ctx, "/proto_pb.UserService/PrintUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) PrintActivity(ctx context.Context, in *PrintActivityReq, opts ...grpc.CallOption) (*PrintActivityRes, error) {
	out := new(PrintActivityRes)
	err := c.cc.Invoke(ctx, "/proto_pb.UserService/PrintActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	AddUser(context.Context, *UserReq) (*UserRes, error)
	UpdateUser(context.Context, *UpdateReq) (*UpdateRes, error)
	AddActivity(context.Context, *ActivityReq) (*ActivityRes, error)
	IsValid(context.Context, *IsValidReq) (*IsValidRes, error)
	IsDone(context.Context, *IsDoneReq) (*IsDoneRes, error)
	RemoveUser(context.Context, *RemoveUserReq) (*RemoveUserRes, error)
	PrintUser(context.Context, *PrintUserReq) (*PrintUserRes, error)
	PrintActivity(context.Context, *PrintActivityReq) (*PrintActivityRes, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) AddUser(context.Context, *UserReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUser(context.Context, *UpdateReq) (*UpdateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) AddActivity(context.Context, *ActivityReq) (*ActivityRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddActivity not implemented")
}
func (UnimplementedUserServiceServer) IsValid(context.Context, *IsValidReq) (*IsValidRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsValid not implemented")
}
func (UnimplementedUserServiceServer) IsDone(context.Context, *IsDoneReq) (*IsDoneRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsDone not implemented")
}
func (UnimplementedUserServiceServer) RemoveUser(context.Context, *RemoveUserReq) (*RemoveUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedUserServiceServer) PrintUser(context.Context, *PrintUserReq) (*PrintUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrintUser not implemented")
}
func (UnimplementedUserServiceServer) PrintActivity(context.Context, *PrintActivityReq) (*PrintActivityRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrintActivity not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_pb.UserService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddUser(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_pb.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_pb.UserService/AddActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddActivity(ctx, req.(*ActivityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_IsValid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsValidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).IsValid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_pb.UserService/IsValid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).IsValid(ctx, req.(*IsValidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_IsDone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsDoneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).IsDone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_pb.UserService/IsDone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).IsDone(ctx, req.(*IsDoneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_pb.UserService/RemoveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RemoveUser(ctx, req.(*RemoveUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_PrintUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrintUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).PrintUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_pb.UserService/PrintUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).PrintUser(ctx, req.(*PrintUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_PrintActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrintActivityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).PrintActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_pb.UserService/PrintActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).PrintActivity(ctx, req.(*PrintActivityReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto_pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _UserService_AddUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "AddActivity",
			Handler:    _UserService_AddActivity_Handler,
		},
		{
			MethodName: "IsValid",
			Handler:    _UserService_IsValid_Handler,
		},
		{
			MethodName: "IsDone",
			Handler:    _UserService_IsDone_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _UserService_RemoveUser_Handler,
		},
		{
			MethodName: "PrintUser",
			Handler:    _UserService_PrintUser_Handler,
		},
		{
			MethodName: "PrintActivity",
			Handler:    _UserService_PrintActivity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto_pb/task.proto",
}
