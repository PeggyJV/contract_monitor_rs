// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: steward.proto

package steward_proto

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

// ContractCallServiceClient is the client API for ContractCallService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContractCallServiceClient interface {
	// Handles scheduled contract call submission
	Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type contractCallServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContractCallServiceClient(cc grpc.ClientConnInterface) ContractCallServiceClient {
	return &contractCallServiceClient{cc}
}

func (c *contractCallServiceClient) Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error) {
	out := new(ScheduleResponse)
	err := c.cc.Invoke(ctx, "/steward.v3.ContractCallService/Schedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractCallServiceClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/steward.v3.ContractCall/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContractCallServer is the server API for ContractCall service.
// All implementations must embed UnimplementedContractCallServer
// for forward compatibility
type ContractCallServiceServer interface {
	// Handles scheduled contract call submission
	Schedule(context.Context, *ScheduleRequest) (*ScheduleResponse, error)
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	mustEmbedUnimplementedContractCallServer()
}

// UnimplementedContractCallServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContractCallServiceServer struct {
}

func (UnimplementedContractCallServiceServer) Schedule(context.Context, *ScheduleRequest) (*ScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Schedule not implemented")
}
func (UnimplementedContractCallServiceServer) Status(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedContractCallServiceServer) mustEmbedUnimplementedContractCallServer() {}

// UnsafeContractCallServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContractCallServiceServer will
// result in compilation errors.
type UnsafeContractCallServiceServer interface {
	mustEmbedUnimplementedContractCallServiceServer()
}

func RegisterContractCallServiceServer(s grpc.ServiceRegistrar, srv ContractCallServiceServer) {
	s.RegisterService(&ContractCallService_ServiceDesc, srv)
}

func _ContractCallService_Schedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractCallServiceServer).Schedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/steward.v3.ContractCallService/Schedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractCallServiceServer).Schedule(ctx, req.(*ScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractCall_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractCallServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/steward.v3.ContractCall/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractCallServiceServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContractCall_ServiceDesc is the grpc.ServiceDesc for ContractCall service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContractCallService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "steward.v3.ContractCallService",
	HandlerType: (*ContractCallServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Schedule",
			Handler:    _ContractCallService_Schedule_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _ContractCall_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "steward.proto",
}
