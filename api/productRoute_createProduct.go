package api

import (
	"context"
	"net/http"

	"github.com/GRACENOBLE/ecommerce/internal/types"
	"github.com/gin-gonic/gin"
)

func (dbConfig *DBConfig) CreateProduct(c *gin.Context) {
	var newProduct types.Product

	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO products (name, description, price, stock_quantity, image_url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) RETURNING id`
	
	err := dbConfig.DB.QueryRow(context.Background(), query, newProduct.Name, newProduct.Description, newProduct.Price, newProduct.StockQuantity, newProduct.ImageURL).Scan(&newProduct.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newProduct)
}
