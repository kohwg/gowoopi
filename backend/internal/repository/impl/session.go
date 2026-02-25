package impl

import (
	"time"

	"github.com/kohwg/gowoopi/backend/internal/model"
	"gorm.io/gorm"
)

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *sessionRepository {
	return &sessionRepository{db: db}
}

func (r *sessionRepository) Create(session *model.TableSession) error {
	return r.db.Create(session).Error
}

func (r *sessionRepository) FindActiveByTable(tableID uint) (*model.TableSession, error) {
	var session model.TableSession
	if err := r.db.First(&session, "table_id = ? AND is_active = ?", tableID, true).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *sessionRepository) End(sessionID string) error {
	now := time.Now()
	return r.db.Model(&model.TableSession{}).Where("id = ?", sessionID).
		Updates(map[string]interface{}{"is_active": false, "ended_at": now}).Error
}
