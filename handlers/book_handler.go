package handlers

import (
	"net/http"

	"github.com/Fuzz-Head/models"
	"github.com/gin-gonic/gin"
)

var books = []models.Book{
	{ID: "1", Title: "1984", Author: "George Orwell", Price: 9.99, ISBN: "9780452284234"},
	{ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee", Price: 10.99, ISBN: "9780061120084"},
	{ID: "3", Title: "PrIDe and Prejudice", Author: "Jane Austen", Price: 8.95, ISBN: "9780141439518"},
	{ID: "4", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Price: 9.49, ISBN: "9780743273565"},
	{ID: "5", Title: "The Catcher in the Rye", Author: "J.D. Salinger", Price: 9.79, ISBN: "9780316769488"},
	{ID: "6", Title: "The Hobbit", Author: "J.R.R. Tolkien", Price: 12.99, ISBN: "9780547928227"},
	{ID: "7", Title: "Fahrenheit 451", Author: "Ray Bradbury", Price: 9.89, ISBN: "9781451673319"},
	{ID: "8", Title: "Jane Eyre", Author: "Charlotte Brontë", Price: 11.50, ISBN: "9780141441146"},
	{ID: "9", Title: "Brave New World", Author: "Aldous Huxley", Price: 10.25, ISBN: "9780060850524"},
	{ID: "10", Title: "The Alchemist", Author: "Paulo Coelho", Price: 10.99, ISBN: "9780061122415"},
	{ID: "11", Title: "The Book Thief", Author: "Markus Zusak", Price: 11.95, ISBN: "9780375842207"},
	{ID: "12", Title: "The Road", Author: "Cormac McCarthy", Price: 9.50, ISBN: "9780307387899"},
	{ID: "13", Title: "Sapiens", Author: "Yuval Noah Harari", Price: 14.99, ISBN: "9780062316097"},
	{ID: "14", Title: "Educated", Author: "Tara Westover", Price: 13.95, ISBN: "9780399590504"},
	{ID: "15", Title: "The Silent Patient", Author: "Alex MichaelIDes", Price: 12.00, ISBN: "9781250301697"},
	{ID: "16", Title: "Where the Crawdads Sing", Author: "Delia Owens", Price: 13.45, ISBN: "9780735219106"},
	{ID: "17", Title: "Atomic Habits", Author: "James Clear", Price: 11.79, ISBN: "9780735211292"},
	{ID: "18", Title: "The MIDnight Library", Author: "Matt Haig", Price: 12.49, ISBN: "9780525559474"},
	{ID: "19", Title: "The Subtle Art of Not Giving a F*ck", Author: "Mark Manson", Price: 10.59, ISBN: "9780062457714"},
	{ID: "20", Title: "Becoming", Author: "Michelle Obama", Price: 15.99, ISBN: "9781524763138"},
}

func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	for _, b := range books {
		if b.ID == id {
			c.JSON(http.StatusOK, b)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func CreateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newBook.Prepare()
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook models.Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedBook.Prepare()
	for i, b := range books {
		if b.ID == id {
			books[i] = updatedBook
			c.JSON(http.StatusOK, updatedBook)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
