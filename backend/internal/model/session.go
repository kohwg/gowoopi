package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TableSession struct {
	ID        string     `gorm:"type:char(36);primaryKey" json:"id"`
	TableID   uint       `gorm:"not null;index" json:"tableId"`
	StoreID   string     `gorm:"type:char(36);not null;index" json:"storeId"`
	StartedAt time.Time  `gorm:"not null" json:"startedAt"`
	EndedAt   *time.Time `json:"endedAt,omitempty"`
	IsActive  bool       `gorm:"not null;default:true" json:"isActive"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	Table     Table      `gorm:"foreignKey:TableID" json:"-"`
	Store     Store      `gorm:"foreignKey:StoreID" json:"-"`
}

func (s *TableSession) BeforeCreate(tx *gorm.DB) error {
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return nil
}
