package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kohwg/gowoopi/backend/internal/mock"
	"github.com/kohwg/gowoopi/backend/internal/model"
	"go.uber.org/mock/gomock"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestAuthMiddleware_ValidCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authSvc := mock.NewMockAuthService(ctrl)
	claims := &model.Claims{StoreID: "store1", Role: "customer"}
	authSvc.EXPECT().ValidateToken("valid-token").Return(claims, nil)

	r := gin.New()
	r.Use(AuthMiddleware(authSvc))
	r.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/test", nil)
	req.AddCookie(&http.Cookie{Name: "access_token", Value: "valid-token"})
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
}

func TestAuthMiddleware_ValidBearerHeader(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authSvc := mock.NewMockAuthService(ctrl)
	claims := &model.Claims{StoreID: "store1", Role: "admin"}
	authSvc.EXPECT().ValidateToken("bearer-token").Return(claims, nil)

	r := gin.New()
	r.Use(AuthMiddleware(authSvc))
	r.GET("/test", func(c *gin.Context) { c.Status(http.StatusOK) })

	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer bearer-token")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
}

func TestAuthMiddleware_NoToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authSvc := mock.NewMockAuthService(ctrl)

	r := gin.New()
	r.Use(AuthMiddleware(authSvc))
	r.GET("/test", func(c *gin.Context) { c.Status(http.StatusOK) })

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("status = %d, want 401", w.Code)
	}
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authSvc := mock.NewMockAuthService(ctrl)
	authSvc.EXPECT().ValidateToken("bad-token").Return(nil, model.ErrUnauthorized)

	r := gin.New()
	r.Use(AuthMiddleware(authSvc))
	r.GET("/test", func(c *gin.Context) { c.Status(http.StatusOK) })

	req := httptest.NewRequest("GET", "/test", nil)
	req.AddCookie(&http.Cookie{Name: "access_token", Value: "bad-token"})
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("status = %d, want 401", w.Code)
	}
}

func TestRequireRole_Allowed(t *testing.T) {
	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Set("claims", &model.Claims{Role: "admin"})
		c.Next()
	})
	r.Use(RequireRole("admin"))
	r.GET("/test", func(c *gin.Context) { c.Status(http.StatusOK) })

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
}

func TestRequireRole_Forbidden(t *testing.T) {
	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Set("claims", &model.Claims{Role: "customer"})
		c.Next()
	})
	r.Use(RequireRole("admin"))
	r.GET("/test", func(c *gin.Context) { c.Status(http.StatusOK) })

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusForbidden {
		t.Errorf("status = %d, want 403", w.Code)
	}
}

func TestRequireRole_NoClaims(t *testing.T) {
	r := gin.New()
	r.Use(RequireRole("admin"))
	r.GET("/test", func(c *gin.Context) { c.Status(http.StatusOK) })

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("status = %d, want 401", w.Code)
	}
}
