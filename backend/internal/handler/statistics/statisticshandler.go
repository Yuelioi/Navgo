package statistics

import (
	"net/http"

	"backend/internal/logic/statistics"
	"backend/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取网站数量
func StatisticsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := statistics.NewStatisticsLogic(r.Context(), svcCtx)
		resp, err := l.Statistics()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
