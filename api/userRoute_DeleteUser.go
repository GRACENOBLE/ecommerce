package api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/GRACENOBLE/ecommerce/internal/types"
	"github.com/gin-gonic/gin"
)

func (dbConfig DBConfig) DeleteUser(c *gin.Context) {
	userId := c.Param("id")

	var user types.User

	checkIfUserExists := "SELECT name FROM users where id = $1"
	if err := dbConfig.DB.QueryRow(context.Background(), checkIfUserExists, userId).Scan(&user.Name); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check if user exists"})
		}
		return
	}

	query := "DELETE FROM users where id = $1"
	_, err := dbConfig.DB.Exec(context.Background(), query, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted user",
	})
}
