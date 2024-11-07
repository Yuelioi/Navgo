package types

import "gorm.io/gorm"

type DBComment struct {
	gorm.Model
	Comment
}

type DBCollection struct {
	gorm.Model
	Collection
}
