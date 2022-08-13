// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: proto/reverseNumber/reverseNumber.proto

package proto

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

// ReverseNumberClient is the client API for ReverseNumber service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReverseNumberClient interface {
	ReverseNumber(ctx context.Context, in *Number, opts ...grpc.CallOption) (*Reverse, error)
}

type reverseNumberClient struct {
	cc grpc.ClientConnInterface
}

func NewReverseNumberClient(cc grpc.ClientConnInterface) ReverseNumberClient {
	return &reverseNumberClient{cc}
}

func (c *reverseNumberClient) ReverseNumber(ctx context.Context, in *Number, opts ...grpc.CallOption) (*Reverse, error) {
	out := new(Reverse)
	err := c.cc.Invoke(ctx, "/ReverseNumber/ReverseNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReverseNumberServer is the server API for ReverseNumber service.
// All implementations must embed UnimplementedReverseNumberServer
// for forward compatibility
type ReverseNumberServer interface {
	ReverseNumber(context.Context, *Number) (*Reverse, error)
	mustEmbedUnimplementedReverseNumberServer()
}

// UnimplementedReverseNumberServer must be embedded to have forward compatible implementations.
type UnimplementedReverseNumberServer struct {
}

func (UnimplementedReverseNumberServer) ReverseNumber(context.Context, *Number) (*Reverse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReverseNumber not implemented")
}
func (UnimplementedReverseNumberServer) mustEmbedUnimplementedReverseNumberServer() {}

// UnsafeReverseNumberServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReverseNumberServer will
// result in compilation errors.
type UnsafeReverseNumberServer interface {
	mustEmbedUnimplementedReverseNumberServer()
}

func RegisterReverseNumberServer(s grpc.ServiceRegistrar, srv ReverseNumberServer) {
	s.RegisterService(&ReverseNumber_ServiceDesc, srv)
}

func _ReverseNumber_ReverseNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Number)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReverseNumberServer).ReverseNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ReverseNumber/ReverseNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReverseNumberServer).ReverseNumber(ctx, req.(*Number))
	}
	return interceptor(ctx, in, info, handler)
}

// ReverseNumber_ServiceDesc is the grpc.ServiceDesc for ReverseNumber service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReverseNumber_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ReverseNumber",
	HandlerType: (*ReverseNumberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReverseNumber",
			Handler:    _ReverseNumber_ReverseNumber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/reverseNumber/reverseNumber.proto",
}
