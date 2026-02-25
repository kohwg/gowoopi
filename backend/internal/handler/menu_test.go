package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gowoopi/backend/internal/mock"
	"github.com/gowoopi/backend/internal/model"
	"go.uber.org/mock/gomock"
)

func setClaims(c *gin.Context) {
	c.Set("claims", &model.Claims{StoreID: "store1", Role: "admin"})
	c.Next()
}

func TestMenuHandler_GetMenus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	menuSvc := mock.NewMockMenuService(ctrl)
	menus := []model.Menu{{ID: 1, Name: "김치찌개", Price: 9000}}
	menuSvc.EXPECT().GetMenusByStore("store1").Return(menus, nil)

	h := NewMenuHandler(menuSvc)
	r := gin.New()
	r.Use(setClaims)
	r.GET("/menus", h.GetMenus)

	req := httptest.NewRequest("GET", "/menus", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
}

func TestMenuHandler_CreateMenu_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	menuSvc := mock.NewMockMenuService(ctrl)
	menuSvc.EXPECT().CreateMenu("store1", gomock.Any()).Return(&model.Menu{ID: 1, Name: "된장찌개"}, nil)

	h := NewMenuHandler(menuSvc)
	r := gin.New()
	r.Use(setClaims)
	r.POST("/menus", h.CreateMenu)

	body, _ := json.Marshal(model.MenuCreateRequest{CategoryID: 1, Name: "된장찌개", Price: 8000})
	req := httptest.NewRequest("POST", "/menus", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("status = %d, want 201", w.Code)
	}
}

func TestMenuHandler_CreateMenu_ValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	menuSvc := mock.NewMockMenuService(ctrl)
	h := NewMenuHandler(menuSvc)
	r := gin.New()
	r.Use(setClaims)
	r.POST("/menus", h.CreateMenu)

	req := httptest.NewRequest("POST", "/menus", bytes.NewReader([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", w.Code)
	}
}

func TestMenuHandler_UpdateMenu_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	menuSvc := mock.NewMockMenuService(ctrl)
	name := "업데이트"
	menuSvc.EXPECT().UpdateMenu(uint(1), gomock.Any()).Return(&model.Menu{ID: 1, Name: name}, nil)

	h := NewMenuHandler(menuSvc)
	r := gin.New()
	r.PUT("/menus/:id", h.UpdateMenu)

	body, _ := json.Marshal(model.MenuUpdateRequest{Name: &name})
	req := httptest.NewRequest("PUT", "/menus/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
}

func TestMenuHandler_UpdateMenu_InvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	menuSvc := mock.NewMockMenuService(ctrl)
	h := NewMenuHandler(menuSvc)
	r := gin.New()
	r.PUT("/menus/:id", h.UpdateMenu)

	req := httptest.NewRequest("PUT", "/menus/abc", bytes.NewReader([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", w.Code)
	}
}

func TestMenuHandler_DeleteMenu_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	menuSvc := mock.NewMockMenuService(ctrl)
	menuSvc.EXPECT().DeleteMenu(uint(1)).Return(nil)

	h := NewMenuHandler(menuSvc)
	r := gin.New()
	r.DELETE("/menus/:id", h.DeleteMenu)

	req := httptest.NewRequest("DELETE", "/menus/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("status = %d, want 204", w.Code)
	}
}

func TestMenuHandler_UpdateMenuOrder_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	menuSvc := mock.NewMockMenuService(ctrl)
	menuSvc.EXPECT().UpdateMenuOrder(gomock.Any()).Return(nil)

	h := NewMenuHandler(menuSvc)
	r := gin.New()
	r.PATCH("/menus/order", h.UpdateMenuOrder)

	body, _ := json.Marshal([]model.MenuOrderRequest{{ID: 1, SortOrder: 2}})
	req := httptest.NewRequest("PATCH", "/menus/order", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("status = %d, want 204", w.Code)
	}
}
