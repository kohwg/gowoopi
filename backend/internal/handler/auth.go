package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kohwg/gowoopi/backend/internal/model"
	"github.com/kohwg/gowoopi/backend/internal/service"
)

type AuthHandler struct {
	authSvc service.AuthService
}

func NewAuthHandler(authSvc service.AuthService) *AuthHandler {
	return &AuthHandler{authSvc: authSvc}
}

// CustomerLogin godoc
// @Summary 고객 로그인
// @Description QR코드 스캔 후 테이블 세션 생성
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body model.CustomerLoginRequest true "로그인 요청"
// @Success 200 {object} model.AuthResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /api/customer/login [post]
func (h *AuthHandler) CustomerLogin(c *gin.Context) {
	var req model.CustomerLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: model.ErrorDetail{Code: "VALIDATION_ERROR", Message: err.Error()}})
		return
	}
	resp, tokens, err := h.authSvc.CustomerLogin(req)
	if err != nil {
		handleError(c, err)
		return
	}
	setTokenCookies(c, tokens)
	c.JSON(http.StatusOK, resp)
}

// AdminLogin godoc
// @Summary 관리자 로그인
// @Description 매장 관리자 로그인
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body model.AdminLoginRequest true "로그인 요청"
// @Success 200 {object} model.AuthResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Router /api/admin/login [post]
func (h *AuthHandler) AdminLogin(c *gin.Context) {
	var req model.AdminLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: model.ErrorDetail{Code: "VALIDATION_ERROR", Message: err.Error()}})
		return
	}
	resp, tokens, err := h.authSvc.AdminLogin(req)
	if err != nil {
		handleError(c, err)
		return
	}
	setTokenCookies(c, tokens)
	c.JSON(http.StatusOK, resp)
}

// RefreshToken godoc
// @Summary 토큰 갱신
// @Description Access token 갱신
// @Tags Auth
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 401 {object} model.ErrorResponse
// @Router /api/auth/refresh [post]
// @Security BearerAuth
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse{Error: model.ErrorDetail{Code: "UNAUTHORIZED", Message: "Refresh token이 없습니다"}})
		return
	}
	accessToken, err := h.authSvc.RefreshToken(refreshToken)
	if err != nil {
		handleError(c, err)
		return
	}
	c.SetCookie("access_token", accessToken, 900, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "토큰이 갱신되었습니다"})
}

func setTokenCookies(c *gin.Context, tokens *model.TokenPair) {
	c.SetCookie("access_token", tokens.AccessToken, 900, "/", "", false, true)
	c.SetCookie("refresh_token", tokens.RefreshToken, 2592000, "/", "", false, true)
}

func handleError(c *gin.Context, err error) {
	if appErr, ok := err.(*model.AppError); ok {
		c.JSON(appErr.Status, model.ErrorResponse{Error: model.ErrorDetail{Code: appErr.Code, Message: appErr.Message}})
		return
	}
	c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: model.ErrorDetail{Code: "INTERNAL_ERROR", Message: "서버 내부 오류"}})
}
