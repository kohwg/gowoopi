package impl

import (
	"testing"

	"github.com/google/uuid"
	"github.com/kohwg/gowoopi/backend/internal/model"
	"github.com/kohwg/gowoopi/backend/internal/testutil"
)

func TestAdminRepository_CreateAndFind(t *testing.T) {
	db := testutil.SetupTestDB(t)
	repo := NewAdminRepository(db)

	storeID := uuid.New().String()
	db.Create(&model.Store{ID: storeID, Name: "Store"})

	admin := &model.Admin{StoreID: storeID, Username: "testadmin", PasswordHash: "hash", Name: "테스트"}
	if err := repo.Create(admin); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	found, err := repo.FindByStoreAndUsername(storeID, "testadmin")
	if err != nil {
		t.Fatalf("FindByStoreAndUsername() error = %v", err)
	}
	if found.Username != "testadmin" || found.Name != "테스트" {
		t.Errorf("unexpected admin: %+v", found)
	}

	// cleanup
	testutil.CleanupTables(t, db, "admins", "stores")
}

func TestAdminRepository_FindByStoreAndUsername_NotFound(t *testing.T) {
	db := testutil.SetupTestDB(t)
	repo := NewAdminRepository(db)

	_, err := repo.FindByStoreAndUsername("nonexistent", "nobody")
	if err == nil {
		t.Fatal("expected error for non-existent admin")
	}
}
