package service

import (
	"time"

	"github.com/kohwg/gowoopi/backend/internal/model"
	"github.com/kohwg/gowoopi/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type tableService struct {
	tableRepo   repository.TableRepository
	sessionRepo repository.SessionRepository
	orderRepo   repository.OrderRepository
	sse         SSEManager
}

func NewTableService(tableRepo repository.TableRepository, sessionRepo repository.SessionRepository, orderRepo repository.OrderRepository, sse SSEManager) TableService {
	return &tableService{tableRepo: tableRepo, sessionRepo: sessionRepo, orderRepo: orderRepo, sse: sse}
}

func (s *tableService) SetupTable(storeID string, req model.TableSetupRequest) (*model.Table, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return nil, model.ErrInternal
	}
	tbl := &model.Table{StoreID: storeID, TableNumber: req.TableNumber, PasswordHash: string(hash), IsActive: true}
	if err := s.tableRepo.Create(tbl); err != nil {
		return nil, model.ErrInternal
	}
	return tbl, nil
}

func (s *tableService) CompleteTable(tableID uint) error {
	session, err := s.sessionRepo.FindActiveByTable(tableID)
	if err != nil {
		return model.WrapNotFound("활성 세션")
	}
	if err := s.orderRepo.MoveToHistory(session.ID); err != nil {
		return model.ErrInternal
	}
	if err := s.sessionRepo.End(session.ID); err != nil {
		return model.ErrInternal
	}
	s.sse.Broadcast(session.StoreID, model.SSEEvent{Type: model.SSETableReset, Data: map[string]uint{"table_id": tableID}})
	return nil
}

func (s *tableService) GetTableHistory(tableID uint, from, to *time.Time) ([]model.OrderHistory, error) {
	return s.orderRepo.FindHistory(tableID, from, to)
}
