package auth

import (
	"errors"
	"log"
	"os"

	"github.com/GRACENOBLE/ecommerce/internal/types"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func RefreshAccessToken(refreshToken string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using defaults.")
	}

	jwtSecretKey := []byte(os.Getenv("JWT_SECRET"))

	// Parse and validate the refresh token
	token, err := jwt.ParseWithClaims(refreshToken, &types.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if err != nil || !token.Valid {
		return "", errors.New("invalid refresh token")
	}

	// Extract claims from the token
	claims, ok := token.Claims.(*types.Claims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	// Generate a new access token
	accessToken, err := CreateAccessToken(claims.ID, claims.Email, claims.Role)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
