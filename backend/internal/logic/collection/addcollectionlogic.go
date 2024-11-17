package collection

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"

	"backend/internal/db"
	"backend/internal/global"
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

func initDst(inReviewDir, inReviewMeta string) error {
	if _, err := os.Stat(inReviewDir); os.IsNotExist(err) {
		return os.MkdirAll(inReviewDir, os.ModePerm)
	}

	if _, err := os.Stat(inReviewMeta); os.IsNotExist(err) {

		f, err := os.OpenFile(inReviewMeta, os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.Write([]byte(
			`title: 待审核
cid: InReview
order: 1
collections:
`,
		))
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *AddCollectionLogic) AddCollection(req *types.Collection) (resp *types.Collection, err error) {

	// TODO 临时提交也加入重复计算

	// 检测是否重复
	var c *types.Collection
	if err = db.DB.Model(&types.Collection{}).Where("cid =?", req.CID).First(&c).Error; err == nil {
		return nil, errors.New("重复提交")
	}

	inReviewDir := filepath.Join(global.ConfInst.Resource.Collections, "_待审核")
	inReviewMeta := filepath.Join(inReviewDir, global.ConfInst.Resource.MetaFile)

	err = initDst(inReviewDir, inReviewMeta)
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(inReviewMeta, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	yamlFile, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	meta := &types.Group{}
	yaml.Unmarshal(yamlFile, meta)
	meta.Collections = append(meta.Collections, req)

	data, err := yaml.Marshal(meta)
	if err != nil {
		return nil, err
	}

	// 截断文件并写入新内容
	_, err = f.Seek(0, 0) // 移动文件指针到文件开头
	if err != nil {
		return nil, err
	}
	err = f.Truncate(0) // 清空文件内容
	if err != nil {
		return nil, err
	}
	_, err = f.Write(data)
	if err != nil {
		return nil, err
	}

	return
}
