package service

import (
	"github.com/google/uuid"
	"github.com/gowoopi/backend/internal/model"
	"github.com/gowoopi/backend/internal/repository"
)

type orderService struct {
	orderRepo repository.OrderRepository
	menuRepo  repository.MenuRepository
	sse       SSEManager
}

func NewOrderService(orderRepo repository.OrderRepository, menuRepo repository.MenuRepository, sse SSEManager) OrderService {
	return &orderService{orderRepo: orderRepo, menuRepo: menuRepo, sse: sse}
}

func (s *orderService) CreateOrder(storeID, sessionID string, tableID uint, req model.OrderCreateRequest) (*model.Order, error) {
	var items []model.OrderItem
	for _, ri := range req.Items {
		menu, err := s.menuRepo.FindByID(ri.MenuID)
		if err != nil {
			return nil, model.WrapNotFound("메뉴")
		}
		if !menu.IsAvailable {
			return nil, model.NewAppError("VALIDATION_ERROR", menu.Name+"은(는) 현재 판매 중지 상태입니다", 400)
		}
		items = append(items, model.OrderItem{MenuID: menu.ID, MenuName: menu.Name, Price: menu.Price, Quantity: ri.Quantity})
	}

	order := &model.Order{ID: uuid.New().String(), SessionID: sessionID, StoreID: storeID, TableID: tableID, Status: model.OrderStatusPending}
	if err := s.orderRepo.Create(order, items); err != nil {
		return nil, model.ErrInternal
	}

	found, _ := s.orderRepo.FindByID(order.ID)
	s.sse.Broadcast(storeID, model.SSEEvent{Type: model.SSEOrderCreated, Data: found})
	return found, nil
}

func (s *orderService) GetOrdersBySession(sessionID string) ([]model.Order, error) {
	return s.orderRepo.FindBySession(sessionID)
}

func (s *orderService) GetOrdersByStore(storeID string) ([]model.Order, error) {
	return s.orderRepo.FindByStore(storeID)
}

func (s *orderService) UpdateOrderStatus(id string, req model.StatusUpdateRequest) (*model.Order, error) {
	order, err := s.orderRepo.FindByID(id)
	if err != nil {
		return nil, model.WrapNotFound("주문")
	}
	newStatus := model.OrderStatus(req.Status)
	if !newStatus.IsValid() {
		return nil, model.ErrValidation
	}
	if !order.Status.CanTransitionTo(newStatus) {
		return nil, model.ErrInvalidStatusTransition
	}
	if err := s.orderRepo.UpdateStatus(id, newStatus); err != nil {
		return nil, model.ErrInternal
	}

	order.Status = newStatus
	s.sse.Broadcast(order.StoreID, model.SSEEvent{Type: model.SSEOrderStatusChanged, Data: order})
	return order, nil
}

func (s *orderService) DeleteOrder(id string) error {
	order, err := s.orderRepo.FindByID(id)
	if err != nil {
		return model.WrapNotFound("주문")
	}
	if order.Status == model.OrderStatusCompleted {
		return model.NewAppError("VALIDATION_ERROR", "완료된 주문은 삭제할 수 없습니다", 400)
	}
	if err := s.orderRepo.Delete(id); err != nil {
		return model.ErrInternal
	}
	s.sse.Broadcast(order.StoreID, model.SSEEvent{Type: model.SSEOrderDeleted, Data: map[string]string{"id": id}})
	return nil
}
