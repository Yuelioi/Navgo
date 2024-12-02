package collection

import (
	"context"
	"errors"

	"backend/internal/common/cache"
	"backend/internal/common/dao/collection"
	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 增加页面
func NewAddCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCollectionLogic {
	return &AddCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCollectionLogic) AddCollection(req *types.CollectionUpdateParams, isAdmin bool) (resp *types.Collection, err error) {
	if !isAdmin {
		// 检测数据库是否重复
		isDuplicate := collection.IsDuplicate(req.CID)
		if isDuplicate {
			return nil, errors.New("重复提交")
		}

		// 添加成功 添加到缓存区
		controller := cache.Manager.GetController(cache.InReviewCacheID)
		if controller != nil {
			err = controller.Add(req.CID, req)
		}

		return
	} else {
		// 添加数据库
		c, err := collection.Add(req)
		if err != nil {
			return nil, err
		}

		// 修改本地源
		cc := cache.Manager.GetController(cache.LocalCacheID)
		err = cc.Add(c.Category.Path, c)
		if err != nil {
			return nil, err
		}
	}

	return
}
