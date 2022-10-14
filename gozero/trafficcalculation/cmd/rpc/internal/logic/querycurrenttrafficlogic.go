package logic

import (
	"context"

	"gosummary/gozero/trafficcalculation/cmd/rpc/internal/svc"
	"gosummary/gozero/trafficcalculation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCurrentTrafficLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCurrentTrafficLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCurrentTrafficLogic {
	return &QueryCurrentTrafficLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  查询拥堵状态
func (l *QueryCurrentTrafficLogic) QueryCurrentTraffic(in *pb.QueryReq) (*pb.TrafficStatusResp, error) {
	// todo: add your logic here and delete this line

	return &pb.TrafficStatusResp{}, nil
}
