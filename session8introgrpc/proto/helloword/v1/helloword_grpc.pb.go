// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/helloword/v1/helloword.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GreaterService_SayHello_FullMethodName = "/proto.helloword.v1.GreaterService/SayHello"
)

// GreaterServiceClient is the client API for GreaterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreaterServiceClient interface {
	SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error)
}

type greaterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreaterServiceClient(cc grpc.ClientConnInterface) GreaterServiceClient {
	return &greaterServiceClient{cc}
}

func (c *greaterServiceClient) SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SayHelloResponse)
	err := c.cc.Invoke(ctx, GreaterService_SayHello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreaterServiceServer is the server API for GreaterService service.
// All implementations must embed UnimplementedGreaterServiceServer
// for forward compatibility.
type GreaterServiceServer interface {
	SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error)
	mustEmbedUnimplementedGreaterServiceServer()
}

// UnimplementedGreaterServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGreaterServiceServer struct{}

func (UnimplementedGreaterServiceServer) SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreaterServiceServer) mustEmbedUnimplementedGreaterServiceServer() {}
func (UnimplementedGreaterServiceServer) testEmbeddedByValue()                        {}

// UnsafeGreaterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreaterServiceServer will
// result in compilation errors.
type UnsafeGreaterServiceServer interface {
	mustEmbedUnimplementedGreaterServiceServer()
}

func RegisterGreaterServiceServer(s grpc.ServiceRegistrar, srv GreaterServiceServer) {
	// If the following call pancis, it indicates UnimplementedGreaterServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GreaterService_ServiceDesc, srv)
}

func _GreaterService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreaterServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GreaterService_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreaterServiceServer).SayHello(ctx, req.(*SayHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GreaterService_ServiceDesc is the grpc.ServiceDesc for GreaterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreaterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.helloword.v1.GreaterService",
	HandlerType: (*GreaterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _GreaterService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/helloword/v1/helloword.proto",
}
