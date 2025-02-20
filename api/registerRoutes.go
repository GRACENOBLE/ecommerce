package api

import (
	"github.com/GRACENOBLE/ecommerce/internal/helpers/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, dbConfig *DBConfig) {

	productRoutes := r.Group("/products")
	productRoutes.Use(middleware.JWTAuthMiddleware())
	{
		productRoutes.GET("/", dbConfig.GetProducts)
		productRoutes.GET("/:id", dbConfig.GetProductById)
		productRoutes.POST("/", dbConfig.CreateProduct)
		productRoutes.PUT("/:id", dbConfig.UpdateProduct)
		productRoutes.DELETE("/:id", dbConfig.DeleteProduct)
	}

	// r.GET("/user/:id", db.GetUser)
	// r.DELETE("/user/:id", db.DeleteUser)
	r.POST("/user", dbConfig.CreateUser)
	r.POST("/login", dbConfig.Login)
	r.POST("/refresh", dbConfig.RefreshToken)

	// r.PUT("/user/:id", db.UpdateUser)
}
