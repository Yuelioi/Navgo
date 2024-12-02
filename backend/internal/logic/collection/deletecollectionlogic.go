package collection

import (
	"context"

	"backend/internal/common/cache"
	"backend/internal/common/dao/collection"
	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除页面
func NewDeleteCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCollectionLogic {
	return &DeleteCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCollectionLogic) DeleteCollection(req *types.IDRequest) (resp *types.Collection, err error) {
	// 修改数据库
	c, err := collection.Delete(req)
	if err != nil {
		return nil, err
	}

	// 修改本地源
	cc := cache.Manager.GetController(cache.LocalCacheID)
	cc.Remove(c.Category.Path, c)

	return
}
