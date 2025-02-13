package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, dbConfig *DBConfig) {

	r.GET("/products", dbConfig.GetProducts)
	r.GET("/product/:id", dbConfig.GetProductById)
	r.POST("/products", dbConfig.CreateProduct)
	r.PUT("/products/:id", dbConfig.UpdateProduct)
	r.DELETE("/products/:id", dbConfig.DeleteProduct)
	// r.GET("/user/:id", db.GetUser)
	// r.DELETE("/user/:id", db.DeleteUser)
	// r.POST("/user/:name/:passsword/:email", db.CreateUser)
	// r.PUT("/user/:id", db.UpdateUser)
}
