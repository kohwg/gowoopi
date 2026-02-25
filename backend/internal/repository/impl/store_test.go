package impl

import (
	"testing"

	"github.com/google/uuid"
	"github.com/kohwg/gowoopi/backend/internal/model"
	"github.com/kohwg/gowoopi/backend/internal/testutil"
)

func TestStoreRepository_FindByID(t *testing.T) {
	db := testutil.SetupTestDB(t)
	repo := NewStoreRepository(db)

	store := &model.Store{ID: uuid.New().String(), Name: "Test Store", DefaultLanguage: "ko"}
	db.Create(store)
	defer testutil.CleanupTables(t, db, "stores")

	found, err := repo.FindByID(store.ID)
	if err != nil {
		t.Fatalf("FindByID() error = %v", err)
	}
	if found.Name != "Test Store" {
		t.Errorf("FindByID() name = %q, want %q", found.Name, "Test Store")
	}
}

func TestStoreRepository_FindByID_NotFound(t *testing.T) {
	db := testutil.SetupTestDB(t)
	repo := NewStoreRepository(db)

	_, err := repo.FindByID("nonexistent")
	if err == nil {
		t.Error("FindByID() expected error for nonexistent store")
	}
}
