package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Fuzz-Head/database"
	"github.com/Fuzz-Head/internal/api/middleware"
	"github.com/Fuzz-Head/test"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var mockID = uuid.New().String()

func setupTestRouter() *gin.Engine {
	os.Setenv("ENV", "test")
	gin.SetMode(gin.TestMode)
	database.DB = test.SetupTestDB()

	r := gin.Default()

	// inject mock claims
	// r.Use(test.MockAuthMiddleware())

	r.GET("/books", middleware.JWTAuthMiddleware(), GetBooks)
	r.GET("/book/:id", middleware.JWTAuthMiddleware(), GetBook)
	r.POST("/book", middleware.JWTAuthMiddleware(), CreateBook)
	r.PUT("/book/:id", middleware.JWTAuthMiddleware(), UpdateBook)
	r.DELETE("/book/:id", middleware.JWTAuthMiddleware(), DeleteBook)

	return r
}

func TestGetBooks_Unauthorized(t *testing.T) {
	// router := gin.Default()
	// router.GET("/books", GetBook)
	r := setupTestRouter()

	req, _ := http.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestCreateBook(t *testing.T) {
	r := setupTestRouter()
	token, _ := test.GenerateMockAccessToken(1, "admin")

	input := CreateBookInput{
		Title:  "Test book",
		Author: "Test Author",
		Price:  12.34,
		ISBN:   "9780140449136",
	}

	jsonValue, _ := json.Marshal(input)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/book", bytes.NewBuffer(jsonValue))

	req.Header.Set("Authorization", "Bearer "+token)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetBooks(t *testing.T) {
	r := setupTestRouter()
	token, _ := test.GenerateMockAccessToken(1, "admin")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetBook_NotFound(t *testing.T) {
	r := setupTestRouter()
	token, _ := test.GenerateMockAccessToken(1, "admin")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/book/"+mockID, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUpdateBook_NotFound(t *testing.T) {
	r := setupTestRouter()
	token, _ := test.GenerateMockAccessToken(1, "admin")

	input := UpdateBookInput{
		Title:  "Updated Book",
		Author: "updated Author",
		Price:  15.00,
		ISBN:   "9780140449136",
	}
	jsonValue, _ := json.Marshal(input)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/book/"+mockID, bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", "Bearer "+token)

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteBook_NotFound(t *testing.T) {
	r := setupTestRouter()
	token, _ := test.GenerateMockAccessToken(1, "admin")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/book/"+mockID, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)

}
