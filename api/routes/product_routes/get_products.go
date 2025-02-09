package productroutes

import (
	"context"

	"github.com/GRACENOBLE/ecommerce/database"
	"github.com/GRACENOBLE/ecommerce/internal/types"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	db := database.ConnectDatabase()
	defer db.Close()

	var products []types.Product
	rows, err := db.Query(context.Background(), "SELECT id, name FROM products")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product types.Product
		if err := rows.Scan(&product.Id, &product.Name); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, products)
}