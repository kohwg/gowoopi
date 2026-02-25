package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kohwg/gowoopi/backend/internal/mock"
	"github.com/kohwg/gowoopi/backend/internal/model"
	"go.uber.org/mock/gomock"
)

func TestTableHandler_SetupTable_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tableSvc := mock.NewMockTableService(ctrl)
	tableSvc.EXPECT().SetupTable("store1", gomock.Any()).Return(&model.Table{ID: 1, TableNumber: 5, IsActive: true}, nil)

	h := NewTableHandler(tableSvc)
	r := gin.New()
	r.Use(setAdminClaims)
	r.POST("/tables/setup", h.SetupTable)

	body, _ := json.Marshal(model.TableSetupRequest{TableNumber: 5, Password: "1234"})
	req := httptest.NewRequest("POST", "/tables/setup", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("status = %d, want 201", w.Code)
	}
}

func TestTableHandler_SetupTable_ValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tableSvc := mock.NewMockTableService(ctrl)
	h := NewTableHandler(tableSvc)
	r := gin.New()
	r.Use(setAdminClaims)
	r.POST("/tables/setup", h.SetupTable)

	req := httptest.NewRequest("POST", "/tables/setup", bytes.NewReader([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", w.Code)
	}
}

func TestTableHandler_CompleteTable_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tableSvc := mock.NewMockTableService(ctrl)
	tableSvc.EXPECT().CompleteTable(uint(1)).Return(nil)

	h := NewTableHandler(tableSvc)
	r := gin.New()
	r.POST("/tables/:id/complete", h.CompleteTable)

	req := httptest.NewRequest("POST", "/tables/1/complete", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
}

func TestTableHandler_CompleteTable_InvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tableSvc := mock.NewMockTableService(ctrl)
	h := NewTableHandler(tableSvc)
	r := gin.New()
	r.POST("/tables/:id/complete", h.CompleteTable)

	req := httptest.NewRequest("POST", "/tables/abc/complete", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", w.Code)
	}
}

func TestTableHandler_GetTableHistory_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tableSvc := mock.NewMockTableService(ctrl)
	tableSvc.EXPECT().GetTableHistory(uint(1), nil, nil).Return([]model.OrderHistory{{ID: 1}}, nil)

	h := NewTableHandler(tableSvc)
	r := gin.New()
	r.GET("/tables/:id/history", h.GetTableHistory)

	req := httptest.NewRequest("GET", "/tables/1/history", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
}

func TestTableHandler_GetTableHistory_WithDateFilter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tableSvc := mock.NewMockTableService(ctrl)
	tableSvc.EXPECT().GetTableHistory(uint(1), gomock.Any(), gomock.Any()).Return([]model.OrderHistory{}, nil)

	h := NewTableHandler(tableSvc)
	r := gin.New()
	r.GET("/tables/:id/history", h.GetTableHistory)

	req := httptest.NewRequest("GET", "/tables/1/history?from=2026-01-01&to=2026-02-25", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
}
