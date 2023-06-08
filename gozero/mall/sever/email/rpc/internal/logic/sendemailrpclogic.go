package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"mall/sever/email/rpc/internal/svc"
	"mall/sever/email/rpc/pb/pb"
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
	// 调用model 写数据，需要对写数据进行加锁
	resp := new(pb.SendResponse)
	//data := new(model.Mail)
	//data.Name.String = "10600"
	//data.Name.Valid = true
	//data.ToEmail.String = "yaxuan.kang@uisee.com"
	//data.ToEmail.Valid = true
	//data.CreatedAt.Int64 = time.Now().Unix()
	//data.CreatedAt.Valid = true // 需要赋值
	//data.IsSend = 1

	// 插入带缓存
	//_, err := l.svcCtx.MailMode.Insert(l.ctx, data)
	//if err != nil {
	//	l.Logger.Error("MailModel Insert error")
	//	return resp, err
	//}

	// 插入不带缓存
	//_, err := l.svcCtx.MailMode.InsertNoCache(l.ctx, data)
	//if err != nil {
	//	l.Logger.Error("MailModel Insert error")
	//	return resp, err
	//}

	// 批量插入
	//l.svcCtx.MailMode.Insert()

	// 查询
	rs, _ := l.svcCtx.MailMode.FindOne(l.ctx, 7)
	fmt.Println(rs.ToEmail)

	// 分页查询(不带缓存）
	//pageNum := 1
	//pageSize := 10
	//where := map[string]interface{}{
	//	"is_send": 0,
	//}
	//rs, err := l.svcCtx.MailMode.FindsNoCache(l.ctx, pageNum, pageSize, where)
	//if err != nil {
	//	l.Logger.Error("MailModel Insert error")
	//}
	//
	//fmt.Println(rs)

	resp.Code = "4006"
	return resp, nil
}
