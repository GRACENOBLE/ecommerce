package productroutes

import (
	"context"
	"database/sql"

	"github.com/GRACENOBLE/ecommerce/database"
	"github.com/GRACENOBLE/ecommerce/internal/types"
	"github.com/gin-gonic/gin"
)

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Missing product ID"})
		return
	}

	var updatedProduct types.Product

	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := database.ConnectDatabase()
	defer db.Close()

	query := `UPDATE products SET name = COALESCE($1, name), description = COALESCE($2, description), price = COALESCE($3, price), stock_quantity = COALESCE($4, stock_quantity), image_url = COALESCE($5, image_url), updated_at = CURRENT_TIMESTAMP WHERE id = $6 RETURNING id, name, description, price, stock_quantity, image_url, updated_at;`
	
	err := db.QueryRow(context.Background(), query, updatedProduct.Name, updatedProduct.Description, updatedProduct.Price, updatedProduct.StockQuantity, updatedProduct.ImageURL, id).Scan(&updatedProduct.Id, &updatedProduct.Name, &updatedProduct.Description, &updatedProduct.Price, &updatedProduct.StockQuantity, &updatedProduct.ImageURL, &updatedProduct.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"error": "Product not found"})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(200, updatedProduct)
}