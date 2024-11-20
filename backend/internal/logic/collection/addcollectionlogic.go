package collection

import (
	"context"
	"errors"

	"backend/internal/common/cache"
	"backend/internal/common/db"
	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type AddCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 增加页面
func NewAddCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCollectionLogic {
	return &AddCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCollectionLogic) AddCollection(req *types.Collection) (resp *types.Collection, err error) {
	// 检测数据库是否重复
	err = db.DB.Model(&types.Collection{}).Where("cid =?", req.CID).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("重复提交")
	}

	controller := cache.Manager.GetController(cache.InReviewCacheID)
	err = controller.Add("", req)

	return
}
