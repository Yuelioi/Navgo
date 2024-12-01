package collection

import (
	"context"

	"backend/internal/common/dao/category"
	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NavsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 页面集合
func NewNavsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NavsLogic {
	return &NavsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NavsLogic) Navs() (resp *types.NavsResponse, err error) {
	navs, err := category.Navs()
	if err != nil {
		return nil, err
	}

	return &types.NavsResponse{
		Navs: navs,
	}, err
}
