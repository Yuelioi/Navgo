package announce

import (
	"net/http"

	"backend/internal/logic/announce"
	"backend/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取通知
func AnnouncesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := announce.NewAnnouncesLogic(r.Context(), svcCtx)
		resp, err := l.Announces()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
