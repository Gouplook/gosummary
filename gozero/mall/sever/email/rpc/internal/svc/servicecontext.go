package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mall/sever/email/model"
	"mall/sever/email/rpc/internal/config"
)

type ServiceContext struct {
	Config   config.Config
	MailMode model.MailModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:   c,
		MailMode: model.NewMailModel(conn, c.CacheRedis),
	}
}
