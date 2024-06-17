package kafka

import "github.com/akulsharma1/distributed-analytics-platform/services/orders/internal/api/models"

type InventoryMessage struct {
	MessageID string `json:"message_id"`
	OrderItems []models.OrderItem `json:"order_items"`
}

type OrderResponseMessage struct {
	Success bool `json:"success"`
	MessageID string `json:"message_id"`
	OutOfStockVariants []models.Variant `json:"variants"`
}

type OrderCreationMessage struct {
	MessageID string `json:"message_id"`
	Order models.Order `json:"order"`
}