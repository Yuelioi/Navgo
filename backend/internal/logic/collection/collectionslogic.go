package collection

import (
	"context"
	"strings"

	"backend/internal/db"
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

func (l *CollectionsLogic) Collections(req *types.AnyRequest) (resp *types.CollectionsResponse, err error) {
	resp = &types.CollectionsResponse{
		Datas: []*types.CollectionsData{},
	}

	var collections []*types.Collection
	db.DB.Preload("Category").Model(&types.Collection{}).Find(&collections)

	for _, collection := range collections {

		var topData *types.CollectionsData
		var group *types.Group

		// 顶级分类
		topID := strings.Split(collection.Category.Path, "\\")[0]

		// 查找顶级数据
		for _, top := range resp.Datas {
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
			resp.Datas = append(resp.Datas, topData)
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

	return
}
