package middleware

import (
	"net/http"
	"strings"

	"github.com/Fuzz-Head/models"
	"github.com/gin-gonic/gin"
)

const UserClaimsKey = "userClaims"

func InjectClaims() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mock source
		role := c.GetHeader("X-Role")
		scopeHeader := c.GetHeader("X-Scopes")

		scopes := strings.Split(scopeHeader, ",")
		for i := range scopes {
			scopes[i] = strings.TrimSpace(scopes[i])
		}

		claims := models.UserClaims{
			Role:   role,
			Scopes: scopes,
		}

		c.Set(UserClaimsKey, claims)
		c.Next()
	}
}

func RoleRequired(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetHeader("X-Role")
		for _, allowed := range allowedRoles {
			if role == allowed {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied for role:" + role})
		c.Abort()
	}
}

func ScopeRequired(requiredScopes ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exists := c.Get(UserClaimsKey)
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Missing claims"})
			c.Abort()
			return
		}

		claims := value.(models.UserClaims)
		scopesSet := make(map[string]struct{})
		for _, s := range claims.Scopes {
			scopesSet[s] = struct{}{}
		}

		for _, required := range requiredScopes {
			if _, ok := scopesSet[required]; !ok {
				c.JSON(http.StatusForbidden, gin.H{"error": "Missing required scope: " + required})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
