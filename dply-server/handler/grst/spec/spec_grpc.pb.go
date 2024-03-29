// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package spec

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

// SpecApiClient is the client API for SpecApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpecApiClient interface {
	GetEnvar(ctx context.Context, in *GetEnvarReq, opts ...grpc.CallOption) (*Envar, error)
	UpsertEnvar(ctx context.Context, in *UpsertEnvarReq, opts ...grpc.CallOption) (*empty.Empty, error)
	GetScale(ctx context.Context, in *GetScaleReq, opts ...grpc.CallOption) (*Scale, error)
	UpsertScale(ctx context.Context, in *UpsertScaleReq, opts ...grpc.CallOption) (*empty.Empty, error)
	GetPort(ctx context.Context, in *GetPortReq, opts ...grpc.CallOption) (*Ports, error)
	UpsertPort(ctx context.Context, in *UpsertPortReq, opts ...grpc.CallOption) (*empty.Empty, error)
	GetAffinity(ctx context.Context, in *GetAffinityReq, opts ...grpc.CallOption) (*Affinity, error)
	UpsertAffinity(ctx context.Context, in *UpsertAffinityReq, opts ...grpc.CallOption) (*empty.Empty, error)
	GetDeploymentConfig(ctx context.Context, in *GetDeploymentConfigReq, opts ...grpc.CallOption) (*DeploymentConfig, error)
	UpsertDeploymentConfig(ctx context.Context, in *UpsertDeploymentConfigReq, opts ...grpc.CallOption) (*empty.Empty, error)
}

type specApiClient struct {
	cc grpc.ClientConnInterface
}

func NewSpecApiClient(cc grpc.ClientConnInterface) SpecApiClient {
	return &specApiClient{cc}
}

func (c *specApiClient) GetEnvar(ctx context.Context, in *GetEnvarReq, opts ...grpc.CallOption) (*Envar, error) {
	out := new(Envar)
	err := c.cc.Invoke(ctx, "/spec.SpecApi/GetEnvar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *specApiClient) UpsertEnvar(ctx context.Context, in *UpsertEnvarReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/spec.SpecApi/UpsertEnvar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *specApiClient) GetScale(ctx context.Context, in *GetScaleReq, opts ...grpc.CallOption) (*Scale, error) {
	out := new(Scale)
	err := c.cc.Invoke(ctx, "/spec.SpecApi/GetScale", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *specApiClient) UpsertScale(ctx context.Context, in *UpsertScaleReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/spec.SpecApi/UpsertScale", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *specApiClient) GetPort(ctx context.Context, in *GetPortReq, opts ...grpc.CallOption) (*Ports, error) {
	out := new(Ports)
	err := c.cc.Invoke(ctx, "/spec.SpecApi/GetPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *specApiClient) UpsertPort(ctx context.Context, in *UpsertPortReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/spec.SpecApi/UpsertPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *specApiClient) GetAffinity(ctx context.Context, in *GetAffinityReq, opts ...grpc.CallOption) (*Affinity, error) {
	out := new(Affinity)
	err := c.cc.Invoke(ctx, "/spec.SpecApi/GetAffinity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *specApiClient) UpsertAffinity(ctx context.Context, in *UpsertAffinityReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/spec.SpecApi/UpsertAffinity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *specApiClient) GetDeploymentConfig(ctx context.Context, in *GetDeploymentConfigReq, opts ...grpc.CallOption) (*DeploymentConfig, error) {
	out := new(DeploymentConfig)
	err := c.cc.Invoke(ctx, "/spec.SpecApi/GetDeploymentConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *specApiClient) UpsertDeploymentConfig(ctx context.Context, in *UpsertDeploymentConfigReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/spec.SpecApi/UpsertDeploymentConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpecApiServer is the server API for SpecApi service.
// All implementations must embed UnimplementedSpecApiServer
// for forward compatibility
type SpecApiServer interface {
	GetEnvar(context.Context, *GetEnvarReq) (*Envar, error)
	UpsertEnvar(context.Context, *UpsertEnvarReq) (*empty.Empty, error)
	GetScale(context.Context, *GetScaleReq) (*Scale, error)
	UpsertScale(context.Context, *UpsertScaleReq) (*empty.Empty, error)
	GetPort(context.Context, *GetPortReq) (*Ports, error)
	UpsertPort(context.Context, *UpsertPortReq) (*empty.Empty, error)
	GetAffinity(context.Context, *GetAffinityReq) (*Affinity, error)
	UpsertAffinity(context.Context, *UpsertAffinityReq) (*empty.Empty, error)
	GetDeploymentConfig(context.Context, *GetDeploymentConfigReq) (*DeploymentConfig, error)
	UpsertDeploymentConfig(context.Context, *UpsertDeploymentConfigReq) (*empty.Empty, error)
	mustEmbedUnimplementedSpecApiServer()
}

// UnimplementedSpecApiServer must be embedded to have forward compatible implementations.
type UnimplementedSpecApiServer struct {
}

func (UnimplementedSpecApiServer) GetEnvar(context.Context, *GetEnvarReq) (*Envar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEnvar not implemented")
}
func (UnimplementedSpecApiServer) UpsertEnvar(context.Context, *UpsertEnvarReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertEnvar not implemented")
}
func (UnimplementedSpecApiServer) GetScale(context.Context, *GetScaleReq) (*Scale, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScale not implemented")
}
func (UnimplementedSpecApiServer) UpsertScale(context.Context, *UpsertScaleReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertScale not implemented")
}
func (UnimplementedSpecApiServer) GetPort(context.Context, *GetPortReq) (*Ports, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPort not implemented")
}
func (UnimplementedSpecApiServer) UpsertPort(context.Context, *UpsertPortReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertPort not implemented")
}
func (UnimplementedSpecApiServer) GetAffinity(context.Context, *GetAffinityReq) (*Affinity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAffinity not implemented")
}
func (UnimplementedSpecApiServer) UpsertAffinity(context.Context, *UpsertAffinityReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertAffinity not implemented")
}
func (UnimplementedSpecApiServer) GetDeploymentConfig(context.Context, *GetDeploymentConfigReq) (*DeploymentConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeploymentConfig not implemented")
}
func (UnimplementedSpecApiServer) UpsertDeploymentConfig(context.Context, *UpsertDeploymentConfigReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertDeploymentConfig not implemented")
}
func (UnimplementedSpecApiServer) mustEmbedUnimplementedSpecApiServer() {}

// UnsafeSpecApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpecApiServer will
// result in compilation errors.
type UnsafeSpecApiServer interface {
	mustEmbedUnimplementedSpecApiServer()
}

func RegisterSpecApiServer(s *grpc.Server, srv SpecApiServer) {
	s.RegisterService(&_SpecApi_serviceDesc, srv)
}

func _SpecApi_GetEnvar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEnvarReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpecApiServer).GetEnvar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spec.SpecApi/GetEnvar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpecApiServer).GetEnvar(ctx, req.(*GetEnvarReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpecApi_UpsertEnvar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertEnvarReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpecApiServer).UpsertEnvar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spec.SpecApi/UpsertEnvar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpecApiServer).UpsertEnvar(ctx, req.(*UpsertEnvarReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpecApi_GetScale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScaleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpecApiServer).GetScale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spec.SpecApi/GetScale",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpecApiServer).GetScale(ctx, req.(*GetScaleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpecApi_UpsertScale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertScaleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpecApiServer).UpsertScale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spec.SpecApi/UpsertScale",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpecApiServer).UpsertScale(ctx, req.(*UpsertScaleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpecApi_GetPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPortReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpecApiServer).GetPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spec.SpecApi/GetPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpecApiServer).GetPort(ctx, req.(*GetPortReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpecApi_UpsertPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertPortReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpecApiServer).UpsertPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spec.SpecApi/UpsertPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpecApiServer).UpsertPort(ctx, req.(*UpsertPortReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpecApi_GetAffinity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAffinityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpecApiServer).GetAffinity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spec.SpecApi/GetAffinity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpecApiServer).GetAffinity(ctx, req.(*GetAffinityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpecApi_UpsertAffinity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertAffinityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpecApiServer).UpsertAffinity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spec.SpecApi/UpsertAffinity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpecApiServer).UpsertAffinity(ctx, req.(*UpsertAffinityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpecApi_GetDeploymentConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeploymentConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpecApiServer).GetDeploymentConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spec.SpecApi/GetDeploymentConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpecApiServer).GetDeploymentConfig(ctx, req.(*GetDeploymentConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpecApi_UpsertDeploymentConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertDeploymentConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpecApiServer).UpsertDeploymentConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spec.SpecApi/UpsertDeploymentConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpecApiServer).UpsertDeploymentConfig(ctx, req.(*UpsertDeploymentConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _SpecApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spec.SpecApi",
	HandlerType: (*SpecApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEnvar",
			Handler:    _SpecApi_GetEnvar_Handler,
		},
		{
			MethodName: "UpsertEnvar",
			Handler:    _SpecApi_UpsertEnvar_Handler,
		},
		{
			MethodName: "GetScale",
			Handler:    _SpecApi_GetScale_Handler,
		},
		{
			MethodName: "UpsertScale",
			Handler:    _SpecApi_UpsertScale_Handler,
		},
		{
			MethodName: "GetPort",
			Handler:    _SpecApi_GetPort_Handler,
		},
		{
			MethodName: "UpsertPort",
			Handler:    _SpecApi_UpsertPort_Handler,
		},
		{
			MethodName: "GetAffinity",
			Handler:    _SpecApi_GetAffinity_Handler,
		},
		{
			MethodName: "UpsertAffinity",
			Handler:    _SpecApi_UpsertAffinity_Handler,
		},
		{
			MethodName: "GetDeploymentConfig",
			Handler:    _SpecApi_GetDeploymentConfig_Handler,
		},
		{
			MethodName: "UpsertDeploymentConfig",
			Handler:    _SpecApi_UpsertDeploymentConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spec.proto",
}
