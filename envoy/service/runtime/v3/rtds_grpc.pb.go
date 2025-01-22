// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.2
// source: envoy/service/runtime/v3/rtds.proto

package runtimev3

import (
	context "context"
	v3 "github.com/sefaphlvn/versioned-go-control-plane/envoy/service/discovery/v3"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RuntimeDiscoveryService_StreamRuntime_FullMethodName = "/envoy.service.runtime.v3.RuntimeDiscoveryService/StreamRuntime"
	RuntimeDiscoveryService_DeltaRuntime_FullMethodName  = "/envoy.service.runtime.v3.RuntimeDiscoveryService/DeltaRuntime"
	RuntimeDiscoveryService_FetchRuntime_FullMethodName  = "/envoy.service.runtime.v3.RuntimeDiscoveryService/FetchRuntime"
)

// RuntimeDiscoveryServiceClient is the client API for RuntimeDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RuntimeDiscoveryServiceClient interface {
	StreamRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_StreamRuntimeClient, error)
	DeltaRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_DeltaRuntimeClient, error)
	FetchRuntime(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error)
}

type runtimeDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRuntimeDiscoveryServiceClient(cc grpc.ClientConnInterface) RuntimeDiscoveryServiceClient {
	return &runtimeDiscoveryServiceClient{cc}
}

func (c *runtimeDiscoveryServiceClient) StreamRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_StreamRuntimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &RuntimeDiscoveryService_ServiceDesc.Streams[0], RuntimeDiscoveryService_StreamRuntime_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &runtimeDiscoveryServiceStreamRuntimeClient{stream}
	return x, nil
}

type RuntimeDiscoveryService_StreamRuntimeClient interface {
	Send(*v3.DiscoveryRequest) error
	Recv() (*v3.DiscoveryResponse, error)
	grpc.ClientStream
}

type runtimeDiscoveryServiceStreamRuntimeClient struct {
	grpc.ClientStream
}

func (x *runtimeDiscoveryServiceStreamRuntimeClient) Send(m *v3.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceStreamRuntimeClient) Recv() (*v3.DiscoveryResponse, error) {
	m := new(v3.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *runtimeDiscoveryServiceClient) DeltaRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_DeltaRuntimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &RuntimeDiscoveryService_ServiceDesc.Streams[1], RuntimeDiscoveryService_DeltaRuntime_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &runtimeDiscoveryServiceDeltaRuntimeClient{stream}
	return x, nil
}

type RuntimeDiscoveryService_DeltaRuntimeClient interface {
	Send(*v3.DeltaDiscoveryRequest) error
	Recv() (*v3.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type runtimeDiscoveryServiceDeltaRuntimeClient struct {
	grpc.ClientStream
}

func (x *runtimeDiscoveryServiceDeltaRuntimeClient) Send(m *v3.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceDeltaRuntimeClient) Recv() (*v3.DeltaDiscoveryResponse, error) {
	m := new(v3.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *runtimeDiscoveryServiceClient) FetchRuntime(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error) {
	out := new(v3.DiscoveryResponse)
	err := c.cc.Invoke(ctx, RuntimeDiscoveryService_FetchRuntime_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RuntimeDiscoveryServiceServer is the server API for RuntimeDiscoveryService service.
// All implementations should embed UnimplementedRuntimeDiscoveryServiceServer
// for forward compatibility
type RuntimeDiscoveryServiceServer interface {
	StreamRuntime(RuntimeDiscoveryService_StreamRuntimeServer) error
	DeltaRuntime(RuntimeDiscoveryService_DeltaRuntimeServer) error
	FetchRuntime(context.Context, *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error)
}

// UnimplementedRuntimeDiscoveryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRuntimeDiscoveryServiceServer struct {
}

func (UnimplementedRuntimeDiscoveryServiceServer) StreamRuntime(RuntimeDiscoveryService_StreamRuntimeServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRuntime not implemented")
}
func (UnimplementedRuntimeDiscoveryServiceServer) DeltaRuntime(RuntimeDiscoveryService_DeltaRuntimeServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaRuntime not implemented")
}
func (UnimplementedRuntimeDiscoveryServiceServer) FetchRuntime(context.Context, *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchRuntime not implemented")
}

// UnsafeRuntimeDiscoveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RuntimeDiscoveryServiceServer will
// result in compilation errors.
type UnsafeRuntimeDiscoveryServiceServer interface {
	mustEmbedUnimplementedRuntimeDiscoveryServiceServer()
}

func RegisterRuntimeDiscoveryServiceServer(s grpc.ServiceRegistrar, srv RuntimeDiscoveryServiceServer) {
	s.RegisterService(&RuntimeDiscoveryService_ServiceDesc, srv)
}

func _RuntimeDiscoveryService_StreamRuntime_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RuntimeDiscoveryServiceServer).StreamRuntime(&runtimeDiscoveryServiceStreamRuntimeServer{stream})
}

type RuntimeDiscoveryService_StreamRuntimeServer interface {
	Send(*v3.DiscoveryResponse) error
	Recv() (*v3.DiscoveryRequest, error)
	grpc.ServerStream
}

type runtimeDiscoveryServiceStreamRuntimeServer struct {
	grpc.ServerStream
}

func (x *runtimeDiscoveryServiceStreamRuntimeServer) Send(m *v3.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceStreamRuntimeServer) Recv() (*v3.DiscoveryRequest, error) {
	m := new(v3.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RuntimeDiscoveryService_DeltaRuntime_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RuntimeDiscoveryServiceServer).DeltaRuntime(&runtimeDiscoveryServiceDeltaRuntimeServer{stream})
}

type RuntimeDiscoveryService_DeltaRuntimeServer interface {
	Send(*v3.DeltaDiscoveryResponse) error
	Recv() (*v3.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type runtimeDiscoveryServiceDeltaRuntimeServer struct {
	grpc.ServerStream
}

func (x *runtimeDiscoveryServiceDeltaRuntimeServer) Send(m *v3.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceDeltaRuntimeServer) Recv() (*v3.DeltaDiscoveryRequest, error) {
	m := new(v3.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RuntimeDiscoveryService_FetchRuntime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuntimeDiscoveryServiceServer).FetchRuntime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuntimeDiscoveryService_FetchRuntime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuntimeDiscoveryServiceServer).FetchRuntime(ctx, req.(*v3.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RuntimeDiscoveryService_ServiceDesc is the grpc.ServiceDesc for RuntimeDiscoveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RuntimeDiscoveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.runtime.v3.RuntimeDiscoveryService",
	HandlerType: (*RuntimeDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchRuntime",
			Handler:    _RuntimeDiscoveryService_FetchRuntime_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRuntime",
			Handler:       _RuntimeDiscoveryService_StreamRuntime_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaRuntime",
			Handler:       _RuntimeDiscoveryService_DeltaRuntime_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/runtime/v3/rtds.proto",
}
