package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nursultan/go-project-nurs/internal/db"
	"github.com/nursultan/go-project-nurs/internal/handler"
	"github.com/nursultan/go-project-nurs/internal/models"
	"github.com/nursultan/go-project-nurs/internal/repository"
	"github.com/nursultan/go-project-nurs/internal/routes"
	"github.com/nursultan/go-project-nurs/internal/service"
)

func main() {
	// Базамен байланыс және миграция
	db.InitDB()

	// GORM арқылы база объектісін алу
	database := db.DB

	// Авто миграция (қаласаң қалдыруға да болады, себебі migrate бар)
	database.AutoMigrate(&models.User{})

	// Репозиторий → Сервис → Хендлер
	bookRepo := repository.NewBookRepository(database)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	// Gin маршруты
	r := gin.Default()
	routes.SetupRoutes(r, bookHandler)

	// Серверді қосу
	r.Run(":8080")
}
