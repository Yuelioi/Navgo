package collection

import (
	"context"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 单页面
func NewCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectionLogic {
	return &CollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollectionLogic) Collection(req *types.IDRequest) (resp *types.Collection, err error) {
	// todo: add your logic here and delete this line

	return
}
