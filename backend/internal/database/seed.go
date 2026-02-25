package database

import (
	"github.com/kohwg/gowoopi/backend/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	storeID := "00000000-0000-0000-0000-000000000001"

	store := model.Store{ID: storeID, Name: "테스트 매장", DefaultLanguage: "ko"}
	if err := db.FirstOrCreate(&store, model.Store{ID: storeID}).Error; err != nil {
		return err
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), 10)
	admin := model.Admin{StoreID: storeID, Username: "admin", PasswordHash: string(hash), Name: "관리자"}
	db.FirstOrCreate(&admin, model.Admin{StoreID: storeID, Username: "admin"})

	categories := []model.Category{
		{StoreID: storeID, Name: "메인", SortOrder: 1},
		{StoreID: storeID, Name: "사이드", SortOrder: 2},
		{StoreID: storeID, Name: "음료", SortOrder: 3},
	}
	for i := range categories {
		db.FirstOrCreate(&categories[i], model.Category{StoreID: storeID, Name: categories[i].Name})
	}

	tblHash, _ := bcrypt.GenerateFromPassword([]byte("1234"), 10)
	table := model.Table{StoreID: storeID, TableNumber: 1, PasswordHash: string(tblHash), IsActive: true}
	db.FirstOrCreate(&table, model.Table{StoreID: storeID, TableNumber: 1})

	menus := []model.Menu{
		{StoreID: storeID, CategoryID: categories[0].ID, Name: "김치찌개", Price: 9000, IsAvailable: true, SortOrder: 1},
		{StoreID: storeID, CategoryID: categories[0].ID, Name: "된장찌개", Price: 8000, IsAvailable: true, SortOrder: 2},
		{StoreID: storeID, CategoryID: categories[1].ID, Name: "계란말이", Price: 5000, IsAvailable: true, SortOrder: 1},
		{StoreID: storeID, CategoryID: categories[2].ID, Name: "콜라", Price: 2000, IsAvailable: true, SortOrder: 1},
	}
	for i := range menus {
		db.FirstOrCreate(&menus[i], model.Menu{StoreID: storeID, Name: menus[i].Name})
	}

	return nil
}
