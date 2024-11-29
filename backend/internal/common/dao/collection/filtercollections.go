package collection

import (
	"backend/internal/common/dao/category"
	"backend/internal/common/db"
	"backend/internal/types"
)

func FilterCollections(req *types.CollectionFilter) ([]*types.Collection, error) {

	catIds, err := category.GetCategoryIDs(req.Categories)

	if err != nil {
		return nil, err
	}

	var collections []*types.Collection

	offset := (req.Page - 1) * req.Limit
	result := db.DB.Preload("Category").Limit(req.Limit).
		Offset(offset).Where("category_id IN ?", catIds).Find(&collections)

	if result.Error != nil {
		return nil, result.Error
	}
	return collections, nil
}
