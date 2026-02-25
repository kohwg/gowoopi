package impl

import (
	"testing"

	"github.com/google/uuid"
	"github.com/kohwg/gowoopi/backend/internal/model"
	"github.com/kohwg/gowoopi/backend/internal/testutil"
)

func TestCategoryRepository_Create(t *testing.T) {
	db := testutil.SetupTestDB(t)
	repo := NewCategoryRepository(db)

	storeID := uuid.New().String()
	db.Create(&model.Store{ID: storeID, Name: "Store"})
	defer testutil.CleanupTables(t, db, "categories", "stores")

	cat := &model.Category{StoreID: storeID, Name: "Main", SortOrder: 1}
	if err := repo.Create(cat); err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if cat.ID == 0 {
		t.Error("Create() should assign ID")
	}
}

func TestCategoryRepository_FindByStore(t *testing.T) {
	db := testutil.SetupTestDB(t)
	repo := NewCategoryRepository(db)

	storeID := uuid.New().String()
	db.Create(&model.Store{ID: storeID, Name: "Store"})
	db.Create(&model.Category{StoreID: storeID, Name: "B", SortOrder: 2})
	db.Create(&model.Category{StoreID: storeID, Name: "A", SortOrder: 1})
	defer testutil.CleanupTables(t, db, "categories", "stores")

	cats, err := repo.FindByStore(storeID)
	if err != nil {
		t.Fatalf("FindByStore() error = %v", err)
	}
	if len(cats) != 2 {
		t.Fatalf("FindByStore() count = %d, want 2", len(cats))
	}
	if cats[0].Name != "A" {
		t.Errorf("FindByStore() first = %q, want %q (sorted by sort_order)", cats[0].Name, "A")
	}
}

func TestCategoryRepository_Delete_SoftDelete(t *testing.T) {
	db := testutil.SetupTestDB(t)
	repo := NewCategoryRepository(db)

	storeID := uuid.New().String()
	db.Create(&model.Store{ID: storeID, Name: "Store"})
	cat := &model.Category{StoreID: storeID, Name: "ToDelete", SortOrder: 1}
	db.Create(cat)
	defer testutil.CleanupTables(t, db, "categories", "stores")

	if err := repo.Delete(cat.ID); err != nil {
		t.Fatalf("Delete() error = %v", err)
	}

	// 일반 조회에서 제외
	cats, _ := repo.FindByStore(storeID)
	if len(cats) != 0 {
		t.Errorf("FindByStore() after delete count = %d, want 0", len(cats))
	}

	// Unscoped로 확인
	var count int64
	db.Unscoped().Model(&model.Category{}).Where("id = ?", cat.ID).Count(&count)
	if count != 1 {
		t.Error("Soft deleted record should still exist in DB")
	}
}
