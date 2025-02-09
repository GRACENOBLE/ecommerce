package api

import (
	productroutes "github.com/GRACENOBLE/ecommerce/api/routes/product_routes"
	userroutes "github.com/GRACENOBLE/ecommerce/api/routes/user_routes"
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	r.GET("/products", productroutes.GetProducts)
	r.GET("/product/:id", productroutes.GetProductById)
	r.POST("/products", productroutes.CreateProduct)
	r.PUT("/products/:id", productroutes.UpdateProduct)
	r.DELETE("/products/:id", productroutes.DeleteProduct)
}


func RegisterUserRoutes(r *gin.Engine) {
	r.GET("/user/:id", userroutes.GetUser)
	r.DELETE("/user/:id", userroutes.DeleteUser)
	r.POST("/user/:name/:passsword/:email", userroutes.CreateUser)
	r.PUT("/user/:id", userroutes.UpdateUser)
}