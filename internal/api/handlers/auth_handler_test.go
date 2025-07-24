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
	"github.com/Fuzz-Head/pkg/utils"
	"github.com/Fuzz-Head/test"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type AccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token" binding:"required"`
}

type TokenRequest struct {
	AccessToken  string `json:"access_token" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func setupAuthTestRouter() *gin.Engine {
	os.Setenv("ENV", "test")
	gin.SetMode(gin.TestMode)
	database.DB = test.SetupTestDB()

	r := gin.Default()

	r.POST("/register", Register)
	//r.POST("/login", Login)
	r.POST("/refresh", RefreshToken)
	//r.POST("/forgot-password", ForgotPassword)
	//r.POST("/reset-password", ResetPassword)
	//r.POST("/logout", Logout)

	return r
}

func seedTestUser_NotToDB(t *testing.T) models.User {
	username := "testuser_" + strconv.FormatInt(time.Now().UnixNano(), 10)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte("test123"), 10)
	assert.NoError(t, err)

	user := models.User{
		Username: username,
		Email:    username + "@example.com",
		Password: string(passwordHash),
		Role:     "user",
	}

	return user
}
func TestHashPassword(t *testing.T) {
	pw := "my-secret"
	hashed, _ := utils.HashPassword(pw)

	assert.NotEqual(t, pw, hashed)
	assert.True(t, utils.CheckPassword(pw, hashed))
}

func TestProtectedRoutes(t *testing.T) {
	r := setupTestRouter()

	accessToken, _ := test.GenerateMockAccessToken(7, "admin")

	req, _ := http.NewRequest("GET", "/books", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRefreshAccessToken(t *testing.T) {
	r := setupAuthTestRouter()
	user := seedTestUser(t)

	refreshToken, err := test.GenerateMockRefreshToken(user.ID, user.Role)
	assert.NoError(t, err)

	user.RefreshToken = refreshToken
	database.DB.Save(&user)

	// Simulate login request with correct credentials
	reqBody := AccessTokenRequest{
		RefreshToken: user.RefreshToken,
	}

	jsonValue, _ := json.Marshal(reqBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/refresh", bytes.NewBuffer(jsonValue))
	//req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	// Assert success and new access token
	assert.Equal(t, http.StatusOK, w.Code)

	var resp AccessTokenResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.AccessToken)
}

func TestRegister(t *testing.T) {
	r := setupAuthTestRouter()
	user := seedTestUser_NotToDB(t)

	jsonValue, _ := json.Marshal(user)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
