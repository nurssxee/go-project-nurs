package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nursultan/go-project-nurs/internal/handler"
	"github.com/nursultan/go-project-nurs/internal/models"
	"github.com/nursultan/go-project-nurs/internal/repository"
	"github.com/nursultan/go-project-nurs/internal/routes"
	"github.com/nursultan/go-project-nurs/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "postgres://nurs:4469@localhost:5432/postgres?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatal("Error on migrating to the DB:", err)
	}

	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	r := gin.Default()
	routes.SetupRoutes(r, bookHandler)

	r.Run(":8080")
}
