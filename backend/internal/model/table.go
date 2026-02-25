package model

type Table struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	StoreID      string `gorm:"type:char(36);not null;index"`
	TableNumber  int    `gorm:"not null"`
	PasswordHash string `gorm:"type:varchar(255);not null"`
	IsActive     bool   `gorm:"not null;default:true"`
	BaseModel

	Store Store `gorm:"foreignKey:StoreID"`
}
