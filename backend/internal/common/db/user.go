package db

import (
	"backend/internal/common/biz"
	"backend/internal/common/constants"
	"backend/internal/types"

	"gorm.io/gorm/clause"
)

type UserManager struct {
	user *types.User
}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (m *UserManager) Read() error {
	hashedPwd, err := biz.HashPassword(constants.ConfInst.System.Password)
	if err != nil {
		return err
	}

	m.user = &types.User{
		Username: constants.ConfInst.System.Username,
		Password: hashedPwd,
		Role:     "admin",
	}
	return nil
}

func (m *UserManager) Build() error {
	DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "username"}},
		UpdateAll: true,
	}).Create(m.user)
	return nil
}

func (m *UserManager) Reduce() error {
	return nil
}
