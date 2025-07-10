package routes

import (
	"github.com/bipro-b/ecommerce-backend/controllers"
	"github.com/bipro-b/ecommerce-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	api.GET("/products", controllers.GetProducts)
	api.POST("/products", middleware.RequireAuth, controllers.CreateProduct)

	api.POST("/orders", middleware.RequireAuth, controllers.CreateOrder)
}
