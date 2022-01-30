// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package atmosphere

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

// AtmosphereServiceClient is the client API for AtmosphereService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AtmosphereServiceClient interface {
	// https://wiki.hoggitworld.com/view/DCS_func_getWind
	GetWind(ctx context.Context, in *GetWindRequest, opts ...grpc.CallOption) (*GetWindResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_getWindWithTurbulence
	GetWindWithTurbulence(ctx context.Context, in *GetWindWithTurbulenceRequest, opts ...grpc.CallOption) (*GetWindWithTurbulenceResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_getWindWithTurbulence
	GetTemperatureAndPressure(ctx context.Context, in *GetTemperatureAndPressureRequest, opts ...grpc.CallOption) (*GetTemperatureAndPressureResponse, error)
}

type atmosphereServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAtmosphereServiceClient(cc grpc.ClientConnInterface) AtmosphereServiceClient {
	return &atmosphereServiceClient{cc}
}

func (c *atmosphereServiceClient) GetWind(ctx context.Context, in *GetWindRequest, opts ...grpc.CallOption) (*GetWindResponse, error) {
	out := new(GetWindResponse)
	err := c.cc.Invoke(ctx, "/dcs.atmosphere.v0.AtmosphereService/GetWind", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *atmosphereServiceClient) GetWindWithTurbulence(ctx context.Context, in *GetWindWithTurbulenceRequest, opts ...grpc.CallOption) (*GetWindWithTurbulenceResponse, error) {
	out := new(GetWindWithTurbulenceResponse)
	err := c.cc.Invoke(ctx, "/dcs.atmosphere.v0.AtmosphereService/GetWindWithTurbulence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *atmosphereServiceClient) GetTemperatureAndPressure(ctx context.Context, in *GetTemperatureAndPressureRequest, opts ...grpc.CallOption) (*GetTemperatureAndPressureResponse, error) {
	out := new(GetTemperatureAndPressureResponse)
	err := c.cc.Invoke(ctx, "/dcs.atmosphere.v0.AtmosphereService/GetTemperatureAndPressure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AtmosphereServiceServer is the server API for AtmosphereService service.
// All implementations must embed UnimplementedAtmosphereServiceServer
// for forward compatibility
type AtmosphereServiceServer interface {
	// https://wiki.hoggitworld.com/view/DCS_func_getWind
	GetWind(context.Context, *GetWindRequest) (*GetWindResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_getWindWithTurbulence
	GetWindWithTurbulence(context.Context, *GetWindWithTurbulenceRequest) (*GetWindWithTurbulenceResponse, error)
	// https://wiki.hoggitworld.com/view/DCS_func_getWindWithTurbulence
	GetTemperatureAndPressure(context.Context, *GetTemperatureAndPressureRequest) (*GetTemperatureAndPressureResponse, error)
	mustEmbedUnimplementedAtmosphereServiceServer()
}

// UnimplementedAtmosphereServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAtmosphereServiceServer struct {
}

func (UnimplementedAtmosphereServiceServer) GetWind(context.Context, *GetWindRequest) (*GetWindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWind not implemented")
}
func (UnimplementedAtmosphereServiceServer) GetWindWithTurbulence(context.Context, *GetWindWithTurbulenceRequest) (*GetWindWithTurbulenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWindWithTurbulence not implemented")
}
func (UnimplementedAtmosphereServiceServer) GetTemperatureAndPressure(context.Context, *GetTemperatureAndPressureRequest) (*GetTemperatureAndPressureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemperatureAndPressure not implemented")
}
func (UnimplementedAtmosphereServiceServer) mustEmbedUnimplementedAtmosphereServiceServer() {}

// UnsafeAtmosphereServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AtmosphereServiceServer will
// result in compilation errors.
type UnsafeAtmosphereServiceServer interface {
	mustEmbedUnimplementedAtmosphereServiceServer()
}

func RegisterAtmosphereServiceServer(s grpc.ServiceRegistrar, srv AtmosphereServiceServer) {
	s.RegisterService(&AtmosphereService_ServiceDesc, srv)
}

func _AtmosphereService_GetWind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AtmosphereServiceServer).GetWind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.atmosphere.v0.AtmosphereService/GetWind",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AtmosphereServiceServer).GetWind(ctx, req.(*GetWindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AtmosphereService_GetWindWithTurbulence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWindWithTurbulenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AtmosphereServiceServer).GetWindWithTurbulence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.atmosphere.v0.AtmosphereService/GetWindWithTurbulence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AtmosphereServiceServer).GetWindWithTurbulence(ctx, req.(*GetWindWithTurbulenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AtmosphereService_GetTemperatureAndPressure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemperatureAndPressureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AtmosphereServiceServer).GetTemperatureAndPressure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcs.atmosphere.v0.AtmosphereService/GetTemperatureAndPressure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AtmosphereServiceServer).GetTemperatureAndPressure(ctx, req.(*GetTemperatureAndPressureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AtmosphereService_ServiceDesc is the grpc.ServiceDesc for AtmosphereService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AtmosphereService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dcs.atmosphere.v0.AtmosphereService",
	HandlerType: (*AtmosphereServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWind",
			Handler:    _AtmosphereService_GetWind_Handler,
		},
		{
			MethodName: "GetWindWithTurbulence",
			Handler:    _AtmosphereService_GetWindWithTurbulence_Handler,
		},
		{
			MethodName: "GetTemperatureAndPressure",
			Handler:    _AtmosphereService_GetTemperatureAndPressure_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dcs/atmosphere/v0/atmosphere.proto",
}