package db

import (
	"backend/internal/common/constants"
	"backend/internal/types"

	"gorm.io/gorm/clause"
)

type UserManager struct {
}

func NewUserManager(root string) *UserManager {
	return &UserManager{}
}

func (m *UserManager) Read() error {
	return nil
}

func (m *UserManager) Build() error {
	user := types.User{
		Username: constants.SVCInst.Config.System.Username,
		Password: constants.SVCInst.Config.System.Password,
	}

	DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "username"}},
		UpdateAll: true,
	}).Create(user)
	return nil
}

func (m *UserManager) Reduce() error {
	return nil
}
