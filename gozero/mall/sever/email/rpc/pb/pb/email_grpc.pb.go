// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: email.proto

package pb

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

// EmailServerClient is the client API for EmailServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailServerClient interface {
	SendEmailRpc(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
}

type emailServerClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailServerClient(cc grpc.ClientConnInterface) EmailServerClient {
	return &emailServerClient{cc}
}

func (c *emailServerClient) SendEmailRpc(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/pb.EmailServer/SendEmailRpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServerServer is the server API for EmailServer service.
// All implementations must embed UnimplementedEmailServerServer
// for forward compatibility
type EmailServerServer interface {
	SendEmailRpc(context.Context, *SendRequest) (*SendResponse, error)
	mustEmbedUnimplementedEmailServerServer()
}

// UnimplementedEmailServerServer must be embedded to have forward compatible implementations.
type UnimplementedEmailServerServer struct {
}

func (UnimplementedEmailServerServer) SendEmailRpc(context.Context, *SendRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmailRpc not implemented")
}
func (UnimplementedEmailServerServer) mustEmbedUnimplementedEmailServerServer() {}

// UnsafeEmailServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailServerServer will
// result in compilation errors.
type UnsafeEmailServerServer interface {
	mustEmbedUnimplementedEmailServerServer()
}

func RegisterEmailServerServer(s grpc.ServiceRegistrar, srv EmailServerServer) {
	s.RegisterService(&EmailServer_ServiceDesc, srv)
}

func _EmailServer_SendEmailRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServerServer).SendEmailRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EmailServer/SendEmailRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServerServer).SendEmailRpc(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EmailServer_ServiceDesc is the grpc.ServiceDesc for EmailServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmailServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.EmailServer",
	HandlerType: (*EmailServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmailRpc",
			Handler:    _EmailServer_SendEmailRpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email.proto",
}
