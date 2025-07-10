package test

import (
	// "github.com/Fuzz-Head/database"
	"github.com/Fuzz-Head/domain/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test DB")
	}

	// needed if no global function calls this function
	// database.DB = db

	db.AutoMigrate(&models.User{}, &models.Book{})
	return db
}

func MockAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("userClaims", models.UserClaims{
			Role: "admin",
			Scopes: []string{
				"can:read:books", "can:read:book", "can:create:book",
				"can:update:book", "can:delete:book",
			},
		})
		c.Next()
	}
}
