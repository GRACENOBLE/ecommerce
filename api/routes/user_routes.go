package routes

import (
	"context"

	"github.com/GRACENOBLE/ecommerce/internal/helpers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	r.GET("/user/:id", getUser)
	r.DELETE("/user/:id", deleteUser)
	r.POST("/user/:name/:passsword/:email", createUser)
	r.PUT("/user/:id", updateUser)
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	db := helpers.ConnectDatabase()
	defer db.Close()

	var user struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	err := db.QueryRow(context.Background(), "SELECT id, name FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

func deleteUser(c *gin.Context) {
	// Handler logic for deleting a user
}

func createUser(c *gin.Context) {
	name, password, email := c.Param("name"), c.Param("password"), c.Param("email")
	db := helpers.ConnectDatabase()
	defer db.Close()
	result, err := db.Exec(context.Background(), "INSERT INTO users (name, password, email) VALUES($1,$2,$3)", name, password, email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"result": result})
}

func updateUser(c *gin.Context) {
	// Handler logic for updating a user
}
