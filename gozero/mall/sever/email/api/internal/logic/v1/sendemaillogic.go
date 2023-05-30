package v1

import (
	"context"
	"fmt"

	"mall/sever/email/api/internal/svc"
	"mall/sever/email/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendEmailLogic) SendEmail(req *types.SendArgs) (resp *types.SendReply, err error) {
	resp = new(types.SendReply)
	resp.Code = 5000
	resp.Msg = "success"
	fmt.Println("====")
	l.Logger.Error("error")
	fmt.Sprintf("=============")
	l.Logger.Debug("debug")

	return
}
