package main

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"mall/sever/email/rpc/internal/config"
	"mall/sever/email/rpc/internal/logic"
	"mall/sever/email/rpc/internal/svc"
	"mall/sever/email/rpc/pb/pb"
	"testing"
)

func TestSendEmail(t *testing.T) {

	var c config.Config
	conf.MustLoad("D:\\gosummaryCode\\gosummary\\gozero\\mall\\sever\\email\\rpc\\etc\\email.yaml", &c)
	ctx := svc.NewServiceContext(c)
	// 初始化
	l := logic.NewSendEmailRpcLogic(context.Background(), ctx)
	args := new(pb.SendRequest)

	rs, err := l.SendEmailRpc(args)
	if err != nil {
		t.Logf("rpc sendEmail err")
	}

	fmt.Println(rs)
}
