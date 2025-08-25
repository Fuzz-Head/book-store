package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Check struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

type HealthResponse struct {
	Status    string           `json:"status"`
	Timestamp int64            `json:"timestamp"`
	Version   string           `json:"version"`
	Service   string           `json:"sevice"`
	Checks    map[string]Check `json:"checks,omitempty"`
}

const (
	StatusHealthy   = "healthy"
	StatusUnhealthy = "unhealthy"
	StatusUp        = "up"
	StatusDown      = "down"
)

// Basic health check endpoint
func HealthCheck(c *gin.Context) {
	response := HealthResponse{
		Status:    StatusHealthy,
		Timestamp: time.Now().Unix(),
		Version:   "1.0.0",
		Service:   "bookstore-api",
	}

	c.JSON(http.StatusOK, response)
}
