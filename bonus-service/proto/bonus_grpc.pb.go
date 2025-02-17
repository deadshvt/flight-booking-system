// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.26.1
// source: bonus-service/proto/bonus.proto

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
	BonusService_GetPrivilegeWithHistory_FullMethodName = "/bonus.BonusService/GetPrivilegeWithHistory"
	BonusService_GetPrivilege_FullMethodName            = "/bonus.BonusService/GetPrivilege"
	BonusService_CreatePrivilege_FullMethodName         = "/bonus.BonusService/CreatePrivilege"
	BonusService_UpdatePrivilege_FullMethodName         = "/bonus.BonusService/UpdatePrivilege"
	BonusService_CreateOperation_FullMethodName         = "/bonus.BonusService/CreateOperation"
)

// BonusServiceClient is the client API for BonusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BonusServiceClient interface {
	GetPrivilegeWithHistory(ctx context.Context, in *GetPrivilegeWithHistoryRequest, opts ...grpc.CallOption) (*GetPrivilegeWithHistoryResponse, error)
	GetPrivilege(ctx context.Context, in *GetPrivilegeRequest, opts ...grpc.CallOption) (*GetPrivilegeResponse, error)
	CreatePrivilege(ctx context.Context, in *CreatePrivilegeRequest, opts ...grpc.CallOption) (*CreatePrivilegeResponse, error)
	UpdatePrivilege(ctx context.Context, in *UpdatePrivilegeRequest, opts ...grpc.CallOption) (*UpdatePrivilegeResponse, error)
	CreateOperation(ctx context.Context, in *CreateOperationRequest, opts ...grpc.CallOption) (*CreateOperationResponse, error)
}

type bonusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBonusServiceClient(cc grpc.ClientConnInterface) BonusServiceClient {
	return &bonusServiceClient{cc}
}

func (c *bonusServiceClient) GetPrivilegeWithHistory(ctx context.Context, in *GetPrivilegeWithHistoryRequest, opts ...grpc.CallOption) (*GetPrivilegeWithHistoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPrivilegeWithHistoryResponse)
	err := c.cc.Invoke(ctx, BonusService_GetPrivilegeWithHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bonusServiceClient) GetPrivilege(ctx context.Context, in *GetPrivilegeRequest, opts ...grpc.CallOption) (*GetPrivilegeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPrivilegeResponse)
	err := c.cc.Invoke(ctx, BonusService_GetPrivilege_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bonusServiceClient) CreatePrivilege(ctx context.Context, in *CreatePrivilegeRequest, opts ...grpc.CallOption) (*CreatePrivilegeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePrivilegeResponse)
	err := c.cc.Invoke(ctx, BonusService_CreatePrivilege_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bonusServiceClient) UpdatePrivilege(ctx context.Context, in *UpdatePrivilegeRequest, opts ...grpc.CallOption) (*UpdatePrivilegeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePrivilegeResponse)
	err := c.cc.Invoke(ctx, BonusService_UpdatePrivilege_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bonusServiceClient) CreateOperation(ctx context.Context, in *CreateOperationRequest, opts ...grpc.CallOption) (*CreateOperationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateOperationResponse)
	err := c.cc.Invoke(ctx, BonusService_CreateOperation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BonusServiceServer is the server API for BonusService service.
// All implementations must embed UnimplementedBonusServiceServer
// for forward compatibility
type BonusServiceServer interface {
	GetPrivilegeWithHistory(context.Context, *GetPrivilegeWithHistoryRequest) (*GetPrivilegeWithHistoryResponse, error)
	GetPrivilege(context.Context, *GetPrivilegeRequest) (*GetPrivilegeResponse, error)
	CreatePrivilege(context.Context, *CreatePrivilegeRequest) (*CreatePrivilegeResponse, error)
	UpdatePrivilege(context.Context, *UpdatePrivilegeRequest) (*UpdatePrivilegeResponse, error)
	CreateOperation(context.Context, *CreateOperationRequest) (*CreateOperationResponse, error)
	mustEmbedUnimplementedBonusServiceServer()
}

// UnimplementedBonusServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBonusServiceServer struct {
}

func (UnimplementedBonusServiceServer) GetPrivilegeWithHistory(context.Context, *GetPrivilegeWithHistoryRequest) (*GetPrivilegeWithHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrivilegeWithHistory not implemented")
}
func (UnimplementedBonusServiceServer) GetPrivilege(context.Context, *GetPrivilegeRequest) (*GetPrivilegeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrivilege not implemented")
}
func (UnimplementedBonusServiceServer) CreatePrivilege(context.Context, *CreatePrivilegeRequest) (*CreatePrivilegeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePrivilege not implemented")
}
func (UnimplementedBonusServiceServer) UpdatePrivilege(context.Context, *UpdatePrivilegeRequest) (*UpdatePrivilegeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePrivilege not implemented")
}
func (UnimplementedBonusServiceServer) CreateOperation(context.Context, *CreateOperationRequest) (*CreateOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOperation not implemented")
}
func (UnimplementedBonusServiceServer) mustEmbedUnimplementedBonusServiceServer() {}

// UnsafeBonusServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BonusServiceServer will
// result in compilation errors.
type UnsafeBonusServiceServer interface {
	mustEmbedUnimplementedBonusServiceServer()
}

func RegisterBonusServiceServer(s grpc.ServiceRegistrar, srv BonusServiceServer) {
	s.RegisterService(&BonusService_ServiceDesc, srv)
}

func _BonusService_GetPrivilegeWithHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrivilegeWithHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BonusServiceServer).GetPrivilegeWithHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BonusService_GetPrivilegeWithHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BonusServiceServer).GetPrivilegeWithHistory(ctx, req.(*GetPrivilegeWithHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BonusService_GetPrivilege_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrivilegeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BonusServiceServer).GetPrivilege(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BonusService_GetPrivilege_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BonusServiceServer).GetPrivilege(ctx, req.(*GetPrivilegeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BonusService_CreatePrivilege_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePrivilegeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BonusServiceServer).CreatePrivilege(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BonusService_CreatePrivilege_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BonusServiceServer).CreatePrivilege(ctx, req.(*CreatePrivilegeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BonusService_UpdatePrivilege_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePrivilegeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BonusServiceServer).UpdatePrivilege(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BonusService_UpdatePrivilege_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BonusServiceServer).UpdatePrivilege(ctx, req.(*UpdatePrivilegeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BonusService_CreateOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BonusServiceServer).CreateOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BonusService_CreateOperation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BonusServiceServer).CreateOperation(ctx, req.(*CreateOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BonusService_ServiceDesc is the grpc.ServiceDesc for BonusService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BonusService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bonus.BonusService",
	HandlerType: (*BonusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPrivilegeWithHistory",
			Handler:    _BonusService_GetPrivilegeWithHistory_Handler,
		},
		{
			MethodName: "GetPrivilege",
			Handler:    _BonusService_GetPrivilege_Handler,
		},
		{
			MethodName: "CreatePrivilege",
			Handler:    _BonusService_CreatePrivilege_Handler,
		},
		{
			MethodName: "UpdatePrivilege",
			Handler:    _BonusService_UpdatePrivilege_Handler,
		},
		{
			MethodName: "CreateOperation",
			Handler:    _BonusService_CreateOperation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bonus-service/proto/bonus.proto",
}
