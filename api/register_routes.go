package api

import (
	"context"
	"database/sql"

	userroutes "github.com/GRACENOBLE/ecommerce/api/routes/user_routes"
	"github.com/GRACENOBLE/ecommerce/database"
	"github.com/GRACENOBLE/ecommerce/internal/types"
	"github.com/gin-gonic/gin"
)

func  RegisterProductRoutes(r *gin.Engine) {
	db := &DBConfig{}

	r.GET("/products", db.GetProducts)
	r.GET("/product/:id", db.GetProductById)
	r.POST("/products", db.CreateProduct)
	r.PUT("/products/:id", db.UpdateProduct)
	r.DELETE("/products/:id", db.DeleteProduct)
}


func (dbConfig *DBConfig) RegisterUserRoutes(r *gin.Engine) {
	r.GET("/user/:id", userroutes.GetUser)
	r.DELETE("/user/:id", userroutes.DeleteUser)
	r.POST("/user/:name/:passsword/:email", userroutes.CreateUser)
	r.PUT("/user/:id", userroutes.UpdateUser)
}


func (dbConfig *DBConfig) CreateProduct(c *gin.Context) {
	var newProduct types.Product

	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := database.ConnectDatabase()
	defer db.Close()

	query := `INSERT INTO products (name, description, price, stock_quantity, image_url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) RETURNING id`
	
	err := db.QueryRow(context.Background(), query, newProduct.Name, newProduct.Description, newProduct.Price, newProduct.StockQuantity, newProduct.ImageURL).Scan(&newProduct.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, newProduct)
}

func (dbConfig *DBConfig) GetProducts(c *gin.Context) {
	db := database.ConnectDatabase()
	defer db.Close()

	var products []types.Product
	rows, err := db.Query(context.Background(), "SELECT id, name, description, price, stock_quantity, image_url, created_at, updated_at FROM products")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product types.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.StockQuantity, &product.ImageURL, &product.CreatedAt, &product.UpdatedAt); err != nil {
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

func (dbConfig *DBConfig) GetProductById(c *gin.Context) {
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

func (dbConfig *DBConfig) UpdateProduct(c *gin.Context) {
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

func (dbConfig *DBConfig) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Missing product ID"})
		return
	}

	db := database.ConnectDatabase()
	defer db.Close()

	query := `DELETE FROM products WHERE id = $1;`

	result, err := db.Exec(context.Background(), query, id)
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
