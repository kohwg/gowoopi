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

// GetMenus godoc
// @Summary 메뉴 목록 조회
// @Description 매장의 카테고리별 메뉴 목록
// @Tags Menu
// @Produce json
// @Success 200 {array} model.Category
// @Failure 401 {object} model.ErrorResponse
// @Router /api/customer/menus [get]
// @Security BearerAuth
func (h *MenuHandler) GetMenus(c *gin.Context) {
	claims := middleware.GetClaims(c)
	menus, err := h.menuSvc.GetMenusByStore(claims.StoreID)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, menus)
}

// CreateMenu godoc
// @Summary 메뉴 생성
// @Description 새 메뉴 추가 (관리자)
// @Tags Menu
// @Accept json
// @Produce json
// @Param request body model.MenuCreateRequest true "메뉴 정보"
// @Success 201 {object} model.Menu
// @Failure 400 {object} model.ErrorResponse
// @Router /api/admin/menus [post]
// @Security BearerAuth
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

// UpdateMenu godoc
// @Summary 메뉴 수정
// @Description 메뉴 정보 수정 (관리자)
// @Tags Menu
// @Accept json
// @Produce json
// @Param id path int true "메뉴 ID"
// @Param request body model.MenuUpdateRequest true "수정할 정보"
// @Success 200 {object} model.Menu
// @Failure 400,404 {object} model.ErrorResponse
// @Router /api/admin/menus/{id} [put]
// @Security BearerAuth
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

// DeleteMenu godoc
// @Summary 메뉴 삭제
// @Description 메뉴 삭제 (관리자)
// @Tags Menu
// @Param id path int true "메뉴 ID"
// @Success 204
// @Failure 400,404 {object} model.ErrorResponse
// @Router /api/admin/menus/{id} [delete]
// @Security BearerAuth
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

// UpdateMenuOrder godoc
// @Summary 메뉴 순서 변경
// @Description 메뉴 정렬 순서 변경 (관리자)
// @Tags Menu
// @Accept json
// @Param request body []model.MenuOrderRequest true "순서 정보"
// @Success 204
// @Failure 400 {object} model.ErrorResponse
// @Router /api/admin/menus/order [patch]
// @Security BearerAuth
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
