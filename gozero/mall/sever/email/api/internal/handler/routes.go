// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	v1 "mall/sever/email/api/internal/handler/v1"
	"mall/sever/email/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/send",
				Handler: v1.SendEmailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/find",
				Handler: v1.FindEmailHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/email"),
	)
}
