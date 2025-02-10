package productroutes

import (
	"context"
	"fmt"

	"github.com/GRACENOBLE/ecommerce/database"
	"github.com/GRACENOBLE/ecommerce/internal/types"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var newProduct types.Product

	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("%v",newProduct)

	db := database.ConnectDatabase()
	defer db.Close()

	query := `INSERT INTO products (name, description, price, stock_quantity, image_url, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) RETURNING id`
	err := db.QueryRow(context.Background(), query, newProduct.Name, newProduct.Description, newProduct.Price, newProduct.StockQuantity, newProduct.ImageURL).Scan(&newProduct.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, newProduct)
}