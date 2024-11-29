package collection

import (
	"net/http"
	"strconv"

	"backend/internal/logic/collection"
	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func parseCollectionsFilterForm(r *http.Request) *types.CollectionFilter {
	page := r.FormValue("page")
	limit := r.FormValue("limit")
	categories := r.Form["categories"]

	pageNum := 1
	limitNum := 20

	v, err := strconv.Atoi(page)
	if err == nil && v != 0 {
		pageNum = v
	}

	v, err = strconv.Atoi(limit)
	if err == nil && v != 0 {
		limitNum = v
	}

	var categoriesList []string
	if len(categories) > 0 {
		categoriesList = categories
	} else {
		categoriesList = []string{}
	}

	return &types.CollectionFilter{
		Categories: categoriesList,
		Page:       pageNum,
		Limit:      limitNum,
	}
}

// 分页集合
func FilteredCollectionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := parseCollectionsFilterForm(r)

		l := collection.NewFilteredCollectionsLogic(r.Context(), svcCtx)
		resp, err := l.FilteredCollections(req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
