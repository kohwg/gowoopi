package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Store struct {
	ID              string `gorm:"type:char(36);primaryKey" json:"id"`
	Name            string `gorm:"type:varchar(100);not null" json:"name"`
	DefaultLanguage string `gorm:"type:varchar(5);not null;default:'ko'" json:"defaultLanguage"`
	BaseModel
}

func (s *Store) BeforeCreate(tx *gorm.DB) error {
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return nil
}
