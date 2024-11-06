package collection

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
