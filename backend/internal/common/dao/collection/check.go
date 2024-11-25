package collection

import (
	"backend/internal/common/db"
	"backend/internal/types"
	"errors"

	"gorm.io/gorm"
)

func IsDuplicate(cid string) bool {
	var c types.Collection
	err := db.DB.Model(&types.Collection{}).Where("cid =?", cid).First(&c).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	}
	return false
}
