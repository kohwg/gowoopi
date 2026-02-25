package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// OrderStatus - 주문 상태 ENUM
type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "PENDING"
	OrderStatusConfirmed OrderStatus = "CONFIRMED"
	OrderStatusPreparing OrderStatus = "PREPARING"
	OrderStatusCompleted OrderStatus = "COMPLETED"
)

// IsValid - 유효한 상태인지 확인
func (s OrderStatus) IsValid() bool {
	switch s {
	case OrderStatusPending, OrderStatusConfirmed, OrderStatusPreparing, OrderStatusCompleted:
		return true
	}
	return false
}

// CanTransitionTo - 허용된 상태 전이인지 확인
func (s OrderStatus) CanTransitionTo(next OrderStatus) bool {
	transitions := map[OrderStatus]OrderStatus{
		OrderStatusPending:   OrderStatusConfirmed,
		OrderStatusConfirmed: OrderStatusPreparing,
		OrderStatusPreparing: OrderStatusCompleted,
	}
	allowed, ok := transitions[s]
	return ok && allowed == next
}

type Order struct {
	ID          string      `gorm:"type:char(36);primaryKey" json:"id"`
	SessionID   string      `gorm:"type:char(36);not null;index" json:"sessionId"`
	StoreID     string      `gorm:"type:char(36);not null;index" json:"storeId"`
	TableID     uint        `gorm:"not null" json:"tableId"`
	Status      OrderStatus `gorm:"type:enum('PENDING','CONFIRMED','PREPARING','COMPLETED');not null;default:'PENDING'" json:"status"`
	TotalAmount uint        `gorm:"not null;default:0" json:"totalAmount"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Session     TableSession `gorm:"foreignKey:SessionID" json:"-"`
	Store       Store        `gorm:"foreignKey:StoreID" json:"-"`
	Table       Table        `gorm:"foreignKey:TableID" json:"-"`
	Items       []OrderItem  `gorm:"foreignKey:OrderID" json:"items"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	if o.ID == "" {
		o.ID = uuid.New().String()
	}
	return nil
}

type OrderItem struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID   string    `gorm:"type:char(36);not null;index" json:"orderId"`
	MenuID    uint      `gorm:"not null" json:"menuId"`
	MenuName  string    `gorm:"type:varchar(100);not null" json:"menuName"`
	Price     uint      `gorm:"not null" json:"price"`
	Quantity  uint      `gorm:"not null" json:"quantity"`
	Subtotal  uint      `gorm:"not null" json:"subtotal"`
	CreatedAt time.Time `json:"createdAt"`
	Order     Order     `gorm:"foreignKey:OrderID" json:"-"`
	Menu      Menu      `gorm:"foreignKey:MenuID" json:"-"`
}

// Validate - 주문 항목 유효성 검증
func (i *OrderItem) Validate() error {
	if i.Quantity == 0 {
		return fmt.Errorf("quantity must be at least 1")
	}
	i.Subtotal = i.Price * i.Quantity
	return nil
}
