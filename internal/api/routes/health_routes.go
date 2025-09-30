package routes

import (
	"github.com/Fuzz-Head/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupHealthRoutes(router *gin.Engine) {
	router.GET("/health", handlers.HealthCheck)
	router.GET("/ready", handlers.ReadinessCheck)
	router.GET("/live", handlers.LivenessCheck)
	router.GET("/startup", handlers.StartupCheck)

	healthGroup := router.Group("/health")
	{
		healthGroup.GET("/", handlers.HealthCheck)
		healthGroup.GET("/ready", handlers.ReadinessCheck)
		healthGroup.GET("/live", handlers.LivenessCheck)
		healthGroup.GET("/startup", handlers.StartupCheck)
	}
}
