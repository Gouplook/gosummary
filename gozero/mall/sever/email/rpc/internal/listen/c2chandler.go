package listen

import (
	"context"
	"fmt"
	"mall/common/mqtt"
	"mall/sever/email/rpc/internal/svc"
)

type C2cHandler struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewC2cHandler(ctx context.Context, svcCtx *svc.ServiceContext) (*C2cHandler, error) {
	handler := &C2cHandler{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
	err := handler.subscribe()
	if err != nil {
		return handler, err
	}
	return handler, nil
}

func (c *C2cHandler) subscribe() error {
	topics := []mqtt.Topic{{
		Name: fmt.Sprintf("%v/format_robot_location/+/+", c.svcCtx.Config.C2cMqttConf.TopicPrefix),
	}}
	var subHandler mqtt.SubscribeHandler = c.onHandle
	err := c.svcCtx.C2CClient.Subscribe(topics, &subHandler)
	if err != nil {
		c.svcCtx.Logger.Error("subscribe err .. ")
		return err
	}
	return nil
}

//
func (c *C2cHandler) onHandle(topic, payload string) {
	// todo  接到内如进行处理，可以调用逻辑
}
