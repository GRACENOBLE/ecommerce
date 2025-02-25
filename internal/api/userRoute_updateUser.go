package api

import (
	"context"

	"github.com/GRACENOBLE/ecommerce/internal/types"
	"github.com/gin-gonic/gin"
)

func (dbConfig DBConfig) UpdateUser(c *gin.Context) {
	userId := c.Param("id")

	var user types.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	query := "UPDATE users SET name = $1, password = $2, email = $3 WHERE id = $4"
	_, err := dbConfig.DB.Exec(context.Background(), query, user.Name, user.Password, user.Email, userId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User updated successfully"})
}