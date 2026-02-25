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

func setCustomerClaims(c *gin.Context) {
	c.Set("claims", &model.Claims{StoreID: "store1", Role: "customer", TableID: 1, SessionID: "sess1"})
	c.Next()
}

func setAdminClaims(c *gin.Context) {
	c.Set("claims", &model.Claims{StoreID: "store1", Role: "admin", AdminID: 1})
	c.Next()
}

func TestOrderHandler_CreateOrder_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orderSvc := mock.NewMockOrderService(ctrl)
	order := &model.Order{ID: "o1", Status: model.OrderStatusPending}
	orderSvc.EXPECT().CreateOrder("store1", "sess1", uint(1), gomock.Any()).Return(order, nil)

	h := NewOrderHandler(orderSvc)
	r := gin.New()
	r.Use(setCustomerClaims)
	r.POST("/orders", h.CreateOrder)

	body, _ := json.Marshal(model.OrderCreateRequest{Items: []model.OrderItemRequest{{MenuID: 1, Quantity: 2}}})
	req := httptest.NewRequest("POST", "/orders", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("status = %d, want 201", w.Code)
	}
}

func TestOrderHandler_CreateOrder_ValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orderSvc := mock.NewMockOrderService(ctrl)
	h := NewOrderHandler(orderSvc)
	r := gin.New()
	r.Use(setCustomerClaims)
	r.POST("/orders", h.CreateOrder)

	req := httptest.NewRequest("POST", "/orders", bytes.NewReader([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", w.Code)
	}
}

func TestOrderHandler_GetCustomerOrders(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orderSvc := mock.NewMockOrderService(ctrl)
	orderSvc.EXPECT().GetOrdersBySession("sess1").Return([]model.Order{{ID: "o1"}}, nil)

	h := NewOrderHandler(orderSvc)
	r := gin.New()
	r.Use(setCustomerClaims)
	r.GET("/orders", h.GetCustomerOrders)

	req := httptest.NewRequest("GET", "/orders", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
}

func TestOrderHandler_GetAdminOrders(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orderSvc := mock.NewMockOrderService(ctrl)
	orderSvc.EXPECT().GetOrdersByStore("store1").Return([]model.Order{{ID: "o1"}}, nil)

	h := NewOrderHandler(orderSvc)
	r := gin.New()
	r.Use(setAdminClaims)
	r.GET("/orders", h.GetAdminOrders)

	req := httptest.NewRequest("GET", "/orders", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
}

func TestOrderHandler_UpdateOrderStatus_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orderSvc := mock.NewMockOrderService(ctrl)
	orderSvc.EXPECT().UpdateOrderStatus("o1", gomock.Any()).Return(&model.Order{ID: "o1", Status: model.OrderStatusConfirmed}, nil)

	h := NewOrderHandler(orderSvc)
	r := gin.New()
	r.PATCH("/orders/:id/status", h.UpdateOrderStatus)

	body, _ := json.Marshal(model.StatusUpdateRequest{Status: "CONFIRMED"})
	req := httptest.NewRequest("PATCH", "/orders/o1/status", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
}

func TestOrderHandler_DeleteOrder_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orderSvc := mock.NewMockOrderService(ctrl)
	orderSvc.EXPECT().DeleteOrder("o1").Return(nil)

	h := NewOrderHandler(orderSvc)
	r := gin.New()
	r.DELETE("/orders/:id", h.DeleteOrder)

	req := httptest.NewRequest("DELETE", "/orders/o1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("status = %d, want 204", w.Code)
	}
}

func TestOrderHandler_DeleteOrder_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orderSvc := mock.NewMockOrderService(ctrl)
	orderSvc.EXPECT().DeleteOrder("o1").Return(model.WrapNotFound("주문"))

	h := NewOrderHandler(orderSvc)
	r := gin.New()
	r.DELETE("/orders/:id", h.DeleteOrder)

	req := httptest.NewRequest("DELETE", "/orders/o1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("status = %d, want 404", w.Code)
	}
}
