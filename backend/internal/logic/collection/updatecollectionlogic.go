package collection

import (
	"context"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新页面
func NewUpdateCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCollectionLogic {
	return &UpdateCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCollectionLogic) UpdateCollection(req *types.CollectionUpdateParams) (resp *types.Collection, err error) {

	// 修改数据库

	// 修改本地源

	return
}
