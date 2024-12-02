package collection

import (
	"backend/internal/common/db"
	"backend/internal/types"
	"path/filepath"
)

func Add(req *types.CollectionUpdateParams) (*types.Collection, error) {
	var cat types.Category

	collection := types.Collection{
		CID:         req.CID,
		Title:       req.Title,
		Description: req.Description,
		Link:        req.Link,
		Tags:        req.Tags,
	}

	// 设置路径
	collection.Path = filepath.Join(req.CategoryPath, req.CID)

	result := db.DB.Model(&types.Category{}).Where("path = ?", req.CategoryPath).Take(&cat)
	if result.Error != nil {
		return nil, result.Error
	}

	collection.Category = &cat

	result = db.DB.Create(&collection)
	if result.Error != nil {
		return nil, result.Error
	}

	return &collection, nil
}
