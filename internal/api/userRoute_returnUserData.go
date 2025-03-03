package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/GRACENOBLE/ecommerce/internal/database/queries"
	"github.com/GRACENOBLE/ecommerce/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (dbConfig DBConfig) ReturnUserData(c *gin.Context) {
	claims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	userID, ok := userClaims["sub"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

	fmt.Printf("user id: %v", userID)
	var user types.User

	err := dbConfig.DB.QueryRow(context.Background(), queries.User.GetUserByID, userID).Scan(&user.Name, &user.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)

}
