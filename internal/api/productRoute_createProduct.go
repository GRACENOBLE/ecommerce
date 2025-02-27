package api

import (
	"context"
	"net/http"

	"github.com/GRACENOBLE/ecommerce/internal/database/queries"
	"github.com/GRACENOBLE/ecommerce/internal/types"
	"github.com/gin-gonic/gin"
)

func (dbConfig *DBConfig) CreateProduct(c *gin.Context) {
	var newProduct types.Product

	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	err := dbConfig.DB.QueryRow(context.Background(), queries.Product.CreateProduct, newProduct.Name, newProduct.Description, newProduct.Price, newProduct.StockQuantity, newProduct.ImageURL).Scan(&newProduct.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newProduct)
}
