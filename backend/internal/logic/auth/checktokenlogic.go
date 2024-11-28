package auth

import (
	"context"

	"backend/internal/common/biz"
	"backend/internal/common/dao/user"
	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 验证token信息
func NewCheckTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckTokenLogic {
	return &CheckTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckTokenLogic) CheckToken(req *types.IDRequest) (resp *types.AuthResponse, err error) {
	username, err := biz.Validate(req.ID)
	if err != nil {
		return nil, err
	}

	u, err := user.Get(username)

	resp = &types.AuthResponse{
		Username: u.Username,
		Email:    u.Email,
		Nickname: u.Nickname,
		Role:     u.Role,
		Token:    req.ID,
	}

	return
}
