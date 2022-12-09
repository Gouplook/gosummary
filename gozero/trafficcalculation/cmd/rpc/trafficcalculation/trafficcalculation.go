// Code generated by goctl. DO NOT EDIT!
// Source: calculation.proto

package trafficcalculation

import (
	"context"

	"gosummary/gozero/trafficcalculation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Box               = pb.Box
	LaneStatus        = pb.LaneStatus
	QueryReq          = pb.QueryReq
	TrafficStatusResp = pb.TrafficStatusResp

	TrafficCalculation interface {
		//  查询拥堵状态
		QueryCurrentTraffic(ctx context.Context, in *QueryReq, opts ...grpc.CallOption) (*TrafficStatusResp, error)
	}

	defaultTrafficCalculation struct {
		cli zrpc.Client
	}
)

func NewTrafficCalculation(cli zrpc.Client) TrafficCalculation {
	return &defaultTrafficCalculation{
		cli: cli,
	}
}

//  查询拥堵状态
func (m *defaultTrafficCalculation) QueryCurrentTraffic(ctx context.Context, in *QueryReq, opts ...grpc.CallOption) (*TrafficStatusResp, error) {
	client := pb.NewTrafficCalculationClient(m.cli.Conn())
	return client.QueryCurrentTraffic(ctx, in, opts...)
}