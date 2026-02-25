package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TableSession struct {
	ID        string    `gorm:"type:char(36);primaryKey"`
	TableID   uint      `gorm:"not null;index"`
	StoreID   string    `gorm:"type:char(36);not null;index"`
	StartedAt time.Time `gorm:"not null"`
	EndedAt   *time.Time
	IsActive  bool      `gorm:"not null;default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Table Table `gorm:"foreignKey:TableID"`
	Store Store `gorm:"foreignKey:StoreID"`
}

func (s *TableSession) BeforeCreate(tx *gorm.DB) error {
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return nil
}
