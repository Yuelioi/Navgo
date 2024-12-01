package collection

import (
	"net/http"

	"backend/internal/logic/collection"
	"backend/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 页面集合
func NavsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := collection.NewNavsLogic(r.Context(), svcCtx)
		resp, err := l.Navs()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
