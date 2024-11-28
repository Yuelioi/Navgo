package auth

import (
	"errors"
	"net/http"

	"backend/internal/logic/auth"
	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func parseForm(r *http.Request) (*types.User, error) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" {
		return nil, errors.New("请填写用户名")
	}
	if username == "" {
		return nil, errors.New("请填写密码")
	}

	return &types.User{
		Username: username,
		Password: password,
	}, nil
}

// 获取验证信息
func AuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req *types.User

		req, err := parseForm(r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}

		l := auth.NewAuthLogic(r.Context(), svcCtx)
		resp, err := l.Auth(req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
