package auth

import (
	"net/http"

	"backend/internal/logic/auth"
	"backend/internal/svc"
	"backend/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取验证信息
func AuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.User
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewAuthLogic(r.Context(), svcCtx)
		resp, err := l.Auth(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
