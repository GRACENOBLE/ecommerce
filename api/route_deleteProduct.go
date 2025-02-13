package api

import (
	"context"

	"github.com/gin-gonic/gin"
)

func (dbConfig *DBConfig) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Missing product ID"})
		return
	}

	query := `DELETE FROM products WHERE id = $1;`

	result, err := dbConfig.DB.Exec(context.Background(), query, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(200, gin.H{
		"rows affected": rowsAffected,
		"deleted":       id,
	})

}
