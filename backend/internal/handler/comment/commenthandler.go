package comment

import (
	"net/http"

	"backend/internal/logic/comment"
	"backend/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取页面评论
func CommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := comment.NewCommentLogic(r.Context(), svcCtx)
		resp, err := l.Comment()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
