// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package deploy

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

// DeployApiClient is the client API for DeployApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeployApiClient interface {
	DeployImage(ctx context.Context, in *DeployImageReq, opts ...grpc.CallOption) (*empty.Empty, error)
	Redeploy(ctx context.Context, in *RedeployReq, opts ...grpc.CallOption) (*empty.Empty, error)
}

type deployApiClient struct {
	cc grpc.ClientConnInterface
}

func NewDeployApiClient(cc grpc.ClientConnInterface) DeployApiClient {
	return &deployApiClient{cc}
}

func (c *deployApiClient) DeployImage(ctx context.Context, in *DeployImageReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/deploy.DeployApi/DeployImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deployApiClient) Redeploy(ctx context.Context, in *RedeployReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/deploy.DeployApi/Redeploy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeployApiServer is the server API for DeployApi service.
// All implementations must embed UnimplementedDeployApiServer
// for forward compatibility
type DeployApiServer interface {
	DeployImage(context.Context, *DeployImageReq) (*empty.Empty, error)
	Redeploy(context.Context, *RedeployReq) (*empty.Empty, error)
	mustEmbedUnimplementedDeployApiServer()
}

// UnimplementedDeployApiServer must be embedded to have forward compatible implementations.
type UnimplementedDeployApiServer struct {
}

func (UnimplementedDeployApiServer) DeployImage(context.Context, *DeployImageReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployImage not implemented")
}
func (UnimplementedDeployApiServer) Redeploy(context.Context, *RedeployReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Redeploy not implemented")
}
func (UnimplementedDeployApiServer) mustEmbedUnimplementedDeployApiServer() {}

// UnsafeDeployApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeployApiServer will
// result in compilation errors.
type UnsafeDeployApiServer interface {
	mustEmbedUnimplementedDeployApiServer()
}

func RegisterDeployApiServer(s *grpc.Server, srv DeployApiServer) {
	s.RegisterService(&_DeployApi_serviceDesc, srv)
}

func _DeployApi_DeployImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployImageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployApiServer).DeployImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deploy.DeployApi/DeployImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployApiServer).DeployImage(ctx, req.(*DeployImageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeployApi_Redeploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedeployReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployApiServer).Redeploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deploy.DeployApi/Redeploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployApiServer).Redeploy(ctx, req.(*RedeployReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeployApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "deploy.DeployApi",
	HandlerType: (*DeployApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeployImage",
			Handler:    _DeployApi_DeployImage_Handler,
		},
		{
			MethodName: "Redeploy",
			Handler:    _DeployApi_Redeploy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deploy.proto",
}
