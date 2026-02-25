package service

import (
	"time"

	"github.com/kohwg/gowoopi/backend/internal/model"
)

type AuthService interface {
	CustomerLogin(req model.CustomerLoginRequest) (*model.AuthResponse, *model.TokenPair, error)
	AdminLogin(req model.AdminLoginRequest) (*model.AuthResponse, *model.TokenPair, error)
	RefreshToken(refreshToken string) (string, error)
	GenerateTokenPair(claims model.Claims) (*model.TokenPair, error)
	ValidateToken(token string) (*model.Claims, error)
}

type MenuService interface {
	GetMenusByStore(storeID string) ([]model.Menu, error)
	CreateMenu(storeID string, req model.MenuCreateRequest) (*model.Menu, error)
	UpdateMenu(id uint, req model.MenuUpdateRequest) (*model.Menu, error)
	DeleteMenu(id uint) error
	UpdateMenuOrder(items []model.MenuOrderRequest) error
}

type OrderService interface {
	CreateOrder(storeID, sessionID string, tableID uint, req model.OrderCreateRequest) (*model.Order, error)
	GetOrdersBySession(sessionID string) ([]model.Order, error)
	GetOrdersByStore(storeID string) ([]model.Order, error)
	UpdateOrderStatus(id string, req model.StatusUpdateRequest) (*model.Order, error)
	DeleteOrder(id string) error
}

type TableService interface {
	SetupTable(storeID string, req model.TableSetupRequest) (*model.Table, error)
	CompleteTable(tableID uint) error
	GetTableHistory(tableID uint, from, to *time.Time) ([]model.OrderHistory, error)
}

type SSEManager interface {
	Subscribe(storeID string) chan model.SSEEvent
	Unsubscribe(storeID string, ch chan model.SSEEvent)
	Broadcast(storeID string, event model.SSEEvent)
}
