package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (dbConfig *DBConfig) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing product ID"})
		return
	}

	query := `DELETE FROM products WHERE id = $1;`

	result, err := dbConfig.DB.Exec(context.Background(), query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"rows affected": rowsAffected,
		"deleted":       id,
	})

}
