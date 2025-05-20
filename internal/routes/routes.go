package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nursultan/go-project-nurs/internal/auth"
	"github.com/nursultan/go-project-nurs/internal/handler"
	"github.com/nursultan/go-project-nurs/internal/middleware"
	"github.com/nursultan/go-project-nurs/internal/repository"
	"github.com/nursultan/go-project-nurs/internal/service"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	// Books (қорғалған)
	books := r.Group("/books", middleware.AuthRequired())
	{
		// Барлық қолданушылар көре алады
		books.GET("", bookHandler.GetAllBooks)
		books.GET("/:id", bookHandler.GetBook)

		// Тек admin рөліндегі қолданушы ғана жасай алады
		books.POST("", middleware.AdminOnly(), bookHandler.CreateBook)
		books.PUT("/:id", middleware.AdminOnly(), bookHandler.UpdateBook)
		books.DELETE("/:id", middleware.AdminOnly(), bookHandler.DeleteBook)
	}

	// Auth
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", auth.Register)
		authRoutes.POST("/login", auth.Login)
		authRoutes.GET("/me", middleware.AuthRequired(), auth.Me)
	}
}
