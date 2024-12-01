package announce

import (
	"context"

	"backend/internal/common/dao/announce"
	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnnouncesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取通知
func NewAnnouncesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnnouncesLogic {
	return &AnnouncesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AnnouncesLogic) Announces() (resp *types.AnnouncesData, err error) {
	return announce.All()
}
