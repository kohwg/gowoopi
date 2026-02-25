package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Store struct {
	ID                string `gorm:"type:char(36);primaryKey"`
	Name              string `gorm:"type:varchar(100);not null"`
	AdminUsername     string `gorm:"type:varchar(50);not null"`
	AdminPasswordHash string `gorm:"type:varchar(255);not null"`
	DefaultLanguage   string `gorm:"type:varchar(5);not null;default:'ko'"`
	BaseModel
}

func (s *Store) BeforeCreate(tx *gorm.DB) error {
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return nil
}
