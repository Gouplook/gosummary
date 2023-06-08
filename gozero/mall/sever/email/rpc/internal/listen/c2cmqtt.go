package listen

import (
	"context"
	"fmt"
	"mall/sever/email/rpc/internal/svc"
)

type C2cMqtt struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewC2cMqtt(ctx context.Context, svcCtx *svc.ServiceContext) *C2cMqtt {
	return &C2cMqtt{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (c *C2cMqtt) Start() {
	// 调用handle
	fmt.Println("C2cMqtt start ")
	_, err := NewC2cHandler(c.ctx, c.svcCtx)
	if err != nil {
		c.svcCtx.Logger.Error("Start C2cMqtt with error:%v", err.Error())
		return
	}
	select {}
}

func (l *C2cMqtt) Stop() {
	fmt.Println("C2cMqtt stop")
}
