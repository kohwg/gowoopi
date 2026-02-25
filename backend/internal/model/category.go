package model

type Category struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	StoreID   string `gorm:"type:char(36);not null;index" json:"storeId"`
	Name      string `gorm:"type:varchar(50);not null" json:"name"`
	SortOrder int    `gorm:"not null;default:0" json:"sortOrder"`
	BaseModel
	Store Store `gorm:"foreignKey:StoreID" json:"-"`
}
