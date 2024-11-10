package types

import (
	"time"

	"gorm.io/gorm"
)

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().Format(time.RFC3339)
	m.CreatedAt = now
	m.UpdatedAt = now
	return
}

func (m *Model) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now().Format(time.RFC3339)
	m.UpdatedAt = now
	return
}
