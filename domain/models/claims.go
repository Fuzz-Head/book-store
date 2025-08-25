package models

import "github.com/golang-jwt/jwt/v5"

type TokenType string

const (
	TokenTypeAccess  TokenType = "access_token"
	TokenTypeRefresh TokenType = "refresh_token"
)

type UserClaims struct {
	UserID string    `json:"user_id"`
	Role   string    `json:"role"`
	Scopes []string  `json:"scopes"`
	Type   TokenType `json:"type"`
	JTI    string    `json:"jti"`
	jwt.RegisteredClaims
}

//type CustomClaims struct {
//	UserID string `json:"user_id"`
//	Role   string `json:"role"`
//	jwt.RegisteredClaims
//}
