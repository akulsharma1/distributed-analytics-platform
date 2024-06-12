package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/IBM/sarama"
	"github.com/akulsharma1/distributed-analytics-platform/services/orders/internal/api/models"
	"github.com/akulsharma1/distributed-analytics-platform/services/orders/internal/kafka"
	"github.com/gofiber/fiber/v2"
)

func CreateOrder (c *fiber.Ctx, inventoryChan chan kafka.InventoryMessage) error {
	var products []models.Product

	if err := c.BodyParser(&products); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error": err,
			"message": "error checking stock",
		})
	}

	message := &kafka.InventoryMessage{Products: products}

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
	
	loop:
	for {
        select {
        case inventoryMessage := <-inventoryChan:
			if !compareInventoryProducts(inventoryMessage, products) {
				continue
			}

			inventoryStock := map[string]int{}
			for _, inventoryProduct := range inventoryMessage.Products {
				for _, variant := range inventoryProduct.Variants {
					inventoryStock[fmt.Sprintf("%v_%v", inventoryProduct.ID, variant.Size)] = variant.Quantity
				}
			}

			for _, product := range products {
				for _, variant := range product.Variants {
					remainingStock, ok := inventoryStock[fmt.Sprintf("%v_%v", product.ID, variant.Size)]
					if !ok || remainingStock <= 0 {
						return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
							"success": false,
							"error": "unable to create order",
							"message": "not enough stock",
						})
					}
				}
			}
			break loop

        case <-timeout:
            return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
                "success": false,
                "error": "timeout waiting for stock information",
                "message": "error checking stock",
            })
        }
    }

	// TODO: create kafka message for adding order
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "created order",
	})

}

func compareInventoryProducts(inventoryMessage kafka.InventoryMessage, products []models.Product) bool {
	inventoryProducts := make(map[uint]bool)

	for _, product := range inventoryMessage.Products {
		inventoryProducts[product.ID] = true
	}

	for _, product := range products {
		_, ok := inventoryProducts[product.ID]
		if !ok {
			return false
		}
	}

	return true
}