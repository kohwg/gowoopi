package model

import "time"

type OrderHistory struct {
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	OriginalOrderID string    `gorm:"type:char(36);not null;index"`
	StoreID         string    `gorm:"type:char(36);not null;index"`
	TableID         uint      `gorm:"not null"`
	TableNumber     int       `gorm:"not null"`
	SessionID       string    `gorm:"type:char(36);not null"`
	Status          string    `gorm:"type:varchar(20);not null"`
	TotalAmount     uint      `gorm:"not null"`
	ItemsJSON       string    `gorm:"type:json;not null"`
	OrderedAt       time.Time `gorm:"not null"`
	CompletedAt     time.Time `gorm:"not null"`
	CreatedAt       time.Time

	Store Store `gorm:"foreignKey:StoreID"`
	Table Table `gorm:"foreignKey:TableID"`
}
