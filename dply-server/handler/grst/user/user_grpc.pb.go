// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package user

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// UserApiClient is the client API for UserApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserApiClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*User, error)
	GetCurrentLogin(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*User, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*empty.Empty, error)
}

type userApiClient struct {
	cc grpc.ClientConnInterface
}

func NewUserApiClient(cc grpc.ClientConnInterface) UserApiClient {
	return &userApiClient{cc}
}

func (c *userApiClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.UserApi/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiClient) GetCurrentLogin(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.UserApi/GetCurrentLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiClient) UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/user.UserApi/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserApiServer is the server API for UserApi service.
// All implementations must embed UnimplementedUserApiServer
// for forward compatibility
type UserApiServer interface {
	Login(context.Context, *LoginReq) (*User, error)
	GetCurrentLogin(context.Context, *empty.Empty) (*User, error)
	UpdatePassword(context.Context, *UpdatePasswordReq) (*empty.Empty, error)
	mustEmbedUnimplementedUserApiServer()
}

// UnimplementedUserApiServer must be embedded to have forward compatible implementations.
type UnimplementedUserApiServer struct {
}

func (UnimplementedUserApiServer) Login(context.Context, *LoginReq) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserApiServer) GetCurrentLogin(context.Context, *empty.Empty) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentLogin not implemented")
}
func (UnimplementedUserApiServer) UpdatePassword(context.Context, *UpdatePasswordReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedUserApiServer) mustEmbedUnimplementedUserApiServer() {}

// UnsafeUserApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserApiServer will
// result in compilation errors.
type UnsafeUserApiServer interface {
	mustEmbedUnimplementedUserApiServer()
}

func RegisterUserApiServer(s *grpc.Server, srv UserApiServer) {
	s.RegisterService(&_UserApi_serviceDesc, srv)
}

func _UserApi_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApiServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserApi/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApiServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserApi_GetCurrentLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApiServer).GetCurrentLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserApi/GetCurrentLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApiServer).GetCurrentLogin(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserApi_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApiServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserApi/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApiServer).UpdatePassword(ctx, req.(*UpdatePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserApi",
	HandlerType: (*UserApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserApi_Login_Handler,
		},
		{
			MethodName: "GetCurrentLogin",
			Handler:    _UserApi_GetCurrentLogin_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _UserApi_UpdatePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
