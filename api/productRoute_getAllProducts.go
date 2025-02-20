package api

import (
	"context"
	"net/http"

	"github.com/GRACENOBLE/ecommerce/internal/types"
	"github.com/gin-gonic/gin"
)

func (dbConfig *DBConfig) GetProducts(c *gin.Context) {

	var products []types.Product
	rows, err := dbConfig.DB.Query(context.Background(), "SELECT id, name, description, price, stock_quantity, image_url, created_at, updated_at FROM products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product types.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.StockQuantity, &product.ImageURL, &product.CreatedAt, &product.UpdatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
