package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	r.GET("/products", getProducts)
	r.POST("/products", createProduct)
	r.PUT("/products/:id", updateProduct)
	r.DELETE("/products/:id", deleteProduct)
	// Add more product-related routes here
}


func RegisterUserRoutes(r *gin.Engine) {
	r.GET("/user/:id", getUser)
	r.DELETE("/user/:id", deleteUser)
	r.POST("/user/:name/:passsword/:email", createUser)
	r.PUT("/user/:id", updateUser)
}