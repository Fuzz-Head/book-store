package handlers

import (
	"net/http"

	"github.com/Fuzz-Head/database"
	"github.com/Fuzz-Head/domain/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreatePublisherInput struct {
	Name     string  `json:"name" binding:"required"`
	Website  *string `json:"website"`
	Location *string `json:"location"`
	Founded  *int    `json:"founded"`
}

func GetPublishers(c *gin.Context) {
	var publishers []models.Publisher

	if err := database.DB.Find(&publishers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch publishers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"publishers": publishers})
}

func CreatePublisher(c *gin.Context) {
	var input CreatePublisherInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	publisher := models.Publisher{
		ID:       uuid.New().String(),
		Name:     input.Name,
		Website:  *input.Website,
		Location: *input.Location,
		Founded:  input.Founded,
	}

	if err := database.DB.Create(&publisher).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create publisher"})
		return
	}

	c.JSON(http.StatusCreated, publisher)
}
