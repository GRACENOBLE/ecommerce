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
	
	protectedUserRoutes := r.Group("/user")	
	protectedUserRoutes.Use(middleware.JWTAuthMiddleware())
	{
		protectedUserRoutes.DELETE("/:id", dbConfig.DeleteUser)
		// r.PUT("/:id", db.UpdateUser)
	}
	r.POST("/refresh", dbConfig.RefreshToken)
	r.POST("/user", dbConfig.CreateUser)
	r.POST("/login", dbConfig.Login)

}
