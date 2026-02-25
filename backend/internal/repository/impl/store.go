package impl

import (
	"github.com/kohwg/gowoopi/backend/internal/model"
	"gorm.io/gorm"
)

type storeRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) *storeRepository {
	return &storeRepository{db: db}
}

func (r *storeRepository) FindByID(id string) (*model.Store, error) {
	var store model.Store
	if err := r.db.First(&store, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &store, nil
}
