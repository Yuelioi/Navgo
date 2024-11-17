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

func (m *Announce) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().Format("2006-01-02")
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
	now := time.Now().Format("2006-01-02")
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
