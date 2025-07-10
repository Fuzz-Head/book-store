package handlers

import (
	"os"
	"testing"

	"github.com/Fuzz-Head/database"
	"github.com/Fuzz-Head/domain/models"
	// "github.com/Fuzz-Head/pkg/utils"
	"github.com/Fuzz-Head/test"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func TestMain(m *testing.M) {
	// Set test environment
	os.Setenv("ENV", "test")

	// Initialize fresh in-memory test DB
	database.DB = test.SetupTestDB()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("isbn", models.IsbnValidator) // add the validator function to utils
	}

	// Run all tests
	code := m.Run()

	// Exit with code
	os.Exit(code)
}
