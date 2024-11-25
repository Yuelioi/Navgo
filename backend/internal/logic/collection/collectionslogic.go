package collection

import (
	"context"
	"path/filepath"
	"sort"
	"strings"

	"backend/internal/common/dao/collection"
	"backend/internal/common/db"
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

func (l *CollectionsLogic) Collections(req *types.AnyRequest) (resp *types.CollectionsResponse, err error) {

	datas := make([]*types.CollectionsData, 0)

	collections, err := collection.All()
	if err != nil {
		return
	}

	for _, collection := range collections {

		var topData *types.CollectionsData
		var group *types.Group

		// 顶级分类
		topID := strings.Split(collection.Category.Path, string(filepath.Separator))[0]

		// 查找顶级数据
		for _, top := range datas {
			if top.Category.CID == topID {
				topData = top
			}
		}
		// 没找到从数据库填充
		// 初始化父分类
		if topData == nil {
			cat := &types.Category{}
			db.DB.Model(&types.Category{}).First(cat, db.DB.Where("path =?", topID))
			topData = &types.CollectionsData{
				Category: cat,
				Groups:   []*types.Group{},
			}
			// 附着到resp
			datas = append(datas, topData)
		}

		// 查找次级数据
		for _, secondary := range topData.Groups {
			if secondary.Category.Path == collection.Category.Path {
				group = secondary
			}
		}
		// 初始化子分类
		if group == nil {
			group = &types.Group{
				Category:    collection.Category,
				Collections: []*types.Collection{},
			}
			// 附着到父类
			topData.Groups = append(topData.Groups, group)
		}

		group.Collections = append(group.Collections, collection)
	}

	resp = &types.CollectionsResponse{
		Datas: sortCollections(datas),
	}

	return
}
