package functions

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/IBM/sarama"
	"github.com/akulsharma1/distributed-analytics-platform/services/orders/internal/api/models"
	"github.com/akulsharma1/distributed-analytics-platform/services/orders/internal/kafka"
	"github.com/gofiber/fiber/v2"
)

/*
Get product stock. Only used for one product
*/
func GetProductStock (c *fiber.Ctx, inventoryChan chan kafka.InventoryMessage) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error": err,
			"message": "error checking stock",
		})
	}

	message := &kafka.InventoryMessage{Products: []models.Product{product}}

	jsonMsg, _ := json.Marshal(message)

	_, _, err := kafka.Producer.SendMessage(&sarama.ProducerMessage{
		Topic: kafka.GET_INVENTORY,
		Key: sarama.ByteEncoder(jsonMsg),
	})

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error": err,
			"message": "error checking stock",
		})
	}

	timeout := time.After(30 * time.Second)
	
	for {
        select {
        case inventoryMessage := <-inventoryChan:
			products := inventoryMessage.Products
			if len(products) == 0 {
				continue
			}
            if products[0].ID == uint(product.ID) {
                return c.Status(http.StatusOK).JSON(&fiber.Map{
                    "success": true,
                    "data": product,
                })
            }
        case <-timeout:
            return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
                "success": false,
                "error": "timeout waiting for stock information",
                "message": "error checking stock",
            })
        }
    }
}