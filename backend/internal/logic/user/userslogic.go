package user

import (
	"context"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// tags页面集合
func NewUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsersLogic {
	return &UsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UsersLogic) Users(req *types.AnyRequest) (resp *types.UserResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
