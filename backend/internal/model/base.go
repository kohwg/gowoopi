package model

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel - GORM 공통 필드 (Soft Delete 대상 엔티티용)
type BaseModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
