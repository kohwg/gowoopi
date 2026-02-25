package impl

import (
	"github.com/gowoopi/backend/internal/model"
	"gorm.io/gorm"
)

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *menuRepository {
	return &menuRepository{db: db}
}

func (r *menuRepository) FindByStore(storeID string) ([]model.Menu, error) {
	var menus []model.Menu
	if err := r.db.Preload("Category").Where("store_id = ?", storeID).Order("sort_order ASC").Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *menuRepository) FindByID(id uint) (*model.Menu, error) {
	var menu model.Menu
	if err := r.db.First(&menu, id).Error; err != nil {
		return nil, err
	}
	return &menu, nil
}

func (r *menuRepository) Create(menu *model.Menu) error {
	return r.db.Create(menu).Error
}

func (r *menuRepository) Update(menu *model.Menu) error {
	return r.db.Save(menu).Error
}

func (r *menuRepository) Delete(id uint) error {
	return r.db.Delete(&model.Menu{}, id).Error
}

func (r *menuRepository) UpdateOrder(items []model.MenuOrderInput) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range items {
			if err := tx.Model(&model.Menu{}).Where("id = ?", item.ID).Update("sort_order", item.SortOrder).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
