package kafka

import "github.com/akulsharma1/distributed-analytics-platform/services/orders/internal/api/models"

type InventoryMessage struct {
	Products []models.Product `json:"products"`
}