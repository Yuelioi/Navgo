package collection

import (
	"net/http"

	"backend/internal/common/dao/statistic"
	"backend/internal/common/utils"
	"backend/internal/logic/collection"
	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 页面集合
func CollectionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AnyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		ip := utils.GetIPFromRequest(r)

		//  添加每日访问(限定ip以及今日)
		err := statistic.AddCustomer(ip)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := collection.NewCollectionsLogic(r.Context(), svcCtx)
		resp, err := l.Collections(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
