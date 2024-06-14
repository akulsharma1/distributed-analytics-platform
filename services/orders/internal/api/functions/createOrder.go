package functions

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/IBM/sarama"
	"github.com/akulsharma1/distributed-analytics-platform/services/orders/internal/api/models"
	"github.com/akulsharma1/distributed-analytics-platform/services/orders/internal/kafka"
	"github.com/gofiber/fiber/v2"
)

func CreateOrder(c *fiber.Ctx, inventoryChan chan kafka.InventoryMessage, orderRespChan chan kafka.OrderResponseMessage) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error": err,
			"message": "error checking stock",
		})
	}

	inventorykafkamsg := &kafka.InventoryMessage{OrderItems: order.OrderItems}

	inventoryJsonMsg, _ := json.Marshal(inventorykafkamsg)

	_, _, err := kafka.Producer.SendMessage(&sarama.ProducerMessage{
		Topic: kafka.GET_INVENTORY,
		Key: sarama.ByteEncoder(inventoryJsonMsg),
	})

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error": err,
			"message": "error checking stock",
		})
	}

	timeout := time.After(10 * time.Second)

	loop:
	for {
		select {
		case inventoryMsg := <-inventoryChan:
			hasEnoughStock, err := compareInventoryProducts(inventoryMsg, order)

			if err != nil {
				continue
			}

			if !hasEnoughStock {
				return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
					"success": false,
					"error": "unable to create order",
					"message": "not enough stock",
				})
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

	orderMsg, _ := json.Marshal(order)

	_, _, err = kafka.Producer.SendMessage(&sarama.ProducerMessage{
		Topic: kafka.CREATE_ORDER,
		Key: sarama.ByteEncoder(orderMsg),
	})

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error": err,
			"message": "error creating order",
		})
	}

	timeout = time.After(10 * time.Second)

	loop2:
	for {
		select {
		case response := <-orderRespChan:
			if response.OrderID != int(order.ID) {
				continue
			}

			if (!response.Success) {
				return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
					"success": false,
					"error": "items out of stock",
					"message": "error creating order",
					"data": response.OutOfStockVariants,
				})
			}

			break loop2
		case <-timeout:
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error": "timeout creating order",
				"message": "error creating order",
			})
		}
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "created order",
	})

	
}

func compareInventoryProducts(inventoryMessage kafka.InventoryMessage, order models.Order) (bool, error) {
	inventoryVariants := make(map[uint]int)

	for _, orderItem := range inventoryMessage.OrderItems {
		inventoryVariants[orderItem.Variant.ID] = orderItem.Variant.Quantity
	}

	for _, orderItem := range order.OrderItems {
		quantity, ok := inventoryVariants[orderItem.Variant.ID]
		if !ok {
			return false, errors.New("invalid inventory msg")
		}
		if quantity < orderItem.Variant.Quantity {
			return false, nil
		}
	}

	return true, nil
}