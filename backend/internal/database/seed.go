package database

import (
	"github.com/google/uuid"
	"github.com/kohwg/gowoopi/backend/internal/model"
	"gorm.io/gorm"
)

// Seed - 개발용 시드 데이터 생성
func Seed(db *gorm.DB) error {
	storeID := uuid.New().String()

	store := model.Store{
		ID:                storeID,
		Name:              "테스트 매장",
		AdminUsername:     "admin",
		AdminPasswordHash: "$2a$10$placeholder", // 실제 사용 시 bcrypt 해시로 교체
		DefaultLanguage:   "ko",
	}
	if err := db.FirstOrCreate(&store, model.Store{ID: storeID}).Error; err != nil {
		return err
	}

	categories := []model.Category{
		{StoreID: storeID, Name: "메인", SortOrder: 1},
		{StoreID: storeID, Name: "사이드", SortOrder: 2},
		{StoreID: storeID, Name: "음료", SortOrder: 3},
	}
	for i := range categories {
		if err := db.FirstOrCreate(&categories[i], model.Category{StoreID: storeID, Name: categories[i].Name}).Error; err != nil {
			return err
		}
	}

	table := model.Table{
		StoreID:      storeID,
		TableNumber:  1,
		PasswordHash: "$2a$10$placeholder",
		IsActive:     true,
	}
	if err := db.FirstOrCreate(&table, model.Table{StoreID: storeID, TableNumber: 1}).Error; err != nil {
		return err
	}

	menus := []model.Menu{
		{StoreID: storeID, CategoryID: categories[0].ID, Name: "김치찌개", Price: 9000, IsAvailable: true, SortOrder: 1},
		{StoreID: storeID, CategoryID: categories[0].ID, Name: "된장찌개", Price: 8000, IsAvailable: true, SortOrder: 2},
		{StoreID: storeID, CategoryID: categories[1].ID, Name: "계란말이", Price: 5000, IsAvailable: true, SortOrder: 1},
		{StoreID: storeID, CategoryID: categories[2].ID, Name: "콜라", Price: 2000, IsAvailable: true, SortOrder: 1},
	}
	for i := range menus {
		if err := db.FirstOrCreate(&menus[i], model.Menu{StoreID: storeID, Name: menus[i].Name}).Error; err != nil {
			return err
		}
	}

	return nil
}
