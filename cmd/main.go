package main

import (
	"github.com/bipro-b/ecommerce-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
