package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kohwg/gowoopi/backend/internal/middleware"
	"github.com/kohwg/gowoopi/backend/internal/model"
	"github.com/kohwg/gowoopi/backend/internal/service"
)

type MenuHandler struct {
	menuSvc service.MenuService
}

func NewMenuHandler(menuSvc service.MenuService) *MenuHandler {
	return &MenuHandler{menuSvc: menuSvc}
}

func (h *MenuHandler) GetMenus(c *gin.Context) {
	claims := middleware.GetClaims(c)
	menus, err := h.menuSvc.GetMenusByStore(claims.StoreID)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, menus)
}

func (h *MenuHandler) CreateMenu(c *gin.Context) {
	claims := middleware.GetClaims(c)
	var req model.MenuCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: model.ErrorDetail{Code: "VALIDATION_ERROR", Message: err.Error()}})
		return
	}
	menu, err := h.menuSvc.CreateMenu(claims.StoreID, req)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, menu)
}

func (h *MenuHandler) UpdateMenu(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: model.ErrorDetail{Code: "VALIDATION_ERROR", Message: "잘못된 메뉴 ID"}})
		return
	}
	var req model.MenuUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: model.ErrorDetail{Code: "VALIDATION_ERROR", Message: err.Error()}})
		return
	}
	menu, err := h.menuSvc.UpdateMenu(uint(id), req)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, menu)
}

func (h *MenuHandler) DeleteMenu(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: model.ErrorDetail{Code: "VALIDATION_ERROR", Message: "잘못된 메뉴 ID"}})
		return
	}
	if err := h.menuSvc.DeleteMenu(uint(id)); err != nil {
		handleError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *MenuHandler) UpdateMenuOrder(c *gin.Context) {
	var items []model.MenuOrderRequest
	if err := c.ShouldBindJSON(&items); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: model.ErrorDetail{Code: "VALIDATION_ERROR", Message: err.Error()}})
		return
	}
	if err := h.menuSvc.UpdateMenuOrder(items); err != nil {
		handleError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
