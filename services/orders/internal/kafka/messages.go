package kafka

import "github.com/akulsharma1/distributed-analytics-platform/services/orders/internal/api/models"

type InventoryMessage struct {
	OrderItems []models.OrderItem `json:"order_items"`
}

type OrderResponseMessage struct {
	Success bool `json:"success"`
	OrderID int `json:"order_id"`
	OutOfStockVariants []models.Variant `json:"variants"`
}