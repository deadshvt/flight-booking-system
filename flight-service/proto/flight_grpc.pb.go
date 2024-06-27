// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.26.1
// source: flight-service/proto/flight.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	FlightService_GetFlightsWithAirports_FullMethodName = "/flight.FlightService/GetFlightsWithAirports"
	FlightService_GetFlightWithAirports_FullMethodName  = "/flight.FlightService/GetFlightWithAirports"
	FlightService_CreateFlight_FullMethodName           = "/flight.FlightService/CreateFlight"
	FlightService_CreateAirport_FullMethodName          = "/flight.FlightService/CreateAirport"
)

// FlightServiceClient is the client API for FlightService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlightServiceClient interface {
	GetFlightsWithAirports(ctx context.Context, in *GetFlightsWithAirportsRequest, opts ...grpc.CallOption) (*GetFlightsWithAirportsResponse, error)
	GetFlightWithAirports(ctx context.Context, in *GetFlightWithAirportsRequest, opts ...grpc.CallOption) (*GetFlightWithAirportsResponse, error)
	CreateFlight(ctx context.Context, in *CreateFlightRequest, opts ...grpc.CallOption) (*CreateFlightResponse, error)
	CreateAirport(ctx context.Context, in *CreateAirportRequest, opts ...grpc.CallOption) (*CreateAirportResponse, error)
}

type flightServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFlightServiceClient(cc grpc.ClientConnInterface) FlightServiceClient {
	return &flightServiceClient{cc}
}

func (c *flightServiceClient) GetFlightsWithAirports(ctx context.Context, in *GetFlightsWithAirportsRequest, opts ...grpc.CallOption) (*GetFlightsWithAirportsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFlightsWithAirportsResponse)
	err := c.cc.Invoke(ctx, FlightService_GetFlightsWithAirports_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightServiceClient) GetFlightWithAirports(ctx context.Context, in *GetFlightWithAirportsRequest, opts ...grpc.CallOption) (*GetFlightWithAirportsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFlightWithAirportsResponse)
	err := c.cc.Invoke(ctx, FlightService_GetFlightWithAirports_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightServiceClient) CreateFlight(ctx context.Context, in *CreateFlightRequest, opts ...grpc.CallOption) (*CreateFlightResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateFlightResponse)
	err := c.cc.Invoke(ctx, FlightService_CreateFlight_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightServiceClient) CreateAirport(ctx context.Context, in *CreateAirportRequest, opts ...grpc.CallOption) (*CreateAirportResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAirportResponse)
	err := c.cc.Invoke(ctx, FlightService_CreateAirport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlightServiceServer is the server API for FlightService service.
// All implementations must embed UnimplementedFlightServiceServer
// for forward compatibility
type FlightServiceServer interface {
	GetFlightsWithAirports(context.Context, *GetFlightsWithAirportsRequest) (*GetFlightsWithAirportsResponse, error)
	GetFlightWithAirports(context.Context, *GetFlightWithAirportsRequest) (*GetFlightWithAirportsResponse, error)
	CreateFlight(context.Context, *CreateFlightRequest) (*CreateFlightResponse, error)
	CreateAirport(context.Context, *CreateAirportRequest) (*CreateAirportResponse, error)
	mustEmbedUnimplementedFlightServiceServer()
}

// UnimplementedFlightServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFlightServiceServer struct {
}

func (UnimplementedFlightServiceServer) GetFlightsWithAirports(context.Context, *GetFlightsWithAirportsRequest) (*GetFlightsWithAirportsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlightsWithAirports not implemented")
}
func (UnimplementedFlightServiceServer) GetFlightWithAirports(context.Context, *GetFlightWithAirportsRequest) (*GetFlightWithAirportsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlightWithAirports not implemented")
}
func (UnimplementedFlightServiceServer) CreateFlight(context.Context, *CreateFlightRequest) (*CreateFlightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFlight not implemented")
}
func (UnimplementedFlightServiceServer) CreateAirport(context.Context, *CreateAirportRequest) (*CreateAirportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAirport not implemented")
}
func (UnimplementedFlightServiceServer) mustEmbedUnimplementedFlightServiceServer() {}

// UnsafeFlightServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlightServiceServer will
// result in compilation errors.
type UnsafeFlightServiceServer interface {
	mustEmbedUnimplementedFlightServiceServer()
}

func RegisterFlightServiceServer(s grpc.ServiceRegistrar, srv FlightServiceServer) {
	s.RegisterService(&FlightService_ServiceDesc, srv)
}

func _FlightService_GetFlightsWithAirports_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFlightsWithAirportsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceServer).GetFlightsWithAirports(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlightService_GetFlightsWithAirports_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceServer).GetFlightsWithAirports(ctx, req.(*GetFlightsWithAirportsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightService_GetFlightWithAirports_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFlightWithAirportsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceServer).GetFlightWithAirports(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlightService_GetFlightWithAirports_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceServer).GetFlightWithAirports(ctx, req.(*GetFlightWithAirportsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightService_CreateFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFlightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceServer).CreateFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlightService_CreateFlight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceServer).CreateFlight(ctx, req.(*CreateFlightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightService_CreateAirport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAirportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceServer).CreateAirport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlightService_CreateAirport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceServer).CreateAirport(ctx, req.(*CreateAirportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FlightService_ServiceDesc is the grpc.ServiceDesc for FlightService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlightService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flight.FlightService",
	HandlerType: (*FlightServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFlightsWithAirports",
			Handler:    _FlightService_GetFlightsWithAirports_Handler,
		},
		{
			MethodName: "GetFlightWithAirports",
			Handler:    _FlightService_GetFlightWithAirports_Handler,
		},
		{
			MethodName: "CreateFlight",
			Handler:    _FlightService_CreateFlight_Handler,
		},
		{
			MethodName: "CreateAirport",
			Handler:    _FlightService_CreateAirport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flight-service/proto/flight.proto",
}
