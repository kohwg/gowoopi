package impl

import (
	"testing"

	"github.com/google/uuid"
	"github.com/kohwg/gowoopi/backend/internal/model"
	"github.com/kohwg/gowoopi/backend/internal/testutil"
	"gorm.io/gorm"
)

func setupMenuTestData(t *testing.T) (*menuRepository, *gorm.DB, string, uint, func()) {
	t.Helper()
	db := testutil.SetupTestDB(t)
	repo := NewMenuRepository(db)

	storeID := uuid.New().String()
	db.Create(&model.Store{ID: storeID, Name: "Store", AdminUsername: "admin", AdminPasswordHash: "hash"})
	cat := &model.Category{StoreID: storeID, Name: "Main", SortOrder: 1}
	db.Create(cat)

	cleanup := func() {
		testutil.CleanupTables(t, db, "menus", "categories", "stores")
	}
	return repo, db, storeID, cat.ID, cleanup
}

func TestMenuRepository_Create(t *testing.T) {
	repo, _, storeID, catID, cleanup := setupMenuTestData(t)
	defer cleanup()

	menu := &model.Menu{StoreID: storeID, CategoryID: catID, Name: "김치찌개", Price: 9000, IsAvailable: true}
	if err := repo.Create(menu); err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if menu.ID == 0 {
		t.Error("Create() should assign ID")
	}
}

func TestMenuRepository_FindByStore(t *testing.T) {
	repo, db, storeID, catID, cleanup := setupMenuTestData(t)
	defer cleanup()

	db.Create(&model.Menu{StoreID: storeID, CategoryID: catID, Name: "B", Price: 8000, SortOrder: 2, IsAvailable: true})
	db.Create(&model.Menu{StoreID: storeID, CategoryID: catID, Name: "A", Price: 9000, SortOrder: 1, IsAvailable: true})

	menus, err := repo.FindByStore(storeID)
	if err != nil {
		t.Fatalf("FindByStore() error = %v", err)
	}
	if len(menus) != 2 {
		t.Fatalf("FindByStore() count = %d, want 2", len(menus))
	}
	if menus[0].Name != "A" {
		t.Errorf("FindByStore() first = %q, want %q (sorted)", menus[0].Name, "A")
	}
}

func TestMenuRepository_Delete_SoftDelete(t *testing.T) {
	repo, _, storeID, catID, cleanup := setupMenuTestData(t)
	defer cleanup()

	menu := &model.Menu{StoreID: storeID, CategoryID: catID, Name: "ToDelete", Price: 5000, IsAvailable: true}
	_ = repo.Create(menu)

	if err := repo.Delete(menu.ID); err != nil {
		t.Fatalf("Delete() error = %v", err)
	}

	menus, _ := repo.FindByStore(storeID)
	if len(menus) != 0 {
		t.Errorf("FindByStore() after delete count = %d, want 0", len(menus))
	}
}

func TestMenuRepository_UpdateOrder(t *testing.T) {
	repo, _, storeID, catID, cleanup := setupMenuTestData(t)
	defer cleanup()

	m1 := &model.Menu{StoreID: storeID, CategoryID: catID, Name: "A", Price: 1000, SortOrder: 1, IsAvailable: true}
	m2 := &model.Menu{StoreID: storeID, CategoryID: catID, Name: "B", Price: 2000, SortOrder: 2, IsAvailable: true}
	_ = repo.Create(m1)
	_ = repo.Create(m2)

	err := repo.UpdateOrder([]model.MenuOrderInput{
		{ID: m1.ID, SortOrder: 2},
		{ID: m2.ID, SortOrder: 1},
	})
	if err != nil {
		t.Fatalf("UpdateOrder() error = %v", err)
	}

	menus, _ := repo.FindByStore(storeID)
	if menus[0].Name != "B" {
		t.Errorf("UpdateOrder() first after reorder = %q, want %q", menus[0].Name, "B")
	}
}
