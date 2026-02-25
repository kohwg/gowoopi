package service

import (
	"testing"

	"github.com/gowoopi/backend/internal/mock"
	"github.com/gowoopi/backend/internal/model"
	"go.uber.org/mock/gomock"
)

func newOrderTestService(t *testing.T) (*gomock.Controller, *mock.MockOrderRepository, *mock.MockMenuRepository, *mock.MockSSEManager, OrderService) {
	ctrl := gomock.NewController(t)
	orderRepo := mock.NewMockOrderRepository(ctrl)
	menuRepo := mock.NewMockMenuRepository(ctrl)
	sseMgr := mock.NewMockSSEManager(ctrl)
	svc := NewOrderService(orderRepo, menuRepo, sseMgr)
	return ctrl, orderRepo, menuRepo, sseMgr, svc
}

func TestOrderService_CreateOrder_Success(t *testing.T) {
	ctrl, orderRepo, menuRepo, sseMgr, svc := newOrderTestService(t)
	defer ctrl.Finish()

	menu := &model.Menu{ID: 1, Name: "김치찌개", Price: 9000, IsAvailable: true}
	menuRepo.EXPECT().FindByID(uint(1)).Return(menu, nil)
	orderRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
	orderRepo.EXPECT().FindByID(gomock.Any()).Return(&model.Order{ID: "order1", StoreID: "store1", Status: model.OrderStatusPending}, nil)
	sseMgr.EXPECT().Broadcast("store1", gomock.Any())

	order, err := svc.CreateOrder("store1", "sess1", 1, model.OrderCreateRequest{
		Items: []model.OrderItemRequest{{MenuID: 1, Quantity: 2}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if order.ID != "order1" {
		t.Errorf("order ID = %q, want order1", order.ID)
	}
}

func TestOrderService_CreateOrder_MenuNotFound(t *testing.T) {
	ctrl, _, menuRepo, _, svc := newOrderTestService(t)
	defer ctrl.Finish()

	menuRepo.EXPECT().FindByID(uint(99)).Return(nil, model.ErrNotFound)

	_, err := svc.CreateOrder("store1", "sess1", 1, model.OrderCreateRequest{
		Items: []model.OrderItemRequest{{MenuID: 99, Quantity: 1}},
	})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestOrderService_CreateOrder_MenuUnavailable(t *testing.T) {
	ctrl, _, menuRepo, _, svc := newOrderTestService(t)
	defer ctrl.Finish()

	menu := &model.Menu{ID: 1, Name: "김치찌개", IsAvailable: false}
	menuRepo.EXPECT().FindByID(uint(1)).Return(menu, nil)

	_, err := svc.CreateOrder("store1", "sess1", 1, model.OrderCreateRequest{
		Items: []model.OrderItemRequest{{MenuID: 1, Quantity: 1}},
	})
	if err == nil {
		t.Fatal("expected error for unavailable menu")
	}
}

func TestOrderService_GetOrdersBySession(t *testing.T) {
	ctrl, orderRepo, _, _, svc := newOrderTestService(t)
	defer ctrl.Finish()

	expected := []model.Order{{ID: "o1", SessionID: "sess1"}}
	orderRepo.EXPECT().FindBySession("sess1").Return(expected, nil)

	orders, err := svc.GetOrdersBySession("sess1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(orders) != 1 {
		t.Errorf("count = %d, want 1", len(orders))
	}
}

func TestOrderService_GetOrdersByStore(t *testing.T) {
	ctrl, orderRepo, _, _, svc := newOrderTestService(t)
	defer ctrl.Finish()

	expected := []model.Order{{ID: "o1", StoreID: "store1"}}
	orderRepo.EXPECT().FindByStore("store1").Return(expected, nil)

	orders, err := svc.GetOrdersByStore("store1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(orders) != 1 {
		t.Errorf("count = %d, want 1", len(orders))
	}
}

func TestOrderService_UpdateOrderStatus_Success(t *testing.T) {
	ctrl, orderRepo, _, sseMgr, svc := newOrderTestService(t)
	defer ctrl.Finish()

	order := &model.Order{ID: "o1", StoreID: "store1", Status: model.OrderStatusPending}
	orderRepo.EXPECT().FindByID("o1").Return(order, nil)
	orderRepo.EXPECT().UpdateStatus("o1", model.OrderStatusConfirmed).Return(nil)
	sseMgr.EXPECT().Broadcast("store1", gomock.Any())

	updated, err := svc.UpdateOrderStatus("o1", model.StatusUpdateRequest{Status: "CONFIRMED"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if updated.Status != model.OrderStatusConfirmed {
		t.Errorf("status = %q, want CONFIRMED", updated.Status)
	}
}

func TestOrderService_UpdateOrderStatus_InvalidTransition(t *testing.T) {
	ctrl, orderRepo, _, _, svc := newOrderTestService(t)
	defer ctrl.Finish()

	order := &model.Order{ID: "o1", Status: model.OrderStatusPending}
	orderRepo.EXPECT().FindByID("o1").Return(order, nil)

	_, err := svc.UpdateOrderStatus("o1", model.StatusUpdateRequest{Status: "COMPLETED"})
	if err != model.ErrInvalidStatusTransition {
		t.Errorf("err = %v, want ErrInvalidStatusTransition", err)
	}
}

func TestOrderService_UpdateOrderStatus_InvalidStatus(t *testing.T) {
	ctrl, orderRepo, _, _, svc := newOrderTestService(t)
	defer ctrl.Finish()

	order := &model.Order{ID: "o1", Status: model.OrderStatusPending}
	orderRepo.EXPECT().FindByID("o1").Return(order, nil)

	_, err := svc.UpdateOrderStatus("o1", model.StatusUpdateRequest{Status: "INVALID"})
	if err != model.ErrValidation {
		t.Errorf("err = %v, want ErrValidation", err)
	}
}

func TestOrderService_DeleteOrder_Success(t *testing.T) {
	ctrl, orderRepo, _, sseMgr, svc := newOrderTestService(t)
	defer ctrl.Finish()

	order := &model.Order{ID: "o1", StoreID: "store1", Status: model.OrderStatusPending}
	orderRepo.EXPECT().FindByID("o1").Return(order, nil)
	orderRepo.EXPECT().Delete("o1").Return(nil)
	sseMgr.EXPECT().Broadcast("store1", gomock.Any())

	if err := svc.DeleteOrder("o1"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestOrderService_DeleteOrder_CompletedOrder(t *testing.T) {
	ctrl, orderRepo, _, _, svc := newOrderTestService(t)
	defer ctrl.Finish()

	order := &model.Order{ID: "o1", Status: model.OrderStatusCompleted}
	orderRepo.EXPECT().FindByID("o1").Return(order, nil)

	err := svc.DeleteOrder("o1")
	if err == nil {
		t.Fatal("expected error for completed order")
	}
}

func TestOrderService_DeleteOrder_NotFound(t *testing.T) {
	ctrl, orderRepo, _, _, svc := newOrderTestService(t)
	defer ctrl.Finish()

	orderRepo.EXPECT().FindByID("o1").Return(nil, model.ErrNotFound)

	err := svc.DeleteOrder("o1")
	if err == nil {
		t.Fatal("expected error")
	}
}
