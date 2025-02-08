package app

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/GRACENOBLE/ecommerce/api"
	"github.com/GRACENOBLE/ecommerce/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//the function below

func Router() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using defaults.")
	}

	db := database.ConnectDatabase()
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
	
	api.RegisterUserRoutes(r)
	api.RegisterProductRoutes(r)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)

}