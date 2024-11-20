package collection

import (
	"errors"
	"net/http"

	"backend/internal/common/db"
	"backend/internal/common/utils"
	"backend/internal/logic/collection"
	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gorm.io/gorm"
)

// 页面集合
func CollectionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AnyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		now := utils.Today()
		ip := utils.GetIPFromRequest(r)

		s := &types.Statistics{
			IP:   ip,
			Date: now,
		}

		// 每天请求时, 追加记录
		if err := db.DB.Model(types.Statistics{}).Where("date =? And ip = ?", ip, now).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			db.DB.Model(types.Statistics{}).Create(s)
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
