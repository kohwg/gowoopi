package service

import (
	"testing"

	"github.com/gowoopi/backend/internal/mock"
	"github.com/gowoopi/backend/internal/model"
	"go.uber.org/mock/gomock"
)

func newMenuTestService(t *testing.T) (*gomock.Controller, *mock.MockMenuRepository, *mock.MockCategoryRepository, MenuService) {
	ctrl := gomock.NewController(t)
	menuRepo := mock.NewMockMenuRepository(ctrl)
	catRepo := mock.NewMockCategoryRepository(ctrl)
	svc := NewMenuService(menuRepo, catRepo)
	return ctrl, menuRepo, catRepo, svc
}

func TestMenuService_GetMenusByStore(t *testing.T) {
	ctrl, menuRepo, _, svc := newMenuTestService(t)
	defer ctrl.Finish()

	expected := []model.Menu{{ID: 1, Name: "김치찌개", Price: 9000}}
	menuRepo.EXPECT().FindByStore("store1").Return(expected, nil)

	menus, err := svc.GetMenusByStore("store1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(menus) != 1 || menus[0].Name != "김치찌개" {
		t.Errorf("unexpected menus: %+v", menus)
	}
}

func TestMenuService_CreateMenu_Success(t *testing.T) {
	ctrl, menuRepo, catRepo, svc := newMenuTestService(t)
	defer ctrl.Finish()

	catRepo.EXPECT().FindByID(uint(1)).Return(&model.Category{ID: 1}, nil)
	menuRepo.EXPECT().Create(gomock.Any()).Return(nil)

	menu, err := svc.CreateMenu("store1", model.MenuCreateRequest{CategoryID: 1, Name: "된장찌개", Price: 8000})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if menu.Name != "된장찌개" || menu.Price != 8000 {
		t.Errorf("unexpected menu: %+v", menu)
	}
}

func TestMenuService_CreateMenu_CategoryNotFound(t *testing.T) {
	ctrl, _, catRepo, svc := newMenuTestService(t)
	defer ctrl.Finish()

	catRepo.EXPECT().FindByID(uint(99)).Return(nil, model.ErrNotFound)

	_, err := svc.CreateMenu("store1", model.MenuCreateRequest{CategoryID: 99, Name: "test", Price: 1000})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestMenuService_UpdateMenu_Success(t *testing.T) {
	ctrl, menuRepo, _, svc := newMenuTestService(t)
	defer ctrl.Finish()

	existing := &model.Menu{ID: 1, Name: "김치찌개", Price: 9000}
	menuRepo.EXPECT().FindByID(uint(1)).Return(existing, nil)
	menuRepo.EXPECT().Update(gomock.Any()).Return(nil)

	newName := "김치찌개 특"
	menu, err := svc.UpdateMenu(1, model.MenuUpdateRequest{Name: &newName})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if menu.Name != "김치찌개 특" {
		t.Errorf("name = %q, want 김치찌개 특", menu.Name)
	}
}

func TestMenuService_UpdateMenu_NotFound(t *testing.T) {
	ctrl, menuRepo, _, svc := newMenuTestService(t)
	defer ctrl.Finish()

	menuRepo.EXPECT().FindByID(uint(99)).Return(nil, model.ErrNotFound)

	_, err := svc.UpdateMenu(99, model.MenuUpdateRequest{})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestMenuService_DeleteMenu_Success(t *testing.T) {
	ctrl, menuRepo, _, svc := newMenuTestService(t)
	defer ctrl.Finish()

	menuRepo.EXPECT().FindByID(uint(1)).Return(&model.Menu{ID: 1}, nil)
	menuRepo.EXPECT().Delete(uint(1)).Return(nil)

	if err := svc.DeleteMenu(1); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestMenuService_DeleteMenu_NotFound(t *testing.T) {
	ctrl, menuRepo, _, svc := newMenuTestService(t)
	defer ctrl.Finish()

	menuRepo.EXPECT().FindByID(uint(99)).Return(nil, model.ErrNotFound)

	err := svc.DeleteMenu(99)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestMenuService_UpdateMenuOrder(t *testing.T) {
	ctrl, menuRepo, _, svc := newMenuTestService(t)
	defer ctrl.Finish()

	menuRepo.EXPECT().UpdateOrder([]model.MenuOrderInput{
		{ID: 1, SortOrder: 2},
		{ID: 2, SortOrder: 1},
	}).Return(nil)

	err := svc.UpdateMenuOrder([]model.MenuOrderRequest{
		{ID: 1, SortOrder: 2},
		{ID: 2, SortOrder: 1},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
