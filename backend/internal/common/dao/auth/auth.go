package auth

import (
	"backend/internal/common/db"
	"backend/internal/types"
	"errors"

	"gorm.io/gorm"
)

func GetUser(username string) (*types.User, error) {
	var user types.User

	err := db.DB.Model(&types.User{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("未找到用户")
		}
		return nil, err // 返回其他数据库错误
	}

	return &user, nil

}
