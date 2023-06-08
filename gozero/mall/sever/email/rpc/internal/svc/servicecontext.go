package svc

import (
	log "github.com/go-ozzo/ozzo-log"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mall/common/mqtt"
	"mall/sever/email/model"
	"mall/sever/email/rpc/internal/config"
)

type ServiceContext struct {
	Logger    *log.Logger // 引入新的日志框架
	Config    config.Config
	MailMode  model.MailModel
	C2CClient mqtt.Messenger
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Logger:    log.NewLogger(),
		Config:    c,
		MailMode:  model.NewMailModel(conn, c.CacheRedis),
		C2CClient: mqtt.NewMessengerImpl(c.C2cMqttConf),
	}
}
