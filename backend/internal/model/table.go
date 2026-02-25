package model

type Table struct {
	ID           uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	StoreID      string `gorm:"type:char(36);not null;index" json:"storeId"`
	TableNumber  int    `gorm:"not null" json:"tableNumber"`
	PasswordHash string `gorm:"type:varchar(255);not null" json:"-"`
	IsActive     bool   `gorm:"not null;default:true" json:"isActive"`
	BaseModel
	Store Store `gorm:"foreignKey:StoreID" json:"-"`
}
