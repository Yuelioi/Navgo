package collection

import (
	"backend/internal/common/db"
	"backend/internal/types"
	"errors"
)

func All() ([]*types.Collection, error) {
	var collections []*types.Collection
	db.DB.Preload("Category").Find(&collections)

	if collections == nil {
		return nil, errors.New("text string")

	}
	return collections, nil
}
