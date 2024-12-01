package user

import (
	"net/http"

	"backend/internal/logic/user"
	"backend/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// tags页面集合
func UsersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := user.NewUsersLogic(r.Context(), svcCtx)
		resp, err := l.Users()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
