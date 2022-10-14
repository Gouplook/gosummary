package logic

import (
	"context"

	"gosummary/gozero/trafficcalculation/cmd/api/internal/svc"
	"gosummary/gozero/trafficcalculation/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTrafficStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTrafficStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTrafficStatusLogic {
	return &GetTrafficStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTrafficStatusLogic) GetTrafficStatus(req *types.TrafficStatusReq) (resp *types.TrafficStatusResp, err error) {
	// todo: add your logic here and delete this line

	return
}
