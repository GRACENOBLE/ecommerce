package api

import (
	"context"
	"net/http"

	"github.com/GRACENOBLE/ecommerce/internal/database/queries"
	"github.com/GRACENOBLE/ecommerce/internal/helpers/auth"
	"github.com/GRACENOBLE/ecommerce/internal/types"
	"github.com/gin-gonic/gin"
)

func (dbConfig *DBConfig) CreateUser(c *gin.Context) {
	var newUser types.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}



	hashedPassword, err := auth.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := dbConfig.DB.QueryRow(context.Background(), queries.User.CreateUser, newUser.Name, hashedPassword, newUser.Email, newUser.Role).Scan(&newUser.ID, &newUser.Name, &newUser.Role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user":  newUser.Name,
		"enail": newUser.Email,
		"role":  newUser.Role,
	})
}
