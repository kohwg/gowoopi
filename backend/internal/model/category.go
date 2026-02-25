package model

type Category struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	StoreID   string `gorm:"type:char(36);not null;index"`
	Name      string `gorm:"type:varchar(50);not null"`
	SortOrder int    `gorm:"not null;default:0"`
	BaseModel

	Store Store `gorm:"foreignKey:StoreID"`
}
