package collection

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"backend/internal/common/cache"
	"backend/internal/common/constants"
	"backend/internal/common/db"
	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/yaml.v3"
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

	// TODO 临时提交也加入重复计算

	// 检测是否重复
	var exists bool
	db.DB.Model(&types.Collection{}).Where("cid =?", req.CID).Scan(&exists)
	if exists {
		return nil, errors.New("重复提交")
	}

	inReviewDir := filepath.Join(constants.ConfInst.Resource.Collections, "_待审核")
	inReviewMeta := filepath.Join(inReviewDir, constants.ConfInst.Resource.MetaFile)

	yamlFile, err := os.ReadFile(inReviewMeta)
	if err != nil {
		return nil, err
	}

	meta := &types.Group{}
	err = yaml.Unmarshal(yamlFile, meta)
	if err != nil {
		return
	}

	// 添加新的集合
	meta.Collections = append(meta.Collections, req)

	// 重新序列化为 YAML
	data, err := yaml.Marshal(meta)
	if err != nil {
		return
	}

	// 写入文件，覆盖原有内容
	err = os.WriteFile(inReviewMeta, data, os.ModePerm)

	// 缓存增加该条记录
	cache.Cache.Add(req.CID, req)

	return
}
