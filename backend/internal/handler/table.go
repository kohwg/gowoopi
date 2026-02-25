package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kohwg/gowoopi/backend/internal/middleware"
	"github.com/kohwg/gowoopi/backend/internal/model"
	"github.com/kohwg/gowoopi/backend/internal/service"
)

type TableHandler struct {
	tableSvc service.TableService
}

func NewTableHandler(tableSvc service.TableService) *TableHandler {
	return &TableHandler{tableSvc: tableSvc}
}

// SetupTable godoc
// @Summary 테이블 설정
// @Description 새 테이블 생성 (관리자)
// @Tags Table
// @Accept json
// @Produce json
// @Param request body model.TableSetupRequest true "테이블 정보"
// @Success 201 {object} model.Table
// @Failure 400 {object} model.ErrorResponse
// @Router /api/admin/tables/setup [post]
// @Security BearerAuth
func (h *TableHandler) SetupTable(c *gin.Context) {
	claims := middleware.GetClaims(c)
	var req model.TableSetupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: model.ErrorDetail{Code: "VALIDATION_ERROR", Message: err.Error()}})
		return
	}
	tbl, err := h.tableSvc.SetupTable(claims.StoreID, req)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, tbl)
}

// CompleteTable godoc
// @Summary 테이블 이용 완료
// @Description 테이블 세션 종료 및 주문 히스토리 이동 (관리자)
// @Tags Table
// @Param id path int true "테이블 ID"
// @Success 200 {object} map[string]string
// @Failure 400,404 {object} model.ErrorResponse
// @Router /api/admin/tables/{id}/complete [post]
// @Security BearerAuth
func (h *TableHandler) CompleteTable(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: model.ErrorDetail{Code: "VALIDATION_ERROR", Message: "잘못된 테이블 ID"}})
		return
	}
	if err := h.tableSvc.CompleteTable(uint(id)); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "테이블 이용이 완료되었습니다"})
}

// GetTableHistory godoc
// @Summary 테이블 주문 히스토리
// @Description 테이블의 과거 주문 내역 조회 (관리자)
// @Tags Table
// @Produce json
// @Param id path int true "테이블 ID"
// @Param from query string false "시작일 (YYYY-MM-DD)"
// @Param to query string false "종료일 (YYYY-MM-DD)"
// @Success 200 {array} model.OrderHistory
// @Failure 400 {object} model.ErrorResponse
// @Router /api/admin/tables/{id}/history [get]
// @Security BearerAuth
func (h *TableHandler) GetTableHistory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: model.ErrorDetail{Code: "VALIDATION_ERROR", Message: "잘못된 테이블 ID"}})
		return
	}
	var from, to *time.Time
	if f := c.Query("from"); f != "" {
		if t, err := time.Parse("2006-01-02", f); err == nil {
			from = &t
		}
	}
	if t := c.Query("to"); t != "" {
		if parsed, err := time.Parse("2006-01-02", t); err == nil {
			to = &parsed
		}
	}
	histories, err := h.tableSvc.GetTableHistory(uint(id), from, to)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, histories)
}
