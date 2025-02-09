package productroutes

import (
	"context"

	"github.com/GRACENOBLE/ecommerce/database"
	"github.com/GRACENOBLE/ecommerce/internal/types"
	"github.com/gin-gonic/gin"
)

func GetProductById(c *gin.Context) {
	id := c.Param("id")

	db := database.ConnectDatabase()
	defer db.Close()

	var product types.Product

	err := db.QueryRow(context.Background(), "SELECT id, name, description, price, stock_quantity, image_url, created_at, updated_at FROM products WHERE id = $1", id).Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.StockQuantity, &product.ImageURL, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, product)

}
