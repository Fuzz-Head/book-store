package models

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	UserID string   `json:"user_id"`
	Role   string   `json:"role"`
	Scopes []string `json:"scopes"`
	Type   string   `json:"type"`
	jwt.RegisteredClaims
}

//type CustomClaims struct {
//	UserID string `json:"user_id"`
//	Role   string `json:"role"`
//	jwt.RegisteredClaims
//}
