// Code generated by goctl. DO NOT EDIT!
// Source: mapmatching.proto

package server

import (
	"context"

	"gosummary/gozero/mapmatching/cmd/rpc/internal/logic"
	"gosummary/gozero/mapmatching/cmd/rpc/internal/svc"
	"gosummary/gozero/mapmatching/cmd/rpc/pb"
)

type MapMatchingServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedMapMatchingServer
}

func NewMapMatchingServer(svcCtx *svc.ServiceContext) *MapMatchingServer {
	return &MapMatchingServer{
		svcCtx: svcCtx,
	}
}

func (s *MapMatchingServer) MapMatch(ctx context.Context, in *pb.MapMatchReq) (*pb.MapMatchResult, error) {
	l := logic.NewMapMatchLogic(ctx, s.svcCtx)
	return l.MapMatch(in)
}
