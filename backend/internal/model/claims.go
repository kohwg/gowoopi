package model

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	StoreID   string `json:"store_id"`
	Role      string `json:"role"`
	TableID   uint   `json:"table_id,omitempty"`
	SessionID string `json:"session_id,omitempty"`
	AdminID   uint   `json:"admin_id,omitempty"`
	jwt.RegisteredClaims
}
