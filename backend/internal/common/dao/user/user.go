package user

import (
	"backend/internal/common/db"
	"backend/internal/types"
)

func Get(username string) (user *types.User, err error) {
	err = db.DB.Model(&types.User{}).Where("username = ?", username).Take(&user).Error
	return
}
