package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectDatabase() *pgxpool.Pool {
	//Load the database URL from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using defaults.")
	}

	databaseURL := os.Getenv("DATABASE_URL")

	// Create a connection pool
	pool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Verify the connection
	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	fmt.Println("Connected to the database successfully!")
	return pool
}
