package announce

import (
	"context"

	"backend/internal/common/db"
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

func (l *AnnouncesLogic) Announces(req *types.AnyRequest) (resp *types.AnnouncesData, err error) {

	var datas []*types.Announce
	if err = db.DB.Model(types.Announce{}).Order("date DESC").Find(&datas).Error; err == nil {
		resp = &types.AnnouncesData{
			Announces: datas,
		}
	}
	return
}
