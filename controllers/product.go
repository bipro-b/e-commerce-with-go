package controllers

import (
	"github.com/bipro-b/ecommerce-backend/database"
	"github.com/bipro-b/ecommerce-backend/models"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)
	c.JSON(200, products)
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	c.ShouldBindJSON(&product)
	database.DB.Create(&product)
	c.JSON(201, product)
}
