package v1

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mall/sever/email/api/internal/logic/v1"
	"mall/sever/email/api/internal/svc"
	"mall/sever/email/api/internal/types"
)

func SendEmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendArgs
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := v1.NewSendEmailLogic(r.Context(), svcCtx)
		resp, err := l.SendEmail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
