package logic

import (
	"context"

	"mall/sever/email/rpc/internal/svc"
	"mall/sever/email/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailRpcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendEmailRpcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailRpcLogic {
	return &SendEmailRpcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendEmailRpcLogic) SendEmailRpc(in *pb.SendRequest) (*pb.SendResponse, error) {
	l.Logger.Error()

	return &pb.SendResponse{
		Code: "4006",
	}, nil
}
