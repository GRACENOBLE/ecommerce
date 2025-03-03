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

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using defaults.")
	}

	jwtSecretKey := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["sub"] = id
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(15 * time.Minute).Unix()

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}