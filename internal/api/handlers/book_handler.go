package handlers

import (
	"log"
	"net/http"

	"github.com/Fuzz-Head/database"
	"github.com/Fuzz-Head/domain/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	if err := database.DB.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.First(&book, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	// Validate UUID format - and not &models.Book{}
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID format"})
		return
	}

	result := database.DB.Delete(&models.Book{}, "id = ?", id)
	if result.Error != nil {
		log.Println("Binding error:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete a book"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

func CreateBook(c *gin.Context) {
	var input CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newBook := models.Book{
		ID:     uuid.New().String(),
		Title:  input.Title,
		Author: input.Author,
		Price:  input.Price,
		ISBN:   input.ISBN,
	}
	// newBook.Prepare()
	// newBook.ID = uuid.New().String()

	if err := database.DB.Create(&newBook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a book"})
		return
	}

	c.JSON(http.StatusCreated, newBook)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook models.Book

	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		log.Println("Binding error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBook.Prepare()
	updatedBook.ID = id

	var existing models.Book
	if err := database.DB.First(&existing, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := database.DB.Save(&updatedBook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}
	c.JSON(http.StatusOK, updatedBook)
}
