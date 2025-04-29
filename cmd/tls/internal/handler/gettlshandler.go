package handler

import (
	"net/http"
	"strings"

	"esay-tls-server/cmd/tls/internal/logic"
	"esay-tls-server/cmd/tls/internal/svc"
	"esay-tls-server/cmd/tls/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func getTlsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqDomains struct {
			Domains string `form:"domains"`
		}
		if err := httpx.Parse(r, &reqDomains); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		idStrings := strings.Split(reqDomains.Domains, ",")
		req := types.GetTlsReq{
			Domains: idStrings,
		}
		l := logic.NewGetTlsLogic(r.Context(), svcCtx)
		resp, err := l.GetTls(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
