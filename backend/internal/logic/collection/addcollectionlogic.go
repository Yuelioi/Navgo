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

func (l *AddCollectionLogic) AddCollection(req *types.Collection) (resp *types.Collection, err error) {
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
}
