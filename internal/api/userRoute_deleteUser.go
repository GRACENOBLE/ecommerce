package api

import (
	"context"
	"net/http"

	"github.com/GRACENOBLE/ecommerce/internal/database/queries"
	"github.com/GRACENOBLE/ecommerce/internal/types"
	"github.com/gin-gonic/gin"
)

func (dbConfig DBConfig) DeleteUser(c *gin.Context) {
	userId := c.Param("id")

	var user types.User

	if err := dbConfig.DB.QueryRow(context.Background(), queries.User.CheckIfUserExists, userId).Scan(&user.Name); err != nil {
		if err.Error() == "no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check if user exists"})
		}
		return
	}

	_, err := dbConfig.DB.Exec(context.Background(), queries.User.DeleteUser, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted user",
	})
}
