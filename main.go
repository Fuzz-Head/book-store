package main

import (
	"log"

	"github.com/Fuzz-Head/database"
	"github.com/Fuzz-Head/models"
	"github.com/Fuzz-Head/routes"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {

	database.Connect()

	// Register custom ISBN validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("isbn", models.IsbnValidator)
	}

	r := routes.SetupRouter()

	log.Println("Bookstore api server is starting on localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
