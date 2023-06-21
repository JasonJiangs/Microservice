// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.3
// source: person.proto

package __

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

// RpcServiceClient is the client API for RpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcServiceClient interface {
	HelloService(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Person, error)
}

type rpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcServiceClient(cc grpc.ClientConnInterface) RpcServiceClient {
	return &rpcServiceClient{cc}
}

func (c *rpcServiceClient) HelloService(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Person, error) {
	out := new(Person)
	err := c.cc.Invoke(ctx, "/pb.rpcService/HelloService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcServiceServer is the server API for RpcService service.
// All implementations must embed UnimplementedRpcServiceServer
// for forward compatibility
type RpcServiceServer interface {
	HelloService(context.Context, *Person) (*Person, error)
	mustEmbedUnimplementedRpcServiceServer()
}

// UnimplementedRpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRpcServiceServer struct {
}

func (UnimplementedRpcServiceServer) HelloService(context.Context, *Person) (*Person, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloService not implemented")
}
func (UnimplementedRpcServiceServer) mustEmbedUnimplementedRpcServiceServer() {}

// UnsafeRpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcServiceServer will
// result in compilation errors.
type UnsafeRpcServiceServer interface {
	mustEmbedUnimplementedRpcServiceServer()
}

func RegisterRpcServiceServer(s grpc.ServiceRegistrar, srv RpcServiceServer) {
	s.RegisterService(&RpcService_ServiceDesc, srv)
}

func _RpcService_HelloService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Person)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServiceServer).HelloService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.rpcService/HelloService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServiceServer).HelloService(ctx, req.(*Person))
	}
	return interceptor(ctx, in, info, handler)
}

// RpcService_ServiceDesc is the grpc.ServiceDesc for RpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.rpcService",
	HandlerType: (*RpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloService",
			Handler:    _RpcService_HelloService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "person.proto",
}
