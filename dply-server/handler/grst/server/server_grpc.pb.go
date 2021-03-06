// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package server

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

// ServerApiClient is the client API for ServerApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerApiClient interface {
	Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatusResp, error)
}

type serverApiClient struct {
	cc grpc.ClientConnInterface
}

func NewServerApiClient(cc grpc.ClientConnInterface) ServerApiClient {
	return &serverApiClient{cc}
}

func (c *serverApiClient) Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatusResp, error) {
	out := new(StatusResp)
	err := c.cc.Invoke(ctx, "/server.ServerApi/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerApiServer is the server API for ServerApi service.
// All implementations must embed UnimplementedServerApiServer
// for forward compatibility
type ServerApiServer interface {
	Status(context.Context, *empty.Empty) (*StatusResp, error)
	mustEmbedUnimplementedServerApiServer()
}

// UnimplementedServerApiServer must be embedded to have forward compatible implementations.
type UnimplementedServerApiServer struct {
}

func (UnimplementedServerApiServer) Status(context.Context, *empty.Empty) (*StatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedServerApiServer) mustEmbedUnimplementedServerApiServer() {}

// UnsafeServerApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerApiServer will
// result in compilation errors.
type UnsafeServerApiServer interface {
	mustEmbedUnimplementedServerApiServer()
}

func RegisterServerApiServer(s *grpc.Server, srv ServerApiServer) {
	s.RegisterService(&_ServerApi_serviceDesc, srv)
}

func _ServerApi_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerApiServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.ServerApi/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerApiServer).Status(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServerApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.ServerApi",
	HandlerType: (*ServerApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _ServerApi_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
