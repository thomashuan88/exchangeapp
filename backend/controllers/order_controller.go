package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"exchangeapp/backend/global"
	"exchangeapp/backend/models"
)

// OrderController handles order-related requests
type OrderController struct{}

// NewOrderController returns a new OrderController instance
func NewOrderController() *OrderController {
	return &OrderController{}
}

// CreateOrder creates a new order
func (oc *OrderController) CreateOrder(c *gin.Context) {
	var order models.Order

	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := global.Db.AutoMigrate(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Insert order into database
	if err := global.Db.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// GetOrdersByCustomer retrieves orders for each customer and sums the total
func (oc *OrderController) GetOrdersByCustomer(c *gin.Context) {

	// Use query to group orders by customer and sum total
	var customerOrders []struct {
		CustomerID int     `json:"customer_id"`
		Total      float64 `json:"total"`
	}
	if err := global.Db.
		Table("orders").
		Select("customer_id, SUM(total) as total").
		Group("customer_id").
		Scan(&customerOrders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return result as JSON
	c.JSON(http.StatusOK, customerOrders)
}
