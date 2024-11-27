package service

import (
	"context"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WallpaperLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWallpaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WallpaperLogic {
	return &WallpaperLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WallpaperLogic) Wallpaper(req *types.AnyRequest) (resp *types.IDResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
