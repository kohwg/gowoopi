package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kohwg/gowoopi/backend/internal/middleware"
	"github.com/kohwg/gowoopi/backend/internal/model"
	"github.com/kohwg/gowoopi/backend/internal/service"
)

type OrderHandler struct {
	orderSvc service.OrderService
}

func NewOrderHandler(orderSvc service.OrderService) *OrderHandler {
	return &OrderHandler{orderSvc: orderSvc}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	claims := middleware.GetClaims(c)
	var req model.OrderCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: model.ErrorDetail{Code: "VALIDATION_ERROR", Message: err.Error()}})
		return
	}
	order, err := h.orderSvc.CreateOrder(claims.StoreID, claims.SessionID, claims.TableID, req)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) GetCustomerOrders(c *gin.Context) {
	claims := middleware.GetClaims(c)
	orders, err := h.orderSvc.GetOrdersBySession(claims.SessionID)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) GetAdminOrders(c *gin.Context) {
	claims := middleware.GetClaims(c)
	orders, err := h.orderSvc.GetOrdersByStore(claims.StoreID)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")
	var req model.StatusUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: model.ErrorDetail{Code: "VALIDATION_ERROR", Message: err.Error()}})
		return
	}
	order, err := h.orderSvc.UpdateOrderStatus(id, req)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if err := h.orderSvc.DeleteOrder(id); err != nil {
		handleError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
