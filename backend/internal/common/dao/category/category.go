package category

import (
	"backend/internal/common/db"
	"backend/internal/types"
)

func GetCategoryIDs(categories []string) ([]uint, error) {
	var categoryIDs []uint

	for _, categoryName := range categories {
		var category types.Category
		result := db.DB.Where("path = ?", categoryName).First(&category)
		if result.Error != nil {
			return nil, result.Error
		}
		categoryIDs = append(categoryIDs, category.ID)
	}

	return categoryIDs, nil
}
