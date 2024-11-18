package category

import (
	"backend/internal/common/db"
	"backend/internal/types"

	"gorm.io/gorm/clause"
)

func FilterTop() []types.Category {
	var cats []types.Category
	db.DB.Model(&types.Category{}).
		Where("path NOT LIKE ?", "%\\%").
		Order(clause.OrderByColumn{
			Column: clause.Column{Name: "order"}, Desc: true}).
		Find(&cats)
	return cats
}
