package app

import (
	"log"
	"os"

	"github.com/GRACENOBLE/ecommerce/internal/api"
	"github.com/GRACENOBLE/ecommerce/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Router() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using defaults.")
	}

	db := database.ConnectDatabase()

	defer func() {
		log.Println("Closing database connection...")
		db.Close()
	}()

	r := gin.Default()

	dbConfig := &api.DBConfig{DB: db}

	api.RegisterRoutes(r, dbConfig)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)

}
