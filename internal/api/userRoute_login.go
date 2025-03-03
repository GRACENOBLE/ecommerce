package api

import (
	"context"
	"net/http"
	"time"

	"github.com/GRACENOBLE/ecommerce/internal/database/queries"
	"github.com/GRACENOBLE/ecommerce/internal/helpers/auth"
	"github.com/GRACENOBLE/ecommerce/internal/types"
	"github.com/gin-gonic/gin"
)

func (dbConfig *DBConfig) Login(c *gin.Context) {
	var loginData types.User

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	var storedUser types.User

	
	err := dbConfig.DB.QueryRow(context.Background(), queries.User.GetUserByEmail, loginData.Email).Scan(&storedUser.ID, &storedUser.Name, &storedUser.Password, &storedUser.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := auth.VerifyPassword(storedUser.Password, loginData.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	accessToken, refreshToken, err := auth.CreateTokenPair(storedUser.ID,storedUser.Email, "user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	c.SetCookie(
		"refresh_token",               // Cookie name
		refreshToken,                  // Cookie value
		int(7*24*time.Hour.Seconds()), // Cookie expiration time in seconds (e.g., 7 days)
		"/",                           // Cookie path
		"",                            // Cookie domain (empty means the current domain)
		true,                          // Secure flag (true ensures the cookie is sent over HTTPS only)
		true,                          // HttpOnly flag (true prevents JavaScript access)
	)

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})

}
