package v1

import (
	"context"
	"fmt"
	"mall/sever/email/api/internal/svc"
	"mall/sever/email/api/internal/types"
	"mall/sever/email/rpc/emailserver"

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

	args := new(emailserver.SendRequest)
	args.ToEmail = "yinjinlin_uplook@163.com"

	//调用RPC服务
	rs, _ := l.svcCtx.EmailRpc.SendEmailRpc(l.ctx, args)
	fmt.Println(rs.Code)

	resp.Code = 5000
	resp.Msg = "success"

	return
}
