package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nursultan/go-project-nurs/internal/handler"
)

func SetupRoutes(r *gin.Engine, bookHandler *handler.BookHandler) {
	books := r.Group("/books")
	{
		books.GET("", bookHandler.GetAllBooks)
		books.GET("/:id", bookHandler.GetBook)
		books.POST("", bookHandler.CreateBook)
		books.PUT("/:id", bookHandler.UpdateBook)
		books.DELETE("/:id", bookHandler.DeleteBook)
	}
}
