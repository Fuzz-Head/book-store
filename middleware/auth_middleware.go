package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/Fuzz-Head/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		role := claims["role"].(string)

		c.Set("userClaims", models.UserClaims{
			Role:   role,
			Scopes: scopesForRole(role),
		})
		c.Next()
	}
}

func scopesForRole(role string) []string {
	switch role {
	case "admin":
		return []string{
			"can:read:books",
			"can:read:book",
			"can:create:book",
			"can:update:book",
			"can:delete:book",
		}
	case "superUser":
		return []string{
			"can:read:books",
			"can:read:book",
			"can:create:book",
			"can:update:book",
		}
	case "user":
		return []string{
			"can:read:book",
			"can:read:books",
		}
	default:
		return []string{}
	}
}
