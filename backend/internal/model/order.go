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
	ID          string      `gorm:"type:char(36);primaryKey"`
	SessionID   string      `gorm:"type:char(36);not null;index"`
	StoreID     string      `gorm:"type:char(36);not null;index"`
	TableID     uint        `gorm:"not null"`
	Status      OrderStatus `gorm:"type:enum('PENDING','CONFIRMED','PREPARING','COMPLETED');not null;default:'PENDING'"`
	TotalAmount uint        `gorm:"not null;default:0"`
	BaseModel

	Session TableSession `gorm:"foreignKey:SessionID"`
	Store   Store        `gorm:"foreignKey:StoreID"`
	Table   Table        `gorm:"foreignKey:TableID"`
	Items   []OrderItem  `gorm:"foreignKey:OrderID"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	if o.ID == "" {
		o.ID = uuid.New().String()
	}
	return nil
}

type OrderItem struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	OrderID   string    `gorm:"type:char(36);not null;index"`
	MenuID    uint      `gorm:"not null"`
	MenuName  string    `gorm:"type:varchar(100);not null"`
	Price     uint      `gorm:"not null"`
	Quantity  uint      `gorm:"not null"`
	Subtotal  uint      `gorm:"not null"`
	CreatedAt time.Time

	Order Order `gorm:"foreignKey:OrderID"`
	Menu  Menu  `gorm:"foreignKey:MenuID"`
}

// Validate - 주문 항목 유효성 검증
func (i *OrderItem) Validate() error {
	if i.Quantity == 0 {
		return fmt.Errorf("quantity must be at least 1")
	}
	i.Subtotal = i.Price * i.Quantity
	return nil
}
