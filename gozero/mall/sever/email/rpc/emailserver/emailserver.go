// Code generated by goctl. DO NOT EDIT!
// Source: email.proto

package emailserver

import (
	"context"

	"mall/sever/email/rpc/pb/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	SendRequest  = pb.SendRequest
	SendResponse = pb.SendResponse

	EmailServer interface {
		SendEmailRpc(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
	}

	defaultEmailServer struct {
		cli zrpc.Client
	}
)

func NewEmailServer(cli zrpc.Client) EmailServer {
	return &defaultEmailServer{
		cli: cli,
	}
}

func (m *defaultEmailServer) SendEmailRpc(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	client := pb.NewEmailServerClient(m.cli.Conn())
	return client.SendEmailRpc(ctx, in, opts...)
}
