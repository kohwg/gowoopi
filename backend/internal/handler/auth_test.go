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

func init() {
	gin.SetMode(gin.TestMode)
}

func TestAuthHandler_CustomerLogin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authSvc := mock.NewMockAuthService(ctrl)
	resp := &model.AuthResponse{Role: "customer", StoreID: "store1"}
	tokens := &model.TokenPair{AccessToken: "at", RefreshToken: "rt"}
	authSvc.EXPECT().CustomerLogin(gomock.Any()).Return(resp, tokens, nil)

	h := NewAuthHandler(authSvc)
	r := gin.New()
	r.POST("/login", h.CustomerLogin)

	body, _ := json.Marshal(model.CustomerLoginRequest{StoreID: "store1", TableNumber: 1, Password: "1234"})
	req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	// Check cookies set
	cookies := w.Result().Cookies()
	found := 0
	for _, c := range cookies {
		if c.Name == "access_token" || c.Name == "refresh_token" {
			found++
		}
	}
	if found != 2 {
		t.Errorf("expected 2 token cookies, got %d", found)
	}
}

func TestAuthHandler_CustomerLogin_ValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authSvc := mock.NewMockAuthService(ctrl)
	h := NewAuthHandler(authSvc)
	r := gin.New()
	r.POST("/login", h.CustomerLogin)

	req := httptest.NewRequest("POST", "/login", bytes.NewReader([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", w.Code)
	}
}

func TestAuthHandler_CustomerLogin_Unauthorized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authSvc := mock.NewMockAuthService(ctrl)
	authSvc.EXPECT().CustomerLogin(gomock.Any()).Return(nil, nil, model.ErrUnauthorized)

	h := NewAuthHandler(authSvc)
	r := gin.New()
	r.POST("/login", h.CustomerLogin)

	body, _ := json.Marshal(model.CustomerLoginRequest{StoreID: "s", TableNumber: 1, Password: "wrong"})
	req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("status = %d, want 401", w.Code)
	}
}

func TestAuthHandler_AdminLogin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authSvc := mock.NewMockAuthService(ctrl)
	resp := &model.AuthResponse{Role: "admin", StoreID: "store1"}
	tokens := &model.TokenPair{AccessToken: "at", RefreshToken: "rt"}
	authSvc.EXPECT().AdminLogin(gomock.Any()).Return(resp, tokens, nil)

	h := NewAuthHandler(authSvc)
	r := gin.New()
	r.POST("/login", h.AdminLogin)

	body, _ := json.Marshal(model.AdminLoginRequest{StoreID: "store1", Username: "admin", Password: "admin123"})
	req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
}

func TestAuthHandler_RefreshToken_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authSvc := mock.NewMockAuthService(ctrl)
	authSvc.EXPECT().RefreshToken("rt").Return("new-at", nil)

	h := NewAuthHandler(authSvc)
	r := gin.New()
	r.POST("/refresh", h.RefreshToken)

	req := httptest.NewRequest("POST", "/refresh", nil)
	req.AddCookie(&http.Cookie{Name: "refresh_token", Value: "rt"})
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
}

func TestAuthHandler_RefreshToken_NoCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authSvc := mock.NewMockAuthService(ctrl)
	h := NewAuthHandler(authSvc)
	r := gin.New()
	r.POST("/refresh", h.RefreshToken)

	req := httptest.NewRequest("POST", "/refresh", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("status = %d, want 401", w.Code)
	}
}
