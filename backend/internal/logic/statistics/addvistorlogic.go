package statistics

import (
	"context"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVistorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 增加网站访问量
func NewAddVistorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVistorLogic {
	return &AddVistorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddVistorLogic) AddVistor(req *types.AnyRequest) (resp *types.IDRequest, err error) {
	// todo: add your logic here and delete this line

	return
}
