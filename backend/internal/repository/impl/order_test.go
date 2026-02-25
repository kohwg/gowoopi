package impl

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kohwg/gowoopi/backend/internal/model"
	"github.com/kohwg/gowoopi/backend/internal/testutil"
	"gorm.io/gorm"
)

func setupOrderTestData(t *testing.T) (*orderRepository, *gorm.DB, string, string, uint, uint, func()) {
	t.Helper()
	db := testutil.SetupTestDB(t)
	repo := NewOrderRepository(db)

	storeID := uuid.New().String()
	db.Create(&model.Store{ID: storeID, Name: "Store", AdminUsername: "admin", AdminPasswordHash: "hash"})
	tbl := &model.Table{StoreID: storeID, TableNumber: 1, PasswordHash: "hash", IsActive: true}
	db.Create(tbl)
	sessionID := uuid.New().String()
	db.Create(&model.TableSession{ID: sessionID, TableID: tbl.ID, StoreID: storeID, StartedAt: time.Now(), IsActive: true})
	cat := &model.Category{StoreID: storeID, Name: "Main", SortOrder: 1}
	db.Create(cat)
	menu := &model.Menu{StoreID: storeID, CategoryID: cat.ID, Name: "김치찌개", Price: 9000, IsAvailable: true}
	db.Create(menu)

	cleanup := func() {
		testutil.CleanupTables(t, db, "order_items", "orders", "order_histories", "menus", "categories", "table_sessions", "tables", "stores")
	}
	return repo, db, storeID, sessionID, tbl.ID, menu.ID, cleanup
}

func TestOrderRepository_Create(t *testing.T) {
	repo, _, storeID, sessionID, tableID, menuID, cleanup := setupOrderTestData(t)
	defer cleanup()

	order := &model.Order{ID: uuid.New().String(), SessionID: sessionID, StoreID: storeID, TableID: tableID, Status: model.OrderStatusPending}
	items := []model.OrderItem{
		{MenuID: menuID, MenuName: "김치찌개", Price: 9000, Quantity: 2},
	}

	if err := repo.Create(order, items); err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if order.TotalAmount != 18000 {
		t.Errorf("Create() total = %d, want 18000", order.TotalAmount)
	}
}

func TestOrderRepository_FindBySession(t *testing.T) {
	repo, _, storeID, sessionID, tableID, menuID, cleanup := setupOrderTestData(t)
	defer cleanup()

	order := &model.Order{ID: uuid.New().String(), SessionID: sessionID, StoreID: storeID, TableID: tableID, Status: model.OrderStatusPending}
	repo.Create(order, []model.OrderItem{{MenuID: menuID, MenuName: "김치찌개", Price: 9000, Quantity: 1}})

	orders, err := repo.FindBySession(sessionID)
	if err != nil {
		t.Fatalf("FindBySession() error = %v", err)
	}
	if len(orders) != 1 {
		t.Fatalf("FindBySession() count = %d, want 1", len(orders))
	}
	if len(orders[0].Items) != 1 {
		t.Errorf("FindBySession() items count = %d, want 1 (Preload)", len(orders[0].Items))
	}
}

func TestOrderRepository_FindByStore(t *testing.T) {
	repo, _, storeID, sessionID, tableID, menuID, cleanup := setupOrderTestData(t)
	defer cleanup()

	order := &model.Order{ID: uuid.New().String(), SessionID: sessionID, StoreID: storeID, TableID: tableID, Status: model.OrderStatusPending}
	repo.Create(order, []model.OrderItem{{MenuID: menuID, MenuName: "김치찌개", Price: 9000, Quantity: 1}})

	orders, err := repo.FindByStore(storeID)
	if err != nil {
		t.Fatalf("FindByStore() error = %v", err)
	}
	if len(orders) != 1 {
		t.Fatalf("FindByStore() count = %d, want 1", len(orders))
	}
}

func TestOrderRepository_UpdateStatus(t *testing.T) {
	repo, _, storeID, sessionID, tableID, menuID, cleanup := setupOrderTestData(t)
	defer cleanup()

	order := &model.Order{ID: uuid.New().String(), SessionID: sessionID, StoreID: storeID, TableID: tableID, Status: model.OrderStatusPending}
	repo.Create(order, []model.OrderItem{{MenuID: menuID, MenuName: "김치찌개", Price: 9000, Quantity: 1}})

	if err := repo.UpdateStatus(order.ID, model.OrderStatusConfirmed); err != nil {
		t.Fatalf("UpdateStatus() error = %v", err)
	}

	found, _ := repo.FindByID(order.ID)
	if found.Status != model.OrderStatusConfirmed {
		t.Errorf("UpdateStatus() status = %q, want %q", found.Status, model.OrderStatusConfirmed)
	}
}

func TestOrderRepository_Delete_SoftDelete(t *testing.T) {
	repo, _, storeID, sessionID, tableID, menuID, cleanup := setupOrderTestData(t)
	defer cleanup()

	order := &model.Order{ID: uuid.New().String(), SessionID: sessionID, StoreID: storeID, TableID: tableID, Status: model.OrderStatusPending}
	repo.Create(order, []model.OrderItem{{MenuID: menuID, MenuName: "김치찌개", Price: 9000, Quantity: 1}})

	if err := repo.Delete(order.ID); err != nil {
		t.Fatalf("Delete() error = %v", err)
	}

	orders, _ := repo.FindByStore(storeID)
	if len(orders) != 0 {
		t.Errorf("FindByStore() after delete count = %d, want 0", len(orders))
	}
}

func TestOrderRepository_MoveToHistory(t *testing.T) {
	repo, _, storeID, sessionID, tableID, menuID, cleanup := setupOrderTestData(t)
	defer cleanup()

	order := &model.Order{ID: uuid.New().String(), SessionID: sessionID, StoreID: storeID, TableID: tableID, Status: model.OrderStatusCompleted}
	repo.Create(order, []model.OrderItem{{MenuID: menuID, MenuName: "김치찌개", Price: 9000, Quantity: 1}})

	if err := repo.MoveToHistory(sessionID); err != nil {
		t.Fatalf("MoveToHistory() error = %v", err)
	}

	// 원본 주문 Soft Delete 확인
	orders, _ := repo.FindBySession(sessionID)
	if len(orders) != 0 {
		t.Errorf("FindBySession() after move count = %d, want 0", len(orders))
	}

	// 이력 생성 확인
	histories, err := repo.FindHistory(tableID, nil, nil)
	if err != nil {
		t.Fatalf("FindHistory() error = %v", err)
	}
	if len(histories) != 1 {
		t.Fatalf("FindHistory() count = %d, want 1", len(histories))
	}
	if histories[0].OriginalOrderID != order.ID {
		t.Errorf("FindHistory() original_order_id = %q, want %q", histories[0].OriginalOrderID, order.ID)
	}
}

func TestOrderRepository_FindHistory_WithDateFilter(t *testing.T) {
	repo, _, storeID, sessionID, tableID, menuID, cleanup := setupOrderTestData(t)
	defer cleanup()

	order := &model.Order{ID: uuid.New().String(), SessionID: sessionID, StoreID: storeID, TableID: tableID, Status: model.OrderStatusCompleted}
	repo.Create(order, []model.OrderItem{{MenuID: menuID, MenuName: "김치찌개", Price: 9000, Quantity: 1}})
	repo.MoveToHistory(sessionID)

	// 미래 날짜 필터 - 결과 없어야 함
	future := time.Now().Add(24 * time.Hour)
	histories, _ := repo.FindHistory(tableID, &future, nil)
	if len(histories) != 0 {
		t.Errorf("FindHistory() with future filter count = %d, want 0", len(histories))
	}
}
