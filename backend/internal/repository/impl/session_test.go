package impl

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kohwg/gowoopi/backend/internal/model"
	"github.com/kohwg/gowoopi/backend/internal/testutil"
)

func setupSessionTestData(t *testing.T) (*sessionRepository, string, uint, func()) {
	t.Helper()
	db := testutil.SetupTestDB(t)
	repo := NewSessionRepository(db)

	storeID := uuid.New().String()
	db.Create(&model.Store{ID: storeID, Name: "Store", AdminUsername: "admin", AdminPasswordHash: "hash"})
	tbl := &model.Table{StoreID: storeID, TableNumber: 1, PasswordHash: "hash", IsActive: true}
	db.Create(tbl)

	cleanup := func() {
		testutil.CleanupTables(t, db, "table_sessions", "tables", "stores")
	}
	return repo, storeID, tbl.ID, cleanup
}

func TestSessionRepository_Create(t *testing.T) {
	repo, storeID, tableID, cleanup := setupSessionTestData(t)
	defer cleanup()

	session := &model.TableSession{ID: uuid.New().String(), TableID: tableID, StoreID: storeID, StartedAt: time.Now(), IsActive: true}
	if err := repo.Create(session); err != nil {
		t.Fatalf("Create() error = %v", err)
	}
}

func TestSessionRepository_FindActiveByTable(t *testing.T) {
	repo, storeID, tableID, cleanup := setupSessionTestData(t)
	defer cleanup()

	session := &model.TableSession{ID: uuid.New().String(), TableID: tableID, StoreID: storeID, StartedAt: time.Now(), IsActive: true}
	repo.Create(session)

	found, err := repo.FindActiveByTable(tableID)
	if err != nil {
		t.Fatalf("FindActiveByTable() error = %v", err)
	}
	if found.ID != session.ID {
		t.Errorf("FindActiveByTable() id = %q, want %q", found.ID, session.ID)
	}
}

func TestSessionRepository_End(t *testing.T) {
	repo, storeID, tableID, cleanup := setupSessionTestData(t)
	defer cleanup()

	session := &model.TableSession{ID: uuid.New().String(), TableID: tableID, StoreID: storeID, StartedAt: time.Now(), IsActive: true}
	repo.Create(session)

	if err := repo.End(session.ID); err != nil {
		t.Fatalf("End() error = %v", err)
	}

	_, err := repo.FindActiveByTable(tableID)
	if err == nil {
		t.Error("FindActiveByTable() should return error after session ended")
	}
}
