package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// GenerateJWT generates a JWT token for a user
func GenerateJWT(userID string) (string, error) {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using defaults.")
	}
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token valid for 24 hours
	})
	return token.SignedString(jwtSecret)

}
