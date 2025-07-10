package routes

import (
	"github.com/Fuzz-Head/internal/api/handlers"
	"github.com/Fuzz-Head/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.InjectClaims())

	// r.GET("/books", middleware.ScopeRequired("can:read:books"), handlers.GetBooks)
	// r.GET("/book/:id", middleware.ScopeRequired("can:read:books"), handlers.GetBook)
	// r.POST("/book", middleware.ScopeRequired("can:create:books"), handlers.CreateBook)
	// r.PUT("/book/:id", middleware.ScopeRequired("can:update:books"), handlers.UpdateBook)
	// r.DELETE("/book/:id", middleware.ScopeRequired("can:delete:books"), handlers.DeleteBook)

	// login and register
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	r.POST("/logout", handlers.Logout)

	// JWT protected routes
	auth := r.Group("/")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.GET("/books", middleware.ScopeRequired("can:read:books"), handlers.GetBooks)
		auth.GET("/book/:id", middleware.ScopeRequired("can:read:books"), handlers.GetBook)
		auth.POST("/book", middleware.ScopeRequired("can:create:books"), handlers.CreateBook)
		auth.PUT("/book/:id", middleware.ScopeRequired("can:update:books"), handlers.UpdateBook)
		auth.DELETE("/book/:id", middleware.ScopeRequired("can:delete:books"), handlers.DeleteBook)
	}

	return r
}
