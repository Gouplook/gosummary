package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gosummary/gozero/trafficcalculation/cmd/api/internal/logic"
	"gosummary/gozero/trafficcalculation/cmd/api/internal/svc"
	"gosummary/gozero/trafficcalculation/cmd/api/internal/types"
)

func GetTrafficStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TrafficStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetTrafficStatusLogic(r.Context(), svcCtx)
		resp, err := l.GetTrafficStatus(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
