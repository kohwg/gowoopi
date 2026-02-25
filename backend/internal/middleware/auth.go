package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gowoopi/backend/internal/model"
	"github.com/gowoopi/backend/internal/service"
)

func AuthMiddleware(authSvc service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("access_token")
		if err != nil || tokenStr == "" {
			// fallback to Authorization header
			auth := c.GetHeader("Authorization")
			if strings.HasPrefix(auth, "Bearer ") {
				tokenStr = auth[7:]
			}
		}
		if tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{Error: model.ErrorDetail{Code: "UNAUTHORIZED", Message: "인증이 필요합니다"}})
			return
		}
		claims, err := authSvc.ValidateToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{Error: model.ErrorDetail{Code: "UNAUTHORIZED", Message: "유효하지 않은 토큰입니다"}})
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{Error: model.ErrorDetail{Code: "UNAUTHORIZED", Message: "인증이 필요합니다"}})
			return
		}
		if claims.(*model.Claims).Role != role {
			c.AbortWithStatusJSON(http.StatusForbidden, model.ErrorResponse{Error: model.ErrorDetail{Code: "FORBIDDEN", Message: "권한이 없습니다"}})
			return
		}
		c.Next()
	}
}

func GetClaims(c *gin.Context) *model.Claims {
	claims, _ := c.Get("claims")
	return claims.(*model.Claims)
}
