package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Fuzz-Head/database"
	"github.com/Fuzz-Head/domain/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthorInput struct {
	FirstName   string     `json:"first_name" binding:"required"`
	LastName    string     `json:"last_name" binding:"required"`
	Biography   *string    `json:"biography"`
	BirthDate   *time.Time `json:"birth_date"`
	DeathDate   *time.Time `json:"death_date"`
	Nationality *string    `json:"nationality"`
	Website     *string    `json:"website"`
}

type AuthorUpdateInput struct {
	FirstName   *string    `json:"first_name"`
	LastName    *string    `json:"last_name"`
	Biography   *string    `json:"biography"`
	BirthDate   *time.Time `json:"birth_date"`
	DeathDate   *time.Time `json:"death_date"`
	Nationality *string    `json:"nationality"`
	Website     *string    `json:"website"`
}

func GetAuthors(c *gin.Context) {
	var authors []models.Author
	var totalCount int64

	limit := 10
	offset := 0

	if l := c.Query("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 {
			limit = parsed
		}
	}

	if o := c.Query("offset"); o != "" {
		if parsed, err := strconv.Atoi(o); err == nil && parsed >= 0 {
			offset = parsed
		}
	}

	// Check if deleted authors are to be included
	includeDeleted := c.Query("include_deleted") == "true"

	// build query - only active records
	baseQuery := database.DB.Model(&models.Author{})
	if !includeDeleted {
		baseQuery = baseQuery.Where("is_active = ?", true)
	}

	// other filtering
	if nationality := c.Query("nationality"); nationality != "" {
		baseQuery = baseQuery.Where("nationality = ?", nationality)
	}

	//if search := c.Query("search"); search != "" {
	//	baseQuery = baseQuery.Where("first_name ILIKE ? OR last_name ILIKE ?", "%"+search+"%", "%"+search+"%")
	//}

	if err := baseQuery.Count(&totalCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count authors"})
		return
	}

	if err := baseQuery.Limit(limit).Offset(offset).Find(&authors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch authors"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"authors":     authors,
		"limit":       limit,
		"offset":      offset,
		"count":       len(authors),
		"total_count": totalCount,
		"has_more":    offset+len(authors) < int(totalCount),
	})

}

func CreateAuthor(c *gin.Context) {
	var input AuthorInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author := models.Author{
		ID:          uuid.New().String(),
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Biography:   *input.Biography,
		BirthDate:   input.BirthDate,
		DeathDate:   input.DeathDate,
		Nationality: *input.Nationality,
		Website:     *input.Website,
	}

	if err := database.DB.Create(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create author"})
		return
	}

	c.JSON(http.StatusCreated, author)
}

func GetAuthor(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}

	var author models.Author

	if err := database.DB.First(&author, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}
	c.JSON(http.StatusOK, author)
}

func UpdateAuthor(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid author ID format"})
		return
	}

	var input AuthorUpdateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingAuthor models.Author
	if err := database.DB.First(&existingAuthor, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	updatedFields := struct {
		FirstName   *string    `json:"first_name"`
		LastName    *string    `json:"last_name"`
		Biography   *string    `json:"biography"`
		BirthDate   *time.Time `json:"birth_date"`
		DeathDate   *time.Time `json:"death_date"`
		Nationality *string    `json:"nationality"`
		Website     *string    `json:"website"`
		UpdatedAt   time.Time  `json:"updated_at"`
	}{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Biography:   input.Biography,
		BirthDate:   input.BirthDate,
		DeathDate:   input.DeathDate,
		Nationality: input.Nationality,
		Website:     input.Website,
		UpdatedAt:   time.Now(),
	}

	if err := database.DB.Model(&existingAuthor).Updates(updatedFields).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update author"})
		return
	}

	var updateAuthor models.Author
	database.DB.First(&updateAuthor, "id = ?", id)

	c.JSON(http.StatusOK, updateAuthor)

}

// A soft delete function for author data
func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid author ID format"})
		return
	}

	// need to get the current userId somehow
	userId := "1"

	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var existingAuthor models.Author
	if err := tx.Where("id = ? AND is_active = ?", id, true).First(&existingAuthor).Error; err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Author not found or already deleted"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete author"})
		}
		return
	}

	now := time.Now()
	updateFields := struct {
		IsActive  bool       `gorm:"column:is_active"`
		DeletedAt *time.Time `gorm:"column:deleted_at"`
		DeletedBy *string    `gorm:"column:deleted_by"`
		UpdatedAt time.Time  `gorm:"column:updated_at"`
	}{
		IsActive:  false,
		DeletedAt: &now,
		DeletedBy: &userId,
		UpdatedAt: now,
	}

	if err := tx.Model(&existingAuthor).Updates(updateFields).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete author"})
		return
	}

	if err := tx.Model(&models.BookAuthor{}).
		Where("author_id = ?", id).
		Updates(map[string]interface{}{
			"is_active":  false,
			"deleted_at": &now,
			"updated_at": now,
		}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete author"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete author"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Author deleted successfully",
		"deleted_at": now,
	})

}

// Re instaiate deleted authors
func ResotreAuthor(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}

	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var deletedAuthor models.Author
	if err := tx.Where("id = ? AND is_active = ?", id, false).First(&deletedAuthor).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Deleted author not found"})
		return
	}

	updateFields := struct {
		IsActive  bool       `gorm:"column:is_active"`
		DeletedAt *time.Time `gorm:"column:deleted_at"`
		DeletedBy *string    `gorm:"column:deleted_by"`
		UpdatedAt time.Time  `gorm:"column:updated_at"`
	}{
		IsActive:  true,
		DeletedAt: nil,
		DeletedBy: nil,
		UpdatedAt: time.Now(),
	}

	if err := tx.Model(&deletedAuthor).Updates(updateFields).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to restore author"})
		return
	}

	if err := tx.Model(&models.BookAuthor{}).
		Where("author_id = ?", id).
		Updates(map[string]interface{}{
			"is_active":  true,
			"deleted_at": nil,
			"updated_at": time.Now(),
		}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to restore author"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to  restore author"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Author restored successfully"})

}
