package userroutes

import (
	"context"

	"github.com/GRACENOBLE/ecommerce/database"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	name, password, email := c.Param("name"), c.Param("password"), c.Param("email")
	db := database.ConnectDatabase()
	defer db.Close()
	result, err := db.Exec(context.Background(), "INSERT INTO users (name, password, email) VALUES($1,$2,$3)", name, password, email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"result": result})
}
