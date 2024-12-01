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

type CC struct {
	ID      string `json:"id"`
	Favicon string `json:"favicon"`
}

func (l *CollectionLogic) Collection() (resp *types.Collection, err error) {

	return
}
