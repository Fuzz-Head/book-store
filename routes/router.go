package routes

import (
	"github.com/Fuzz-Head/handlers"
	"github.com/Fuzz-Head/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.InjectClaims())

	r.GET("/books", middleware.ScopeRequired("can:read:books"), handlers.GetBooks)
	r.GET("/book/:id", middleware.ScopeRequired("can:read:books"), handlers.GetBook)
	r.POST("/book", middleware.ScopeRequired("can:create:books"), handlers.CreateBook)
	r.PUT("/book/:id", middleware.ScopeRequired("can:update:books"), handlers.UpdateBook)
	r.DELETE("/book/:id", middleware.ScopeRequired("can:delete:books"), handlers.DeleteBook)

	return r
}
