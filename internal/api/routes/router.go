package routes

import (
	"io"
	"log"
	"os"

	"github.com/Fuzz-Head/internal/api/handlers"
	"github.com/Fuzz-Head/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	f, err := os.OpenFile("/var/tmp/book-store.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()

	r.Use(middleware.InjectClaims())

	authLimiter := middleware.NewRateLimiter("3-M")

	// login and register
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.POST("/refresh", authLimiter, handlers.RefreshToken)
	r.POST("/forgot-password", authLimiter, handlers.ForgotPassword)
	r.POST("/reset-password", authLimiter, handlers.ResetPassword)

	r.POST("/logout", handlers.Logout)

	r.GET("/authors", middleware.ScopeRequired("can:read:authors"), handlers.GetAuthors)
	r.POST("/author", middleware.ScopeRequired("can:create:author"), handlers.CreateAuthor)
	r.GET("/author/:id", middleware.ScopeRequired("can:read:author"), handlers.GetAuthor)
	r.PUT("/author/:id", middleware.ScopeRequired("can:update:author"), handlers.UpdateAuthor)
	r.DELETE("/author/:id", middleware.ScopeRequired("can:delete:author"), handlers.DeleteAuthor)
	r.POST("/author/:id/restore", handlers.ResotreAuthor)

	r.GET("/publishers", handlers.GetPublishers)
	r.POST("/publisher", handlers.CreatePublisher)

	// JWT protected routes
	auth := r.Group("/")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.GET("/books", middleware.ScopeRequired("can:read:books"), handlers.GetBooks)
		auth.GET("/book/:id", middleware.ScopeRequired("can:read:book"), handlers.GetBook)
		auth.POST("/book", middleware.ScopeRequired("can:create:book"), handlers.CreateBook)
		auth.PUT("/book/:id", middleware.ScopeRequired("can:update:book"), handlers.UpdateBook)
		auth.DELETE("/book/:id", middleware.ScopeRequired("can:delete:book"), handlers.DeleteBook)
	}

	user := r.Group("/user")
	user.Use(middleware.JWTAuthMiddleware())
	{
		user.GET("/me", handlers.GetCurrentUser)
		user.PUT("/:id", handlers.UpdateUser)
		user.PATCH("/:id/password", handlers.ChangePassword)
		user.DELETE("/:id", handlers.DeleteUser)
	}

	admin := user.Group("/")
	admin.Use(middleware.RequireAdminRole())
	{
		admin.GET("/all", handlers.GetAllUsers)
		admin.PUT("/:id/role", handlers.UpdateUserRole)
	}

	// Protect your entire api from abuse - goes in middleware
	// r.Use(middleware.NewRateLimiter(\"1000-H\")) // 1000 requests per hour per IP

	return r
}
