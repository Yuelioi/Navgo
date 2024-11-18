package statistics

import (
	"context"

	"backend/internal/common/db"
	"backend/internal/common/utils"
	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StatisticsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取网站数量
func NewStatisticsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatisticsLogic {
	return &StatisticsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatisticsLogic) Statistics(req *types.AnyRequest) (resp *types.SiteStats, err error) {

	var yesterday int64
	var all int64
	var links int64

	db.DB.Model(types.Statistics{}).Where("date = ?", utils.Yesterday()).Count(&yesterday)
	db.DB.Model(types.Statistics{}).Count(&all)
	db.DB.Model(types.Collection{}).Count(&links)

	resp = &types.SiteStats{
		LastDayVisitors: int(yesterday),
		TotalVisitors:   int(all),
		LinksCount:      int(links),
	}
	return
}
