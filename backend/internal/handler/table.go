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
