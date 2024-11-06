package tag

import (
	"context"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// tags页面集合
func NewTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TagsLogic {
	return &TagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TagsLogic) Tags(req *types.AnyRequest) (resp *types.TagsResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
