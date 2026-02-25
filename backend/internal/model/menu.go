package model

type Menu struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	StoreID     string `gorm:"type:char(36);not null;index"`
	CategoryID  uint   `gorm:"not null;index"`
	Name        string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:text"`
	Price       uint   `gorm:"not null"`
	ImageURL    string `gorm:"type:varchar(500)"`
	IsAvailable bool   `gorm:"not null;default:true"`
	SortOrder   int    `gorm:"not null;default:0"`
	BaseModel

	Store    Store    `gorm:"foreignKey:StoreID"`
	Category Category `gorm:"foreignKey:CategoryID"`
}

// MenuOrderInput - 메뉴 순서 변경 입력
type MenuOrderInput struct {
	ID        uint
	SortOrder int
}
