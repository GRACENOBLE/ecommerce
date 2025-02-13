package api

import (
	"context"

	"github.com/GRACENOBLE/ecommerce/internal/types"
	"github.com/gin-gonic/gin"
)

func (dbConfig *DBConfig) GetProductById(c *gin.Context) {
	id := c.Param("id")


	var product types.Product

	err := dbConfig.DB.QueryRow(context.Background(), "SELECT id, name, description, price, stock_quantity, image_url, created_at, updated_at FROM products WHERE id = $1", id).Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.StockQuantity, &product.ImageURL, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(200, product)
	
}