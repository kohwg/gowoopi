package model

type Admin struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	StoreID      string `gorm:"type:char(36);not null;index"`
	Username     string `gorm:"type:varchar(50);not null"`
	PasswordHash string `gorm:"type:varchar(255);not null"`
	Name         string `gorm:"type:varchar(50);not null"`
	BaseModel

	Store Store `gorm:"foreignKey:StoreID"`
}
