package collection

import (
	"backend/internal/common/db"
	"backend/internal/types"
)

func Delete(req *types.IDRequest) (*types.Collection, error) {
	var collection types.Collection
	result := db.DB.Preload("Category").Where("path = ?", req.ID).First(&collection)
	if result.Error != nil {
		return nil, result.Error
	}

	db.DB.Delete(&collection)
	return &collection, nil
}
