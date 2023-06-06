package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"mall/sever/email/api/internal/config"
	"mall/sever/email/model"
	"mall/sever/email/rpc/emailserver"
)

type ServiceContext struct {
	Config   config.Config
	EmailRpc emailserver.EmailServer // 引用rpc接口
	MailMode model.MailModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:   c,
		EmailRpc: emailserver.NewEmailServer(zrpc.MustNewClient(c.EmailRpc)),
		MailMode: model.NewMailModel(conn, c.CacheRedis),
	}
}
