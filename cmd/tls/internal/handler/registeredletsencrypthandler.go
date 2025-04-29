package handler

import (
	"net/http"

	"esay-tls-server/cmd/tls/internal/logic"
	"esay-tls-server/cmd/tls/internal/svc"
	"esay-tls-server/cmd/tls/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func registeredLetsEncryptHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostRegisteredLetsEncryptReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRegisteredLetsEncryptLogic(r.Context(), svcCtx)
		resp, err := l.RegisteredLetsEncrypt(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
