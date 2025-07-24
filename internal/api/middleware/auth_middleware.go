package middleware

import (
	"net/http"
	"strings"

	"github.com/Fuzz-Head/domain/models"
	"github.com/Fuzz-Head/pkg/jwtutils"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwtutils.ParseToken(tokenStr)

		if err != nil || claims.UserID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Set("userClaims", *claims)
		c.Next()
	}
}

func GetClaims(c *gin.Context) (*models.UserClaims, bool) {
	val, exists := c.Get("userClaims")
	if !exists {
		return nil, false // errors.New("user claims not found in context")
	}
	claims, ok := val.(models.UserClaims)
	if !ok {
		return nil, false //errors.New("invalid user claims type")
	}
	return &claims, ok
}

func RequireAdminRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := GetClaims(c)
		if !ok || claims.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
