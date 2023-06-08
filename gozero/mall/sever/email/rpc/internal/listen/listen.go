package listen

import (
	"context"
	"mall/sever/email/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/service"
)

// Mqtts 返回所有消费者
func Mqtts(serverCtx *svc.ServiceContext) []service.Service {
	ctx := context.Background()
	var services []service.Service
	services = append(services,
		NewC2cMqtt(ctx, serverCtx),
	)
	return services
}
