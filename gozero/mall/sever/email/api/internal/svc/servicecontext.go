package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mall/sever/email/api/internal/config"
	"mall/sever/email/rpc/emailserver"
)

type ServiceContext struct {
	Config   config.Config
	EmailRpc emailserver.EmailServer // 引用rpc接口
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		EmailRpc: emailserver.NewEmailServer(zrpc.MustNewClient(c.EmailRpc)),
	}
}
