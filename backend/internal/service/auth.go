package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kohwg/gowoopi/backend/internal/model"
	"github.com/kohwg/gowoopi/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	storeRepo   repository.StoreRepository
	tableRepo   repository.TableRepository
	sessionRepo repository.SessionRepository
	adminRepo   repository.AdminRepository
	jwtSecret   []byte
}

func NewAuthService(storeRepo repository.StoreRepository, tableRepo repository.TableRepository, sessionRepo repository.SessionRepository, adminRepo repository.AdminRepository, jwtSecret string) AuthService {
	return &authService{storeRepo: storeRepo, tableRepo: tableRepo, sessionRepo: sessionRepo, adminRepo: adminRepo, jwtSecret: []byte(jwtSecret)}
}

func (s *authService) CustomerLogin(req model.CustomerLoginRequest) (*model.AuthResponse, *model.TokenPair, error) {
	tbl, err := s.tableRepo.FindByStoreAndNumber(req.StoreID, req.TableNumber)
	if err != nil {
		return nil, nil, model.ErrUnauthorized
	}
	if err := bcrypt.CompareHashAndPassword([]byte(tbl.PasswordHash), []byte(req.Password)); err != nil {
		return nil, nil, model.ErrUnauthorized
	}

	session, err := s.sessionRepo.FindActiveByTable(tbl.ID)
	if err != nil {
		session = &model.TableSession{TableID: tbl.ID, StoreID: req.StoreID, StartedAt: time.Now(), IsActive: true}
		if err := s.sessionRepo.Create(session); err != nil {
			return nil, nil, model.ErrInternal
		}
	}

	claims := model.Claims{StoreID: req.StoreID, Role: "customer", TableID: tbl.ID, SessionID: session.ID}
	tokens, err := s.GenerateTokenPair(claims)
	if err != nil {
		return nil, nil, model.ErrInternal
	}

	return &model.AuthResponse{Role: "customer", StoreID: req.StoreID}, tokens, nil
}

func (s *authService) AdminLogin(req model.AdminLoginRequest) (*model.AuthResponse, *model.TokenPair, error) {
	admin, err := s.adminRepo.FindByStoreAndUsername(req.StoreID, req.Username)
	if err != nil {
		return nil, nil, model.ErrUnauthorized
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(req.Password)); err != nil {
		return nil, nil, model.ErrUnauthorized
	}

	claims := model.Claims{StoreID: req.StoreID, Role: "admin", AdminID: admin.ID}
	tokens, err := s.GenerateTokenPair(claims)
	if err != nil {
		return nil, nil, model.ErrInternal
	}

	return &model.AuthResponse{Role: "admin", StoreID: req.StoreID}, tokens, nil
}

func (s *authService) RefreshToken(refreshToken string) (string, error) {
	claims, err := s.ValidateToken(refreshToken)
	if err != nil {
		return "", model.ErrUnauthorized
	}

	now := time.Now()
	claims.ExpiresAt = jwt.NewNumericDate(now.Add(15 * time.Minute))
	claims.IssuedAt = jwt.NewNumericDate(now)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}

func (s *authService) GenerateTokenPair(claims model.Claims) (*model.TokenPair, error) {
	now := time.Now()

	claims.IssuedAt = jwt.NewNumericDate(now)
	claims.ExpiresAt = jwt.NewNumericDate(now.Add(15 * time.Minute))
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	access, err := accessToken.SignedString(s.jwtSecret)
	if err != nil {
		return nil, err
	}

	claims.ExpiresAt = jwt.NewNumericDate(now.Add(30 * 24 * time.Hour))
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refresh, err := refreshToken.SignedString(s.jwtSecret)
	if err != nil {
		return nil, err
	}

	return &model.TokenPair{AccessToken: access, RefreshToken: refresh}, nil
}

func (s *authService) ValidateToken(tokenStr string) (*model.Claims, error) {
	claims := &model.Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return s.jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, model.ErrUnauthorized
	}
	return claims, nil
}
