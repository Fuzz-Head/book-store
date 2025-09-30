// internal/api/handlers/startup_handler.go
package handlers

import (
	"net/http"
	"time"

	"github.com/Fuzz-Head/database"
	"github.com/gin-gonic/gin"
)

type StartupResponse struct {
	Status      string                 `json:"status"`
	Timestamp   int64                  `json:"timestamp"`
	StartupTime time.Time              `json:"startup_time"`
	Uptime      string                 `json:"uptime"`
	Environment string                 `json:"environment"`
	Checks      map[string]interface{} `json:"checks"`
}

var startupTime = time.Now()

// StartupCheck - Provides detailed startup information
// GET /startup
func StartupCheck(c *gin.Context) {
	uptime := time.Since(startupTime)

	checks := map[string]interface{}{
		"database_migration": checkDatabaseTables(),
		"seed_data":          checkSeedData(),
		"configuration":      checkConfiguration(),
	}

	response := StartupResponse{
		Status:      StatusHealthy,
		Timestamp:   time.Now().Unix(),
		StartupTime: startupTime,
		Uptime:      uptime.String(),
		Environment: getEnvironment(),
		Checks:      checks,
	}

	c.JSON(http.StatusOK, response)
}

func checkDatabaseTables() map[string]interface{} {
	if database.DB == nil {
		return map[string]interface{}{
			"status": StatusDown,
			"error":  "Database not connected",
		}
	}

	// Check if main tables exist
	tables := []string{"books", "authors", "publishers", "users"}
	existingTables := make(map[string]bool)

	for _, table := range tables {
		var count int64
		err := database.DB.Table(table).Count(&count).Error
		existingTables[table] = err == nil
	}

	return map[string]interface{}{
		"status": StatusUp,
		"tables": existingTables,
	}
}

func checkSeedData() map[string]interface{} {
	if database.DB == nil {
		return map[string]interface{}{
			"status": StatusDown,
			"error":  "Database not connected",
		}
	}

	var bookCount, authorCount, publisherCount int64

	database.DB.Table("books").Count(&bookCount)
	database.DB.Table("authors").Count(&authorCount)
	database.DB.Table("publishers").Count(&publisherCount)

	return map[string]interface{}{
		"status": StatusUp,
		"counts": map[string]int64{
			"books":      bookCount,
			"authors":    authorCount,
			"publishers": publisherCount,
		},
	}
}

func checkConfiguration() map[string]interface{} {
	return map[string]interface{}{
		"status":         StatusUp,
		"jwt_configured": len(getEnvOrDefault("JWT_SECRET", "")) > 0,
		"db_configured":  len(getEnvOrDefault("DB_HOST", "")) > 0,
	}
}

func getEnvironment() string {
	return getEnvOrDefault("ENVIRONMENT", "development")
}
