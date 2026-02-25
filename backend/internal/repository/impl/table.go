package impl

import (
	"github.com/kohwg/gowoopi/backend/internal/model"
	"gorm.io/gorm"
)

type tableRepository struct {
	db *gorm.DB
}

func NewTableRepository(db *gorm.DB) *tableRepository {
	return &tableRepository{db: db}
}

func (r *tableRepository) FindByStoreAndNumber(storeID string, number int) (*model.Table, error) {
	var table model.Table
	if err := r.db.First(&table, "store_id = ? AND table_number = ?", storeID, number).Error; err != nil {
		return nil, err
	}
	return &table, nil
}

func (r *tableRepository) Create(table *model.Table) error {
	return r.db.Create(table).Error
}

func (r *tableRepository) Update(table *model.Table) error {
	return r.db.Save(table).Error
}
