package statistic

import (
	"backend/internal/common/db"
	"backend/internal/common/utils"
	"backend/internal/types"
	"errors"

	"gorm.io/gorm"
)

// 增加访客记录
func AddCustomer(ip string) error {
	now := utils.Today()

	s := &types.Statistics{
		IP:   ip,
		Date: now,
	}

	var stats types.Statistics
	err := db.DB.Model(types.Statistics{}).Where("date =? And ip = ?", ip, now).First(&stats).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		db.DB.Model(types.Statistics{}).Create(s)
	}

	return nil
}
