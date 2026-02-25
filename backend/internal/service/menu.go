package service

import (
	"github.com/kohwg/gowoopi/backend/internal/model"
	"github.com/kohwg/gowoopi/backend/internal/repository"
)

type menuService struct {
	menuRepo     repository.MenuRepository
	categoryRepo repository.CategoryRepository
}

func NewMenuService(menuRepo repository.MenuRepository, categoryRepo repository.CategoryRepository) MenuService {
	return &menuService{menuRepo: menuRepo, categoryRepo: categoryRepo}
}

func (s *menuService) GetMenusByStore(storeID string) ([]model.Menu, error) {
	return s.menuRepo.FindByStore(storeID)
}

func (s *menuService) GetCategoriesByStore(storeID string) ([]model.Category, error) {
	return s.categoryRepo.FindByStore(storeID)
}

func (s *menuService) CreateMenu(storeID string, req model.MenuCreateRequest) (*model.Menu, error) {
	if _, err := s.categoryRepo.FindByID(req.CategoryID); err != nil {
		return nil, model.WrapNotFound("카테고리")
	}
	isAvailable := true
	if req.IsAvailable != nil {
		isAvailable = *req.IsAvailable
	}
	menu := &model.Menu{
		StoreID: storeID, CategoryID: req.CategoryID, Name: req.Name,
		Description: req.Description, Price: req.Price, ImageURL: req.ImageURL, IsAvailable: isAvailable,
	}
	if err := s.menuRepo.Create(menu); err != nil {
		return nil, model.ErrInternal
	}
	return menu, nil
}

func (s *menuService) UpdateMenu(id uint, req model.MenuUpdateRequest) (*model.Menu, error) {
	menu, err := s.menuRepo.FindByID(id)
	if err != nil {
		return nil, model.WrapNotFound("메뉴")
	}
	if req.CategoryID != nil {
		menu.CategoryID = *req.CategoryID
	}
	if req.Name != nil {
		menu.Name = *req.Name
	}
	if req.Description != nil {
		menu.Description = *req.Description
	}
	if req.Price != nil {
		menu.Price = *req.Price
	}
	if req.ImageURL != nil {
		menu.ImageURL = *req.ImageURL
	}
	if req.IsAvailable != nil {
		menu.IsAvailable = *req.IsAvailable
	}
	if err := s.menuRepo.Update(menu); err != nil {
		return nil, model.ErrInternal
	}
	return menu, nil
}

func (s *menuService) DeleteMenu(id uint) error {
	if _, err := s.menuRepo.FindByID(id); err != nil {
		return model.WrapNotFound("메뉴")
	}
	return s.menuRepo.Delete(id)
}

func (s *menuService) UpdateMenuOrder(items []model.MenuOrderRequest) error {
	inputs := make([]model.MenuOrderInput, len(items))
	for i, item := range items {
		inputs[i] = model.MenuOrderInput(item)
	}
	return s.menuRepo.UpdateOrder(inputs)
}
