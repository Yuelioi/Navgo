package tag

import (
	"net/http"

	"backend/internal/logic/tag"
	"backend/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// tags页面集合
func TagsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := tag.NewTagsLogic(r.Context(), svcCtx)
		resp, err := l.Tags()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
