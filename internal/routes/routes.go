package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nursultan/go-project-nurs/internal/auth"
	"github.com/nursultan/go-project-nurs/internal/handler"
	"github.com/nursultan/go-project-nurs/internal/middleware"
)

func SetupRoutes(r *gin.Engine, bookHandler *handler.BookHandler) {
	// Books
	books := r.Group("/books")
	{
		books.GET("", bookHandler.GetAllBooks)
		books.GET("/:id", bookHandler.GetBook)
		books.POST("", bookHandler.CreateBook)
		books.PUT("/:id", bookHandler.UpdateBook)
		books.DELETE("/:id", bookHandler.DeleteBook)
	}

	// Auth
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", auth.Register)
		authRoutes.POST("/login", auth.Login)
		authRoutes.GET("/me", middleware.AuthRequired(), auth.Me)
	}
}
