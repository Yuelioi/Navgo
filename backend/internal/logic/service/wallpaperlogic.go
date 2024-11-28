package service

import (
	"context"
	"math/rand/v2"
	"os"
	"path/filepath"

	"backend/internal/common/constants"
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

	root := constants.ConfInst.Resource.Wallpaper
	wallpapers := make([]string, 0)

	filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		wallpapers = append(wallpapers, d.Name())
		return nil
	})

	id := rand.IntN(len(wallpapers))

	resp = &types.IDResponse{
		ID: constants.ConfInst.Api.Wallpaper + wallpapers[id],
	}

	return
}
