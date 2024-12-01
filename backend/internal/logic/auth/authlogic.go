package auth

import (
	"context"
	"errors"

	"backend/internal/common/biz"
	"backend/internal/common/dao/auth"
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
	user, err := auth.GetUser(req.Username)
	if err != nil {

		return nil, err
		// // 数据库没有数据, 生成临时token
		// token, err := biz.Generate(user.Username, "guests")
		// if err != nil {
		// 	return nil, errors.New("token 生成失败")
		// }
		// return &types.AuthResponse{
		// 	Username: "访客A",
		// 	Nickname: "访客A",
		// 	Role:     "guests",
		// 	Token:    token,
		// }, err
	}

	if biz.CheckPasswordHash(req.Password, user.Password) {
		// todo 角色设置
		token, err := biz.Generate(user.Username, "admin")
		if err != nil {
			return nil, errors.New("token 生成失败")
		}
		return &types.AuthResponse{
			Username: user.Username,
			Email:    user.Email,
			Nickname: user.Nickname,
			Role:     user.Role,
			Token:    token,
		}, err
	}

	return
}
