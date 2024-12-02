package collection

import (
	"context"
	"path/filepath"
	"sort"
	"strings"

	"backend/internal/common/dao/category"
	"backend/internal/common/dao/collection"
	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 页面集合
func NewCollectionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectionsLogic {
	return &CollectionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func sortCollections(datas []*types.CollectionsData) []*types.CollectionsData {
	sort.Slice(datas, func(i, j int) bool {
		return datas[i].Category.Order < datas[j].Category.Order
	})

	for i, data := range datas {
		sort.Slice(data.Groups, func(i, j int) bool {
			return data.Groups[i].Category.Order < data.Groups[j].Category.Order
		})

		datas[i] = data
	}
	return datas
}

func (l *CollectionsLogic) Collections() (resp *types.CollectionsDataResponse, err error) {

	datas := make([]*types.CollectionsData, 0)

	collections, err := collection.All()
	if err != nil {
		return
	}

	topCats, err := category.TopCategories()
	if err != nil {
		return
	}

	// 构建顶层数据
	topCatsMap := make(map[string]*types.CollectionsData, 0)
	for _, cat := range topCats {
		top := &types.CollectionsData{
			Category: cat,
			Groups:   []*types.Group{},
		}
		datas = append(datas, top)
		topCatsMap[cat.CID] = top
	}

	// 构建父级索引
	parentMap := make(map[string]*types.Group, 0)

	for _, collection := range collections {
		// 查找顶级数据
		topID := strings.Split(collection.Category.Path, string(filepath.Separator))[0]

		group, ok := topCatsMap[topID]
		if !ok {
			// 没找到父级分类就继续, 虽然肯定能找到
			continue
		}

		// 查找父级数据
		parent, ok := parentMap[collection.Category.Path]
		if !ok {
			parent = &types.Group{
				Category:    collection.Category,
				Collections: make([]*types.Collection, 0),
			}
			parentMap[collection.Category.Path] = parent
			group.Groups = append(group.Groups, parent)
		}
		parent.Collections = append(parent.Collections, collection)

	}

	resp = &types.CollectionsDataResponse{
		Datas: sortCollections(datas),
	}

	return
}
