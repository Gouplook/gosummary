package main

import (
	"flag"
	"fmt"

	"gosummary/gozero/trafficcalculation/cmd/rpc/internal/config"
	"gosummary/gozero/trafficcalculation/cmd/rpc/internal/server"
	"gosummary/gozero/trafficcalculation/cmd/rpc/internal/svc"
	"gosummary/gozero/trafficcalculation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/calculation.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewTrafficCalculationServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterTrafficCalculationServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	serviceGroup := service.NewServiceGroup()
	serviceGroup.Add(s)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)

	// rpc/mqtt
	//for _,mq := range
	defer serviceGroup.Stop()
	serviceGroup.Start()

}
