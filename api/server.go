package api

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/GRACENOBLE/ecommerce/internal/helpers"
	"github.com/GRACENOBLE/ecommerce/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Router() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using defaults.")
	}

	db := helpers.ConnectDatabase()
	defer db.Close()

	rows, err := db.Query(context.Background(), "SELECT id, name FROM users")
	if err != nil {
		log.Fatalf("Failed to query data: %v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatalf("Failed to scan row: %v\n", err)
		}
		fmt.Printf("ID: %s, Name: %s\n", id, name)
	}

	r := gin.Default()
	
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	//  This function call registers the user routes with the router
	routes.RegisterUserRoutes(r)
	routes.RegisterProductRoutes(r)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)

}
