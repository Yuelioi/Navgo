package collection

import (
	"context"

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
		Collections: []types.Collection{},
	}

	var collections []types.Collection
	db.DB.Preload("Category").Model(&types.Collection{}).Find(&collections)

	if collections != nil {
		resp.Collections = collections
	}
	return
}
