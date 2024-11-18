package comment

import (
	"context"
	"os"
	"path/filepath"

	"backend/internal/common/constants"
	"backend/internal/common/db"
	"backend/internal/common/utils"
	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/yaml.v3"
)

type AddCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 发布评论
func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCommentLogic) AddComment(req *types.Comment) (resp *types.IDResponse, err error) {

	if req.Date == "" {
		req.Date = utils.Today()
	}

	inReviewMeta := filepath.Join(constants.ConfInst.Resource.Comments, constants.ConfInst.Resource.MetaFile)

	// 2. 读取meta文件
	meta := &types.CommentsResponse{}

	yamlFile, err := os.ReadFile(inReviewMeta)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, meta)
	if err != nil {
		return
	}

	// 添加新的集合
	meta.Comments = append(meta.Comments, req)

	// 重新序列化为 YAML
	data, err := yaml.Marshal(meta)
	if err != nil {
		return
	}

	// 写入文件，覆盖原有内容
	err = os.WriteFile(inReviewMeta, data, os.ModePerm)
	if err != nil {
		return
	}

	db.DB.Model(types.Comment{}).Create(&req)
	return
}
