package collection

import (
	"backend/internal/common/db"
	"backend/internal/types"
)

func Update(req *types.CollectionUpdateParams) (*types.Collection, error) {
	var collection types.Collection
	result := db.DB.Preload("Category").Where("cid = ?", req.CID).First(&collection)

	if result.Error != nil {
		return nil, result.Error
	}

	collection.Title = req.Title
	collection.Link = req.Link
	collection.Description = req.Description
	collection.Tags = req.Tags
	db.DB.Save(&collection)
	return &collection, nil
}
