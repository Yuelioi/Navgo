package auth

import (
	"context"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取验证信息
func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogic {
	return &AuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthLogic) Auth(req *types.User) (resp *types.AuthResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
