package model

type Menu struct {
	ID          uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	StoreID     string   `gorm:"type:char(36);not null;index" json:"storeId"`
	CategoryID  uint     `gorm:"not null;index" json:"categoryId"`
	Name        string   `gorm:"type:varchar(100);not null" json:"name"`
	Description string   `gorm:"type:text" json:"description"`
	Price       uint     `gorm:"not null" json:"price"`
	ImageURL    string   `gorm:"type:varchar(500)" json:"imageUrl"`
	IsAvailable bool     `gorm:"not null;default:true" json:"isAvailable"`
	SortOrder   int      `gorm:"not null;default:0" json:"sortOrder"`
	BaseModel
	Store    Store    `gorm:"foreignKey:StoreID" json:"-"`
	Category Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
}

// MenuOrderInput - 메뉴 순서 변경 입력
type MenuOrderInput struct {
	ID        uint
	SortOrder int
}
