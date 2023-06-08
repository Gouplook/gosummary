package v1

import (
	"context"
	"errors"
	"mall/sever/email/model"

	"mall/sever/email/api/internal/svc"
	"mall/sever/email/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindEmailLogic {
	return &FindEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindEmailLogic) FindEmail(req *types.FindArgs) (resp *types.FindReply, err error) {

	resp = new(types.FindReply)

	var id int64
	id = req.Id
	// 调用mode,主要用于查找
	mail, err := l.svcCtx.MailMode.FindOne(l.ctx, id)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errors.New("邮箱不存在")

	default:
		return nil, err

	}
	resp.ToMail = mail.ToEmail.String
	resp.Code = 5000
	resp.Msg = "success"

	return
}
