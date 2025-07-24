package routes

import (
	"github.com/Fuzz-Head/internal/api/handlers"
	"github.com/Fuzz-Head/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.InjectClaims())

	authLimiter := middleware.NewRateLimiter("3-M")

	// r.GET("/books", middleware.ScopeRequired("can:read:books"), handlers.GetBooks)

	// login and register
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.POST("/refresh", authLimiter, handlers.RefreshToken)
	r.POST("/forgot-password", authLimiter, handlers.ForgotPassword)
	r.POST("/reset-password", authLimiter, handlers.ResetPassword)

	r.POST("/logout", handlers.Logout)

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
