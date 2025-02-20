package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// CreateTokenPair generates an access and a refresh token for a given user email and role.
func CreateTokenPair(email, role string) (accessToken string, refreshToken string, err error) {
	err = godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using defaults.")
	}

	jwtSecretKey := []byte(os.Getenv("JWT_SECRET"))

	// Create Access Token
	accessTokenClaims := jwt.MapClaims{
		"authorized": true,
		"email":      email,
		"role":       role,
		"exp":        time.Now().Add(15 * time.Minute).Unix(),
	}
	accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessToken, err = accessTokenObj.SignedString(jwtSecretKey)
	if err != nil {
		return "", "", err
	}

	// Create Refresh Token
	refreshTokenClaims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(7 * 24 * time.Hour).Unix(), // 7 days
	}
	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshToken, err = refreshTokenObj.SignedString(jwtSecretKey)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
