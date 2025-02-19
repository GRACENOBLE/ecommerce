package api

import (
	"context"
	"net/http"

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

	query := "SELECT id, name, password, email FROM users WHERE email = $1"
	err := dbConfig.DB.QueryRow(context.Background(), query, loginData.Email).Scan(&storedUser.ID, &storedUser.Name, &storedUser.Password, &storedUser.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := auth.VerifyPassword(storedUser.Password, loginData.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := auth.CreateToken(storedUser.Email, "user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
