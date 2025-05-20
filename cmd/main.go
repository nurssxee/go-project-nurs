package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nursultan/go-project-nurs/internal/db"
	"github.com/nursultan/go-project-nurs/internal/routes"
)

func main() {
	// Базамен байланыс және миграцияны іске қосу
	db.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r, db.DB)

	r.Run(":8080")
}
