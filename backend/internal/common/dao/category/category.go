package category

import (
	"backend/internal/common/db"
	"backend/internal/types"
	"strings"

	"gorm.io/gorm/clause"
)

func CategoryIDs(categories []string) ([]uint, error) {
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

func Navs() ([]*types.Nav, error) {

	var cats []*types.Category

	result := db.DB.Model(&types.Category{}).Order(clause.OrderBy{Columns: []clause.OrderByColumn{
		{Column: clause.Column{Name: "depth"}, Desc: false},
		{Column: clause.Column{Name: "order"}, Desc: false},
	}}).Find(&cats)
	if result.Error != nil {
		return nil, result.Error
	}

	parent := make([]*types.Nav, 0)
	pathToNav := make(map[string]*types.Nav)
	for _, cat := range cats {
		if cat.Depth == 1 {
			newNav := &types.Nav{
				CID:       cat.CID,
				Title:     cat.Title,
				FullTitle: cat.FullTitle,
				Order:     cat.Order,
				Path:      cat.Path,
				Children:  []types.Nav{},
			}
			parent = append(parent, newNav)
			pathToNav[cat.Path] = newNav
		} else {

			parentPath := strings.Split(cat.Path, "\\")[0]
			if parentNav, exists := pathToNav[parentPath]; exists {
				childNav := types.Nav{
					CID:       cat.CID,
					Title:     cat.Title,
					FullTitle: cat.FullTitle,
					Order:     cat.Order,
					Path:      cat.Path,
				}
				parentNav.Children = append(parentNav.Children, childNav)
			}
		}
	}

	return parent, nil
}
