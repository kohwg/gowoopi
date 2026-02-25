package impl

import (
	"github.com/gowoopi/backend/internal/model"
	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) FindByStoreAndUsername(storeID, username string) (*model.Admin, error) {
	var admin model.Admin
	if err := r.db.First(&admin, "store_id = ? AND username = ?", storeID, username).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *adminRepository) Create(admin *model.Admin) error {
	return r.db.Create(admin).Error
}
