package api

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Router() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using defaults.")
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)

}
