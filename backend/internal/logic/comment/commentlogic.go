package comment

import (
	"context"

	"backend/internal/db"
	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取页面评论
func NewCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentLogic {
	return &CommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentLogic) Comment(req *types.AnyRequest) (resp *types.CommentsResponse, err error) {

	var datas []*types.Comment
	if err = db.DB.Model(types.Comment{}).Order("date DESC").Find(&datas).Error; err == nil {
		resp = &types.CommentsResponse{
			Comments: datas,
		}
	}
	return
}
