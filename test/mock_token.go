package test

import (
	"strconv"
	"time"

	"github.com/Fuzz-Head/domain/models"
	"github.com/Fuzz-Head/internal/common/authutil"
	"github.com/golang-jwt/jwt/v5"
)

var testJWTSecret = []byte("test-key-secret")

func GenerateMockAccessToken(userID uint, role string) (string, error) {
	claims := models.UserClaims{
		UserID: strconv.Itoa(int(userID)),
		Role:   role,
		Scopes: authutil.ScopesForRole(role),
		Type:   "access_token",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(testJWTSecret))
}

func GenerateMockRefreshToken(userID uint, role string) (string, error) {
	claims := models.UserClaims{
		UserID: strconv.Itoa(int(userID)),
		Role:   role,
		Scopes: authutil.ScopesForRole(role),
		Type:   "refresh_token",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(testJWTSecret))
}
