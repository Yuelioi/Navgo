package collection

import (
	"errors"
	"net/http"

	global "backend/internal/common/constants"
	"backend/internal/logic/collection"
	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 更新页面
func UpdateCollectionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CollectionUpdateParams

		// 获取表单字段
		err := httpx.Parse(r, &req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, errors.New("无法解析字段"))
			return
		}

		// 保存文件
		file, _, err := r.FormFile("favicon")
		if err == nil {
			defer file.Close()
			_, err := saveFile(w, r, file, global.ConfInst.Resource.Icons, req.CID)
			if err != nil {
				return
			}
			// req.Favicon = icon
		}

		l := collection.NewUpdateCollectionLogic(r.Context(), svcCtx)
		resp, err := l.UpdateCollection(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
