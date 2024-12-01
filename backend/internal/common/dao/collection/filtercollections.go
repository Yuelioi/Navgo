package collection

import (
	"backend/internal/common/dao/category"
	"backend/internal/common/db"
	"backend/internal/types"
	"math"

	"gorm.io/gorm"
)

func filterCats(catIds []uint) func(db *gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		return d.Where("category_id IN ?", catIds)
	}
}

func filterKeyword(keyword string) func(db *gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		return d.Where("title like ?", "%"+keyword+"%").
			Or("description like ?", "%"+keyword+"%").
			Or("link like ?", "%"+keyword+"%")
	}
}

func filterPage(pageSize, offset int) func(db *gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		return d.Limit(pageSize).
			Offset(offset)
	}
}

func FilterCollections(req *types.CollectionFilter) (*types.CollectionsResponse, error) {
	var collections []*types.Collection
	var count int64

	offset := (req.Page - 1) * req.PageSize
	if len(req.Categories) > 0 {
		catIds, err := category.CategoryIDs(req.Categories)

		if err != nil {
			return nil, err
		}

		if req.Keyword != "" {
			result := db.DB.Model(&types.Collection{}).
				Scopes(filterCats(catIds), filterKeyword(req.Keyword)).
				Count(&count)
			if result.Error != nil {
				return nil, result.Error
			}
		} else {
			result := db.DB.Model(&types.Collection{}).
				Scopes(filterCats(catIds)).
				Count(&count)
			if result.Error != nil {
				return nil, result.Error
			}
		}

		if req.Keyword != "" {
			result := db.DB.Preload("Category").
				Scopes(filterCats(catIds), filterPage(req.PageSize, offset), filterKeyword(req.Keyword)).
				Find(&collections)

			if result.Error != nil {
				return nil, result.Error
			}
		} else {
			result := db.DB.Preload("Category").
				Scopes(filterCats(catIds), filterPage(req.PageSize, offset)).
				Find(&collections)

			if result.Error != nil {
				return nil, result.Error
			}
		}

	} else {

		if req.Keyword != "" {
			result := db.DB.Model(&types.Collection{}).Scopes(filterKeyword(req.Keyword)).Count(&count)
			if result.Error != nil {
				return nil, result.Error
			}
		} else {
			result := db.DB.Model(&types.Collection{}).Count(&count)
			if result.Error != nil {
				return nil, result.Error
			}
		}

		if req.Keyword != "" {
			result := db.DB.Preload("Category").Scopes(filterPage(req.PageSize, offset), filterKeyword(req.Keyword)).Find(&collections)

			if result.Error != nil {
				return nil, result.Error
			}
		} else {
			result := db.DB.Preload("Category").Scopes(filterPage(req.PageSize, offset)).Find(&collections)
			if result.Error != nil {
				return nil, result.Error
			}
		}

	}

	// 计算总页数
	totalPages := int(math.Ceil(float64(count) / float64(req.PageSize)))
	return &types.CollectionsResponse{
		Count:       int(count),
		TotalPages:  totalPages,
		Collections: collections,
	}, nil

}
