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

// CreateOrder godoc
// @Summary 주문 생성
// @Description 새 주문 생성 (고객)
// @Tags Order
// @Accept json
// @Produce json
// @Param request body model.OrderCreateRequest true "주문 정보"
// @Success 201 {object} model.Order
// @Failure 400 {object} model.ErrorResponse
// @Router /api/customer/orders [post]
// @Security BearerAuth
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

// GetCustomerOrders godoc
// @Summary 내 주문 조회
// @Description 현재 세션의 주문 목록 (고객)
// @Tags Order
// @Produce json
// @Success 200 {array} model.Order
// @Router /api/customer/orders [get]
// @Security BearerAuth
func (h *OrderHandler) GetCustomerOrders(c *gin.Context) {
	claims := middleware.GetClaims(c)
	orders, err := h.orderSvc.GetOrdersBySession(claims.SessionID)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, orders)
}

// GetAdminOrders godoc
// @Summary 매장 주문 조회
// @Description 매장의 전체 주문 목록 (관리자)
// @Tags Order
// @Produce json
// @Success 200 {array} model.Order
// @Router /api/admin/orders [get]
// @Security BearerAuth
func (h *OrderHandler) GetAdminOrders(c *gin.Context) {
	claims := middleware.GetClaims(c)
	orders, err := h.orderSvc.GetOrdersByStore(claims.StoreID)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, orders)
}

// UpdateOrderStatus godoc
// @Summary 주문 상태 변경
// @Description 주문 상태 업데이트 (관리자)
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "주문 ID"
// @Param request body model.StatusUpdateRequest true "상태 정보"
// @Success 200 {object} model.Order
// @Failure 400,404 {object} model.ErrorResponse
// @Router /api/admin/orders/{id}/status [patch]
// @Security BearerAuth
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

// DeleteOrder godoc
// @Summary 주문 삭제
// @Description 주문 삭제 (관리자)
// @Tags Order
// @Param id path string true "주문 ID"
// @Success 204
// @Failure 404 {object} model.ErrorResponse
// @Router /api/admin/orders/{id} [delete]
// @Security BearerAuth
func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if err := h.orderSvc.DeleteOrder(id); err != nil {
		handleError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
