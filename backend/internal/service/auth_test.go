package service

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gowoopi/backend/internal/mock"
	"github.com/gowoopi/backend/internal/model"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func newAuthTestService(t *testing.T) (*gomock.Controller, *mock.MockStoreRepository, *mock.MockTableRepository, *mock.MockSessionRepository, *mock.MockAdminRepository, AuthService) {
	ctrl := gomock.NewController(t)
	storeRepo := mock.NewMockStoreRepository(ctrl)
	tableRepo := mock.NewMockTableRepository(ctrl)
	sessionRepo := mock.NewMockSessionRepository(ctrl)
	adminRepo := mock.NewMockAdminRepository(ctrl)
	svc := NewAuthService(storeRepo, tableRepo, sessionRepo, adminRepo, "test-secret")
	return ctrl, storeRepo, tableRepo, sessionRepo, adminRepo, svc
}

func TestAuthService_CustomerLogin_Success(t *testing.T) {
	ctrl, _, tableRepo, sessionRepo, _, svc := newAuthTestService(t)
	defer ctrl.Finish()

	hash, _ := bcrypt.GenerateFromPassword([]byte("1234"), 10)
	tbl := &model.Table{ID: 1, StoreID: "store1", TableNumber: 1, PasswordHash: string(hash)}
	session := &model.TableSession{ID: "sess1", TableID: 1, StoreID: "store1", IsActive: true}

	tableRepo.EXPECT().FindByStoreAndNumber("store1", 1).Return(tbl, nil)
	sessionRepo.EXPECT().FindActiveByTable(uint(1)).Return(session, nil)

	resp, tokens, err := svc.CustomerLogin(model.CustomerLoginRequest{StoreID: "store1", TableNumber: 1, Password: "1234"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Role != "customer" {
		t.Errorf("role = %q, want customer", resp.Role)
	}
	if tokens.AccessToken == "" || tokens.RefreshToken == "" {
		t.Error("tokens should not be empty")
	}
}

func TestAuthService_CustomerLogin_WrongPassword(t *testing.T) {
	ctrl, _, tableRepo, _, _, svc := newAuthTestService(t)
	defer ctrl.Finish()

	hash, _ := bcrypt.GenerateFromPassword([]byte("1234"), 10)
	tbl := &model.Table{ID: 1, PasswordHash: string(hash)}
	tableRepo.EXPECT().FindByStoreAndNumber("store1", 1).Return(tbl, nil)

	_, _, err := svc.CustomerLogin(model.CustomerLoginRequest{StoreID: "store1", TableNumber: 1, Password: "wrong"})
	if err != model.ErrUnauthorized {
		t.Errorf("err = %v, want ErrUnauthorized", err)
	}
}

func TestAuthService_CustomerLogin_TableNotFound(t *testing.T) {
	ctrl, _, tableRepo, _, _, svc := newAuthTestService(t)
	defer ctrl.Finish()

	tableRepo.EXPECT().FindByStoreAndNumber("store1", 99).Return(nil, model.ErrNotFound)

	_, _, err := svc.CustomerLogin(model.CustomerLoginRequest{StoreID: "store1", TableNumber: 99, Password: "1234"})
	if err != model.ErrUnauthorized {
		t.Errorf("err = %v, want ErrUnauthorized", err)
	}
}

func TestAuthService_CustomerLogin_CreatesSession(t *testing.T) {
	ctrl, _, tableRepo, sessionRepo, _, svc := newAuthTestService(t)
	defer ctrl.Finish()

	hash, _ := bcrypt.GenerateFromPassword([]byte("1234"), 10)
	tbl := &model.Table{ID: 1, StoreID: "store1", PasswordHash: string(hash)}

	tableRepo.EXPECT().FindByStoreAndNumber("store1", 1).Return(tbl, nil)
	sessionRepo.EXPECT().FindActiveByTable(uint(1)).Return(nil, model.ErrNotFound)
	sessionRepo.EXPECT().Create(gomock.Any()).Return(nil)

	_, _, err := svc.CustomerLogin(model.CustomerLoginRequest{StoreID: "store1", TableNumber: 1, Password: "1234"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAuthService_AdminLogin_Success(t *testing.T) {
	ctrl, _, _, _, adminRepo, svc := newAuthTestService(t)
	defer ctrl.Finish()

	hash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), 10)
	admin := &model.Admin{ID: 1, StoreID: "store1", Username: "admin", PasswordHash: string(hash)}
	adminRepo.EXPECT().FindByStoreAndUsername("store1", "admin").Return(admin, nil)

	resp, tokens, err := svc.AdminLogin(model.AdminLoginRequest{StoreID: "store1", Username: "admin", Password: "admin123"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Role != "admin" {
		t.Errorf("role = %q, want admin", resp.Role)
	}
	if tokens.AccessToken == "" {
		t.Error("access token should not be empty")
	}
}

func TestAuthService_AdminLogin_WrongPassword(t *testing.T) {
	ctrl, _, _, _, adminRepo, svc := newAuthTestService(t)
	defer ctrl.Finish()

	hash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), 10)
	admin := &model.Admin{ID: 1, PasswordHash: string(hash)}
	adminRepo.EXPECT().FindByStoreAndUsername("store1", "admin").Return(admin, nil)

	_, _, err := svc.AdminLogin(model.AdminLoginRequest{StoreID: "store1", Username: "admin", Password: "wrong"})
	if err != model.ErrUnauthorized {
		t.Errorf("err = %v, want ErrUnauthorized", err)
	}
}

func TestAuthService_RefreshToken_Success(t *testing.T) {
	ctrl, _, _, _, _, svc := newAuthTestService(t)
	defer ctrl.Finish()

	claims := model.Claims{StoreID: "store1", Role: "customer"}
	tokens, _ := svc.GenerateTokenPair(claims)

	newAccess, err := svc.RefreshToken(tokens.RefreshToken)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if newAccess == "" {
		t.Error("new access token should not be empty")
	}
}

func TestAuthService_RefreshToken_InvalidToken(t *testing.T) {
	ctrl, _, _, _, _, svc := newAuthTestService(t)
	defer ctrl.Finish()

	_, err := svc.RefreshToken("invalid-token")
	if err != model.ErrUnauthorized {
		t.Errorf("err = %v, want ErrUnauthorized", err)
	}
}

func TestAuthService_ValidateToken_Success(t *testing.T) {
	ctrl, _, _, _, _, svc := newAuthTestService(t)
	defer ctrl.Finish()

	claims := model.Claims{StoreID: "store1", Role: "admin", AdminID: 5}
	tokens, _ := svc.GenerateTokenPair(claims)

	parsed, err := svc.ValidateToken(tokens.AccessToken)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if parsed.StoreID != "store1" || parsed.Role != "admin" || parsed.AdminID != 5 {
		t.Errorf("claims mismatch: %+v", parsed)
	}
}

func TestAuthService_ValidateToken_Expired(t *testing.T) {
	ctrl, _, _, _, _, svc := newAuthTestService(t)
	defer ctrl.Finish()

	claims := model.Claims{
		StoreID: "store1", Role: "customer",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().Add(-2 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString([]byte("test-secret"))

	_, err := svc.ValidateToken(tokenStr)
	if err != model.ErrUnauthorized {
		t.Errorf("err = %v, want ErrUnauthorized", err)
	}
}
