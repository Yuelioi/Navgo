package collection

import (
	"context"

	"backend/internal/common/dao/collection"
	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilteredCollectionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页集合
func NewFilteredCollectionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilteredCollectionsLogic {
	return &FilteredCollectionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilteredCollectionsLogic) FilteredCollections(req *types.CollectionFilter) (resp *types.CollectionsResponse, err error) {
	resp, err = collection.FilterCollections(req)
	if err != nil {
		return nil, err
	}

	return
}
