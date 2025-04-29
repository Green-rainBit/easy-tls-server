package handler

import (
	"net/http"

	"esay-tls-server/cmd/tls/internal/logic"
	"esay-tls-server/cmd/tls/internal/svc"
	"esay-tls-server/cmd/tls/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func createTlsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostCreateTlsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateTlsLogic(r.Context(), svcCtx)
		resp, err := l.CreateTls(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
