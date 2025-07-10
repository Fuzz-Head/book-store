package test

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var testJWTSecret = []byte("test-key-secret")

func GenerateMockAccessToken(userID uint, role string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"type":    "access",
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}).SignedString(testJWTSecret)
}

func GenerateMockRefreshToken(userID uint, role string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"type":    "refresh",
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}).SignedString(testJWTSecret)
}
