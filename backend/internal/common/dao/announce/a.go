package announce

import (
	"backend/internal/common/db"
	"backend/internal/types"
	"errors"
)

func All() (*types.AnnouncesData, error) {
	var datas []*types.Announce
	if err := db.DB.Model(types.Announce{}).Order("date DESC").Find(&datas).Error; err == nil {
		return &types.AnnouncesData{
			Announces: datas,
		}, nil
	}

	return nil, errors.New("未找到数据")
}
