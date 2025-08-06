package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Fuzz-Head/database"
	"github.com/Fuzz-Head/domain/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetBooks retrieves all books with their relationships
func GetBooks(c *gin.Context) {
	var books []models.Book

	// Parse query parameters for pagination and filtering
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

	// Build query with preloads for relationships
	query := database.DB.
		Preload("Publisher").
		Preload("Series").
		Preload("Categories").
		Preload("Authors").
		Preload("Authors.Author").
		Preload("Editions").
		Preload("Awards").
		Preload("Awards.Award").
		Limit(limit).
		Offset(offset)

	// Add filtering options
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	if featured := c.Query("featured"); featured == "true" {
		query = query.Where("featured = ?", true)
	}

	if bestseller := c.Query("bestseller"); bestseller == "true" {
		query = query.Where("bestseller = ?", true)
	}

	if err := query.Find(&books).Error; err != nil {
		log.Printf("Error fetching books: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"books":  books,
		"limit":  limit,
		"offset": offset,
		"count":  len(books),
	})
}

// GetBook retrieves a single book with all its relationships
func GetBook(c *gin.Context) {
	id := c.Param("id")

	// Validate UUID format
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID format"})
		return
	}

	var book models.Book

	err := database.DB.
		Preload("Publisher").
		Preload("Series").
		Preload("Categories").
		Preload("Authors").
		Preload("Authors.Author").
		Preload("Editions").
		Preload("Awards").
		Preload("Awards.Award").
		Preload("Reviews").
		First(&book, "id = ?", id).Error

	if err != nil {
		log.Printf("Error fetching book %s: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Increment view count
	database.DB.Model(&book).Update("view_count", book.ViewCount+1)

	c.JSON(http.StatusOK, book)
}

// CreateBook creates a new book with relationships
func CreateBook(c *gin.Context) {
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Creating book with input: %+v", input)

	// Generate slug if not provided
	if input.Slug == "" {
		input.Slug = generateSlug(input.Title)
	}

	// Create the book
	newBook := models.Book{
		ID:                  uuid.New().String(),
		Title:               input.Title,
		Subtitle:            input.Subtitle,
		OriginalTitle:       input.OriginalTitle,
		Description:         input.Description,
		ShortDescription:    input.ShortDescription,
		PublisherID:         input.PublisherID,
		PublicationDate:     input.PublicationDate,
		OriginalPublishDate: input.OriginalPublishDate,
		Edition:             input.Edition,
		SeriesID:            input.SeriesID,
		SeriesNumber:        input.SeriesNumber,
		Language:            input.Language,
		OriginalLanguage:    input.OriginalLanguage,
		TranslatedFrom:      input.TranslatedFrom,
		AgeRating:           input.AgeRating,
		ContentWarnings:     input.ContentWarnings,
		Keywords:            input.Keywords,
		Tags:                input.Tags,
		Genres:              input.Genres,
		Subjects:            input.Subjects,
		DeweyDecimal:        input.DeweyDecimal,
		LibraryOfCongress:   input.LibraryOfCongress,
		Status:              getStatusOrDefault(input.Status),
		Featured:            input.Featured,
		Bestseller:          input.Bestseller,
		NewRelease:          input.NewRelease,
		StaffPick:           input.StaffPick,
		SEOTitle:            input.SEOTitle,
		SEODescription:      input.SEODescription,
		MetaKeywords:        input.MetaKeywords,
		Slug:                input.Slug,
		ExternalIDs:         input.ExternalIDs,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	// Start transaction
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create the book
	if err := tx.Create(&newBook).Error; err != nil {
		tx.Rollback()
		log.Printf("Error creating book: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	// Associate categories
	if len(input.CategoryIDs) > 0 {
		var categories []models.Category
		if err := tx.Where("id IN ?", input.CategoryIDs).Find(&categories).Error; err != nil {
			tx.Rollback()
			log.Printf("Error finding categories: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category IDs"})
			return
		}
		if err := tx.Model(&newBook).Association("Categories").Replace(categories); err != nil {
			tx.Rollback()
			log.Printf("Error associating categories: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to associate categories"})
			return
		}
	}

	// Associate authors
	if len(input.AuthorIDs) > 0 {
		for _, authorID := range input.AuthorIDs {
			bookAuthor := models.BookAuthor{
				BookID:   newBook.ID,
				AuthorID: authorID,
				Role:     "author",
			}
			if err := tx.Create(&bookAuthor).Error; err != nil {
				tx.Rollback()
				log.Printf("Error associating author %s: %v", authorID, err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to associate authors"})
				return
			}
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		log.Printf("Error committing transaction: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	// Fetch the created book with relationships
	var createdBook models.Book
	database.DB.
		Preload("Publisher").
		Preload("Categories").
		Preload("Authors.Author").
		First(&createdBook, "id = ?", newBook.ID)

	c.JSON(http.StatusCreated, createdBook)
}

// UpdateBook updates an existing book
func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	// Validate UUID format
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID format"})
		return
	}

	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Binding error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if book exists
	var existingBook models.Book
	if err := database.DB.First(&existingBook, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Start transaction
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Update book fields
	updates := UpdateBookInput{
		Title:               input.Title,
		Subtitle:            input.Subtitle,
		OriginalTitle:       input.OriginalTitle,
		Description:         input.Description,
		ShortDescription:    input.ShortDescription,
		PublisherID:         input.PublisherID,
		PublicationDate:     input.PublicationDate,
		OriginalPublishDate: input.OriginalPublishDate,
		Edition:             input.Edition,
		SeriesID:            input.SeriesID,
		SeriesNumber:        input.SeriesNumber,
		Language:            input.Language,
		OriginalLanguage:    input.OriginalLanguage,
		TranslatedFrom:      input.TranslatedFrom,
		AgeRating:           input.AgeRating,
		ContentWarnings:     input.ContentWarnings,
		Keywords:            input.Keywords,
		Tags:                input.Tags,
		Genres:              input.Genres,
		Subjects:            input.Subjects,
		DeweyDecimal:        input.DeweyDecimal,
		LibraryOfCongress:   input.LibraryOfCongress,
		Status:              getStatusOrDefault(input.Status),
		Featured:            input.Featured,
		Bestseller:          input.Bestseller,
		NewRelease:          input.NewRelease,
		StaffPick:           input.StaffPick,
		SEOTitle:            input.SEOTitle,
		SEODescription:      input.SEODescription,
		MetaKeywords:        input.MetaKeywords,
		Slug:                input.Slug,
		ExternalIDs:         input.ExternalIDs,
	}

	if err := tx.Model(&existingBook).Updates(updates).Error; err != nil {
		tx.Rollback()
		log.Printf("Error updating book: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	// Update categories if provided
	if len(input.CategoryIDs) > 0 {
		var categories []models.Category
		if err := tx.Where("id IN ?", input.CategoryIDs).Find(&categories).Error; err != nil {
			tx.Rollback()
			log.Printf("Error finding categories: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category IDs"})
			return
		}
		if err := tx.Model(&existingBook).Association("Categories").Replace(categories); err != nil {
			tx.Rollback()
			log.Printf("Error updating categories: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update categories"})
			return
		}
	}

	// Update authors if provided
	if len(input.AuthorIDs) > 0 {
		// Remove existing author associations
		if err := tx.Where("book_id = ?", id).Delete(&models.BookAuthor{}).Error; err != nil {
			tx.Rollback()
			log.Printf("Error removing existing authors: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update authors"})
			return
		}

		// Add new author associations
		for _, authorID := range input.AuthorIDs {
			bookAuthor := models.BookAuthor{
				BookID:   id,
				AuthorID: authorID,
				Role:     "author",
			}
			if err := tx.Create(&bookAuthor).Error; err != nil {
				tx.Rollback()
				log.Printf("Error associating author %s: %v", authorID, err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update authors"})
				return
			}
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		log.Printf("Error committing update transaction: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	// Fetch updated book with relationships
	var updatedBook models.Book
	database.DB.
		Preload("Publisher").
		Preload("Categories").
		Preload("Authors.Author").
		First(&updatedBook, "id = ?", id)

	c.JSON(http.StatusOK, updatedBook)
}

// DeleteBook deletes a book and its relationships
func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	// Validate UUID format
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID format"})
		return
	}

	// Start transaction
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Delete related records first
	if err := tx.Where("book_id = ?", id).Delete(&models.BookAuthor{}).Error; err != nil {
		tx.Rollback()
		log.Printf("Error deleting book authors: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	if err := tx.Where("book_id = ?", id).Delete(&models.BookAward{}).Error; err != nil {
		tx.Rollback()
		log.Printf("Error deleting book awards: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	if err := tx.Where("book_id = ?", id).Delete(&models.Edition{}).Error; err != nil {
		tx.Rollback()
		log.Printf("Error deleting book editions: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	if err := tx.Where("book_id = ?", id).Delete(&models.Review{}).Error; err != nil {
		tx.Rollback()
		log.Printf("Error deleting book reviews: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	// Delete the book itself
	result := tx.Delete(&models.Book{}, "id = ?", id)
	if result.Error != nil {
		tx.Rollback()
		log.Printf("Error deleting book: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	if result.RowsAffected == 0 {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		log.Printf("Error committing delete transaction: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

// Helper functions
func generateSlug(title string) string {
	// Simple slug generation - you might want to use a more sophisticated library
	// This is a basic implementation
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")
	// Remove special characters (basic implementation)
	reg := regexp.MustCompile("[^a-z0-9-]")
	slug = reg.ReplaceAllString(slug, "")
	return slug
}

func getStatusOrDefault(status string) string {
	if status == "" {
		return "published"
	}
	return status
}
