package api

import (
	"context"

	"github.com/GRACENOBLE/ecommerce/internal/helpers/auth"
	"github.com/GRACENOBLE/ecommerce/internal/types"
	"github.com/gin-gonic/gin"
)

func (dbConfig *DBConfig) CreateUser(c *gin.Context) {
	var newUser types.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	query := `INSERT INTO users (name, password, email) VALUES ($1, $2, $3) RETURNING id, name, password`

	hashedPassword, err := auth.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := dbConfig.DB.QueryRow(context.Background(), query, newUser.Name, hashedPassword, newUser.Email).Scan(&newUser.ID, &newUser.Name, &newUser.Password); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, newUser)
}
