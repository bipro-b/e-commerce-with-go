package controllers

import (
	"net/http"

	"github.com/bipro-b/ecommerce-backend/database"
	"github.com/bipro-b/ecommerce-backend/models"
	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, order)
}
