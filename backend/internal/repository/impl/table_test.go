package impl

import (
	"testing"

	"github.com/google/uuid"
	"github.com/kohwg/gowoopi/backend/internal/model"
	"github.com/kohwg/gowoopi/backend/internal/testutil"
)

func TestTableRepository_Create(t *testing.T) {
	db := testutil.SetupTestDB(t)
	repo := NewTableRepository(db)

	storeID := uuid.New().String()
	db.Create(&model.Store{ID: storeID, Name: "Store", AdminUsername: "admin", AdminPasswordHash: "hash"})
	defer testutil.CleanupTables(t, db, "tables", "stores")

	tbl := &model.Table{StoreID: storeID, TableNumber: 1, PasswordHash: "hash", IsActive: true}
	if err := repo.Create(tbl); err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if tbl.ID == 0 {
		t.Error("Create() should assign ID")
	}
}

func TestTableRepository_FindByStoreAndNumber(t *testing.T) {
	db := testutil.SetupTestDB(t)
	repo := NewTableRepository(db)

	storeID := uuid.New().String()
	db.Create(&model.Store{ID: storeID, Name: "Store", AdminUsername: "admin", AdminPasswordHash: "hash"})
	db.Create(&model.Table{StoreID: storeID, TableNumber: 5, PasswordHash: "hash", IsActive: true})
	defer testutil.CleanupTables(t, db, "tables", "stores")

	found, err := repo.FindByStoreAndNumber(storeID, 5)
	if err != nil {
		t.Fatalf("FindByStoreAndNumber() error = %v", err)
	}
	if found.TableNumber != 5 {
		t.Errorf("FindByStoreAndNumber() number = %d, want 5", found.TableNumber)
	}
}
