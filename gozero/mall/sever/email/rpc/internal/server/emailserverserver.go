// Code generated by goctl. DO NOT EDIT!
// Source: email.proto

package server

import (
	"context"

	"mall/sever/email/rpc/internal/logic"
	"mall/sever/email/rpc/internal/svc"
	"mall/sever/email/rpc/pb/pb"
)

type EmailServerServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedEmailServerServer
}

func NewEmailServerServer(svcCtx *svc.ServiceContext) *EmailServerServer {
	return &EmailServerServer{
		svcCtx: svcCtx,
	}
}

func (s *EmailServerServer) SendEmailRpc(ctx context.Context, in *pb.SendRequest) (*pb.SendResponse, error) {
	l := logic.NewSendEmailRpcLogic(ctx, s.svcCtx)
	return l.SendEmailRpc(in)
}
