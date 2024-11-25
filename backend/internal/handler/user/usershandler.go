package user

import (
	"net/http"

	"backend/internal/logic/user"
	"backend/internal/svc"
	"backend/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// tags页面集合
func UsersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AnyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUsersLogic(r.Context(), svcCtx)
		resp, err := l.Users(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
