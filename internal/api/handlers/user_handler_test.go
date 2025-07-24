package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/Fuzz-Head/database"
	"github.com/Fuzz-Head/domain/models"
	"github.com/Fuzz-Head/internal/api/middleware"
	"github.com/Fuzz-Head/test"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func setupUserTestRouterWithAuth() *gin.Engine {
	os.Setenv("ENV", "test")
	gin.SetMode(gin.TestMode)
	database.DB = test.SetupTestDB()

	r := gin.Default()

	user := r.Group("/user")
	user.Use(middleware.JWTAuthMiddleware())
	user.GET("/me", GetCurrentUser)
	user.PUT("/:id", UpdateUser)
	user.PATCH("/:id/password", ChangePassword)
	user.DELETE("/:id", DeleteUser)

	admin := user.Group("/")
	admin.Use(middleware.RequireAdminRole())
	admin.GET("/all", GetAllUsers)
	admin.PUT("/:id/role", UpdateUserRole)

	return r
}

func seedTestUser(t *testing.T) models.User {
	username := "testuser_" + strconv.FormatInt(time.Now().UnixNano(), 10)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte("test123"), 10)
	assert.NoError(t, err)

	user := models.User{
		Username: username,
		Email:    username + "@example.com",
		Password: string(passwordHash),
		Role:     "user",
	}

	err = database.DB.Create(&user).Error
	assert.NoError(t, err)

	return user
}

func TestGetCurrentUser(t *testing.T) {
	r := setupUserTestRouterWithAuth()
	user := seedTestUser(t)
	token, _ := test.GenerateMockAccessToken(user.ID, "user")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/me", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestUpdateUser(t *testing.T) {
	r := setupUserTestRouterWithAuth()
	user := seedTestUser(t)
	token, _ := test.GenerateMockAccessToken(user.ID, "user")

	uniqueAppend := strconv.FormatInt(time.Now().UnixNano(), 10)

	input := UpdateUserInput{
		Username: "updatedUsername_" + uniqueAppend,
		Email:    "updated" + uniqueAppend + "@example.com",
	}

	jsonValue, _ := json.Marshal(input)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/user/"+strconv.Itoa(int(user.ID)), bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", "Bearer "+token)

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteuser(t *testing.T) {
	r := setupUserTestRouterWithAuth()
	user := seedTestUser(t)
	token, _ := test.GenerateMockAccessToken(user.ID, "user")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/user/"+strconv.Itoa(int(user.ID)), nil)
	req.Header.Set("Authorization", "Bearer "+token)

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAllUsers_AsAdmin(t *testing.T) {
	r := setupUserTestRouterWithAuth()
	admin := seedTestUser(t)
	admin.Role = "admin"
	database.DB.Save(&admin)
	token, _ := test.GenerateMockAccessToken(admin.ID, "admin")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/all", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateUserRole_AsAdmin(t *testing.T) {
	r := setupUserTestRouterWithAuth()
	admin := seedTestUser(t)
	admin.Role = "admin"
	database.DB.Save(&admin)
	token, _ := test.GenerateMockAccessToken(admin.ID, "admin")

	// later replace this with a struct
	input := UpdateRoleInput{
		Role: "user",
	}

	jsonValue, _ := json.Marshal(input)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/user/"+strconv.Itoa(int(admin.ID))+"/role", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", "Bearer "+token)

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
