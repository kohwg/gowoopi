package repository

import (
	"time"

	"github.com/kohwg/gowoopi/backend/internal/model"
)

type StoreRepository interface {
	FindByID(id string) (*model.Store, error)
}

type CategoryRepository interface {
	FindByStore(storeID string) ([]model.Category, error)
	FindByID(id uint) (*model.Category, error)
	Create(category *model.Category) error
	Update(category *model.Category) error
	Delete(id uint) error
}

type TableRepository interface {
	FindByStoreAndNumber(storeID string, number int) (*model.Table, error)
	Create(table *model.Table) error
	Update(table *model.Table) error
}

type SessionRepository interface {
	Create(session *model.TableSession) error
	FindActiveByTable(tableID uint) (*model.TableSession, error)
	End(sessionID string) error
}

type MenuRepository interface {
	FindByStore(storeID string) ([]model.Menu, error)
	FindByID(id uint) (*model.Menu, error)
	Create(menu *model.Menu) error
	Update(menu *model.Menu) error
	Delete(id uint) error
	UpdateOrder(items []model.MenuOrderInput) error
}

type OrderRepository interface {
	Create(order *model.Order, items []model.OrderItem) error
	FindBySession(sessionID string) ([]model.Order, error)
	FindByStore(storeID string) ([]model.Order, error)
	FindByID(id string) (*model.Order, error)
	UpdateStatus(id string, status model.OrderStatus) error
	Delete(id string) error
	MoveToHistory(sessionID string) error
	FindHistory(tableID uint, from, to *time.Time) ([]model.OrderHistory, error)
}

type AdminRepository interface {
	FindByStoreAndUsername(storeID, username string) (*model.Admin, error)
	Create(admin *model.Admin) error
}
