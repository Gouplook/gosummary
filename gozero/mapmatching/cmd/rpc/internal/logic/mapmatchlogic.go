package logic

import (
	"context"

	"gosummary/gozero/mapmatching/cmd/rpc/internal/svc"
	"gosummary/gozero/mapmatching/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MapMatchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMapMatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MapMatchLogic {
	return &MapMatchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MapMatchLogic) MapMatch(in *pb.MapMatchReq) (*pb.MapMatchResult, error) {
	// todo: add your logic here and delete this line

	return &pb.MapMatchResult{}, nil
}
