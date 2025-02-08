package app

import (
	"log"
	"os"

	"github.com/GRACENOBLE/ecommerce/api"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Router() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using defaults.")
	}

	r := gin.Default()
	
	api.RegisterUserRoutes(r)
	api.RegisterProductRoutes(r)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)

}