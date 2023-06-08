package main

import (
	"flag"
	"fmt"
	"mall/sever/email/rpc/internal/listen"

	"mall/sever/email/rpc/internal/config"
	"mall/sever/email/rpc/internal/server"
	"mall/sever/email/rpc/internal/svc"
	"mall/sever/email/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/email.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	// sever rpc/mqtt
	svr := server.NewEmailServerServer(ctx)
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterEmailServerServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	serviceGroup.Add(s)
	// mqtt
	for _, mq := range listen.Mqtts(ctx) {
		serviceGroup.Add(mq)
	}
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	serviceGroup.Start()
}
