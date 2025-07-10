package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Fuzz-Head/pkg/utils"
	"github.com/Fuzz-Head/test"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	pw := "my-secret"
	hashed, _ := utils.HashPassword(pw)

	assert.NotEqual(t, pw, hashed)
	assert.True(t, utils.CheckPassword(pw, hashed))
}

func TestProtectedRoutes(t *testing.T) {
	r := setupTestRouter()

	accessToken, _ := test.GenerateMockAccessToken(1, "admin")

	req, _ := http.NewRequest("GET", "/books", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
