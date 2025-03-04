package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// CreateToken generates a JWT for a given user email and role.
func CreateAccessToken(id, email, role string) (string, error) {

	log.Printf("ID: %v, EMAIL: %v, ROLE: %v", id, email, role)

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using defaults.")
	}

	jwtSecretKey := []byte(os.Getenv("JWT_SECRET"))

	claims := jwt.MapClaims{
		"authorized": true,
		"sub":        id,
		"email":      email,
		"role":       role,
		"exp":        time.Now().Add(15 * time.Minute).Unix(),
	}

	accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessTokenObj.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
