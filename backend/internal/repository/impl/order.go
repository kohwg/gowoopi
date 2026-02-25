package impl

import (
	"encoding/json"
	"time"

	"github.com/kohwg/gowoopi/backend/internal/model"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(order *model.Order, items []model.OrderItem) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}
		var total uint
		for i := range items {
			items[i].OrderID = order.ID
			items[i].Subtotal = items[i].Price * items[i].Quantity
			total += items[i].Subtotal
		}
		if err := tx.Create(&items).Error; err != nil {
			return err
		}
		return tx.Model(order).Update("total_amount", total).Error
	})
}

func (r *orderRepository) FindBySession(sessionID string) ([]model.Order, error) {
	var orders []model.Order
	if err := r.db.Preload("Items").Where("session_id = ?", sessionID).Order("created_at DESC").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) FindByStore(storeID string) ([]model.Order, error) {
	var orders []model.Order
	if err := r.db.Preload("Items").Where("store_id = ?", storeID).Order("created_at DESC").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) FindByID(id string) (*model.Order, error) {
	var order model.Order
	if err := r.db.Preload("Items").First(&order, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) UpdateStatus(id string, status model.OrderStatus) error {
	return r.db.Model(&model.Order{}).Where("id = ?", id).Update("status", status).Error
}

func (r *orderRepository) Delete(id string) error {
	return r.db.Delete(&model.Order{}, "id = ?", id).Error
}

func (r *orderRepository) MoveToHistory(sessionID string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var orders []model.Order
		if err := tx.Preload("Items").Where("session_id = ?", sessionID).Find(&orders).Error; err != nil {
			return err
		}

		now := time.Now()
		for _, order := range orders {
			itemsJSON, err := json.Marshal(order.Items)
			if err != nil {
				return err
			}

			var table model.Table
			if err := tx.First(&table, order.TableID).Error; err != nil {
				return err
			}

			history := model.OrderHistory{
				OriginalOrderID: order.ID,
				StoreID:         order.StoreID,
				TableID:         order.TableID,
				TableNumber:     table.TableNumber,
				SessionID:       order.SessionID,
				Status:          string(order.Status),
				TotalAmount:     order.TotalAmount,
				ItemsJSON:       string(itemsJSON),
				OrderedAt:       order.CreatedAt,
				CompletedAt:     now,
			}
			if err := tx.Create(&history).Error; err != nil {
				return err
			}

			if err := tx.Delete(&order).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *orderRepository) FindHistory(tableID uint, from, to *time.Time) ([]model.OrderHistory, error) {
	var histories []model.OrderHistory
	query := r.db.Where("table_id = ?", tableID)
	if from != nil {
		query = query.Where("ordered_at >= ?", *from)
	}
	if to != nil {
		query = query.Where("ordered_at <= ?", *to)
	}
	if err := query.Order("completed_at DESC").Find(&histories).Error; err != nil {
		return nil, err
	}
	return histories, nil
}
