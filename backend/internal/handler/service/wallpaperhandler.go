package service

import (
	"net/http"

	"backend/internal/logic/service"
	"backend/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func WallpaperHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := service.NewWallpaperLogic(r.Context(), svcCtx)
		resp, err := l.Wallpaper()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
