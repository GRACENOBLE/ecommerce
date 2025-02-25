package api

import (
	"net/http"

	"github.com/GRACENOBLE/ecommerce/internal/helpers/auth"
	"github.com/gin-gonic/gin"
)

func (dbConfig *DBConfig) RefreshToken(c *gin.Context) {
    // Read the refresh token from the HTTP-only cookie
    refreshToken, err := c.Cookie("refresh_token")
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
        return
    }

    // Validate the refresh token and generate a new access token
    accessToken, err := auth.RefreshAccessToken(refreshToken)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
        return
    }

    // Return the new access token
    c.JSON(http.StatusOK, gin.H{
        "access_token": accessToken,
    })
}
