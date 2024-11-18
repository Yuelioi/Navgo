package types

import (
	"backend/internal/common/utils"

	"gorm.io/gorm"
)

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	now := utils.Today()
	m.CreatedAt = now
	m.UpdatedAt = now
	return
}

func (m *Model) BeforeUpdate(tx *gorm.DB) (err error) {
	now := utils.Today()
	m.UpdatedAt = now
	return
}

func (m *Announce) BeforeCreate(tx *gorm.DB) (err error) {
	now := utils.Today()
	if m.Date == "" {
		m.Date = now
	}
	m.Model.BeforeCreate(tx)
	return
}
func (m *Announce) BeforeUpdate(tx *gorm.DB) (err error) {
	m.Model.BeforeUpdate(tx)
	return
}

func (m *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	now := utils.Today()
	if m.Date == "" {
		m.Date = now
	}
	m.Model.BeforeCreate(tx)
	return
}
func (m *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	m.Model.BeforeUpdate(tx)
	return
}
