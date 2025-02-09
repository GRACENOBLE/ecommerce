package userroutes

import (
	"context"

	"github.com/GRACENOBLE/ecommerce/database"
	"github.com/GRACENOBLE/ecommerce/internal/types"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")
	
	db := database.ConnectDatabase()
	defer db.Close()

	var user types.User
	err := db.QueryRow(context.Background(), "SELECT id, name FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"user": user})
}