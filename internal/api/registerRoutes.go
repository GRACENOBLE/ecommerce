package api

import (
	"time"

	"github.com/GRACENOBLE/ecommerce/internal/helpers/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, dbConfig *DBConfig) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://scalable-ecom-frontend.vercel.app/"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
		protectedUserRoutes.PUT("/:id", dbConfig.UpdateUser)
		protectedUserRoutes.GET("/data", dbConfig.ReturnUserData)
	}
	r.POST("/refresh", dbConfig.RefreshToken)
	r.POST("/user", dbConfig.CreateUser)
	r.POST("/login", dbConfig.Login)

}
