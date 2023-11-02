// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: goServer.proto

package goServerTest

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

const (
	GoTestService_GetTestGO_FullMethodName = "/goServerTest.GoTestService/GetTestGO"
)

// GoTestServiceClient is the client API for GoTestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoTestServiceClient interface {
	GetTestGO(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error)
}

type goTestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoTestServiceClient(cc grpc.ClientConnInterface) GoTestServiceClient {
	return &goTestServiceClient{cc}
}

func (c *goTestServiceClient) GetTestGO(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error) {
	out := new(TestResponse)
	err := c.cc.Invoke(ctx, GoTestService_GetTestGO_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoTestServiceServer is the server API for GoTestService service.
// All implementations must embed UnimplementedGoTestServiceServer
// for forward compatibility
type GoTestServiceServer interface {
	GetTestGO(context.Context, *TestRequest) (*TestResponse, error)
	mustEmbedUnimplementedGoTestServiceServer()
}

// UnimplementedGoTestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGoTestServiceServer struct {
}

func (UnimplementedGoTestServiceServer) GetTestGO(context.Context, *TestRequest) (*TestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestGO not implemented")
}
func (UnimplementedGoTestServiceServer) mustEmbedUnimplementedGoTestServiceServer() {}

// UnsafeGoTestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoTestServiceServer will
// result in compilation errors.
type UnsafeGoTestServiceServer interface {
	mustEmbedUnimplementedGoTestServiceServer()
}

func RegisterGoTestServiceServer(s grpc.ServiceRegistrar, srv GoTestServiceServer) {
	s.RegisterService(&GoTestService_ServiceDesc, srv)
}

func _GoTestService_GetTestGO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoTestServiceServer).GetTestGO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoTestService_GetTestGO_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoTestServiceServer).GetTestGO(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoTestService_ServiceDesc is the grpc.ServiceDesc for GoTestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoTestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goServerTest.GoTestService",
	HandlerType: (*GoTestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTestGO",
			Handler:    _GoTestService_GetTestGO_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goServer.proto",
}