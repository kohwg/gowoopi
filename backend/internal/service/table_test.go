package service

import (
	"testing"

	"github.com/kohwg/gowoopi/backend/internal/mock"
	"github.com/kohwg/gowoopi/backend/internal/model"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func newTableTestService(t *testing.T) (*gomock.Controller, *mock.MockTableRepository, *mock.MockSessionRepository, *mock.MockOrderRepository, *mock.MockSSEManager, TableService) {
	ctrl := gomock.NewController(t)
	tableRepo := mock.NewMockTableRepository(ctrl)
	sessionRepo := mock.NewMockSessionRepository(ctrl)
	orderRepo := mock.NewMockOrderRepository(ctrl)
	sseMgr := mock.NewMockSSEManager(ctrl)
	svc := NewTableService(tableRepo, sessionRepo, orderRepo, sseMgr)
	return ctrl, tableRepo, sessionRepo, orderRepo, sseMgr, svc
}

func TestTableService_SetupTable_Success(t *testing.T) {
	ctrl, tableRepo, _, _, _, svc := newTableTestService(t)
	defer ctrl.Finish()

	tableRepo.EXPECT().Create(gomock.Any()).DoAndReturn(func(tbl *model.Table) error {
		if tbl.StoreID != "store1" || tbl.TableNumber != 5 {
			t.Errorf("unexpected table: %+v", tbl)
		}
		if err := bcrypt.CompareHashAndPassword([]byte(tbl.PasswordHash), []byte("pass")); err != nil {
			t.Errorf("password hash mismatch")
		}
		return nil
	})

	tbl, err := svc.SetupTable("store1", model.TableSetupRequest{TableNumber: 5, Password: "pass"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !tbl.IsActive {
		t.Error("table should be active")
	}
}

func TestTableService_CompleteTable_Success(t *testing.T) {
	ctrl, _, sessionRepo, orderRepo, sseMgr, svc := newTableTestService(t)
	defer ctrl.Finish()

	session := &model.TableSession{ID: "sess1", TableID: 1, StoreID: "store1", IsActive: true}
	sessionRepo.EXPECT().FindActiveByTable(uint(1)).Return(session, nil)
	orderRepo.EXPECT().MoveToHistory("sess1").Return(nil)
	sessionRepo.EXPECT().End("sess1").Return(nil)
	sseMgr.EXPECT().Broadcast("store1", gomock.Any())

	if err := svc.CompleteTable(1); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestTableService_CompleteTable_NoActiveSession(t *testing.T) {
	ctrl, _, sessionRepo, _, _, svc := newTableTestService(t)
	defer ctrl.Finish()

	sessionRepo.EXPECT().FindActiveByTable(uint(1)).Return(nil, model.ErrNotFound)

	err := svc.CompleteTable(1)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestTableService_GetTableHistory(t *testing.T) {
	ctrl, _, _, orderRepo, _, svc := newTableTestService(t)
	defer ctrl.Finish()

	expected := []model.OrderHistory{{ID: 1, TableID: 1, OriginalOrderID: "o1"}}
	orderRepo.EXPECT().FindHistory(uint(1), nil, nil).Return(expected, nil)

	histories, err := svc.GetTableHistory(1, nil, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(histories) != 1 {
		t.Errorf("count = %d, want 1", len(histories))
	}
}
