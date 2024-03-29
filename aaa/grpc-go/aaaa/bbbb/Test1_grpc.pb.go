// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.5.0
// source: Test1.proto

package bbbb

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

// Test1Client is the client API for Test1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Test1Client interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type test1Client struct {
	cc grpc.ClientConnInterface
}

func NewTest1Client(cc grpc.ClientConnInterface) Test1Client {
	return &test1Client{cc}
}

func (c *test1Client) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/aaaa.bbbb.Test1/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Test1Server is the server API for Test1 service.
// All implementations must embed UnimplementedTest1Server
// for forward compatibility
type Test1Server interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedTest1Server()
}

// UnimplementedTest1Server must be embedded to have forward compatible implementations.
type UnimplementedTest1Server struct {
}

func (UnimplementedTest1Server) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedTest1Server) mustEmbedUnimplementedTest1Server() {}

// UnsafeTest1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Test1Server will
// result in compilation errors.
type UnsafeTest1Server interface {
	mustEmbedUnimplementedTest1Server()
}

func RegisterTest1Server(s grpc.ServiceRegistrar, srv Test1Server) {
	s.RegisterService(&Test1_ServiceDesc, srv)
}

func _Test1_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Test1Server).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aaaa.bbbb.Test1/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Test1Server).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Test1_ServiceDesc is the grpc.ServiceDesc for Test1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Test1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aaaa.bbbb.Test1",
	HandlerType: (*Test1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Test1_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Test1.proto",
}
