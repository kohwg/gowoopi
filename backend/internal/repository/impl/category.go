package impl

import (
	"github.com/gowoopi/backend/internal/model"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) FindByStore(storeID string) ([]model.Category, error) {
	var categories []model.Category
	if err := r.db.Where("store_id = ?", storeID).Order("sort_order ASC").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) FindByID(id uint) (*model.Category, error) {
	var category model.Category
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) Create(category *model.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) Update(category *model.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id uint) error {
	return r.db.Delete(&model.Category{}, id).Error
}
