package collection

import (
	"net/http"

	"backend/internal/logic/collection"
	"backend/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 单页面
func CollectionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := collection.NewCollectionLogic(r.Context(), svcCtx)
		resp, err := l.Collection()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
