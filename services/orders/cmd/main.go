package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/akulsharma1/distributed-analytics-platform/services/orders/internal/api/functions"
	"github.com/akulsharma1/distributed-analytics-platform/services/orders/internal/db"
	"github.com/akulsharma1/distributed-analytics-platform/services/orders/internal/kafka"
	"github.com/ashah360/fibertools"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	currPath, _ := os.Getwd()

	godotenv.Load(filepath.Join(currPath, "..", "..", "..", ".env"))

	db.DATABASE = db.SetUpDb()

	// db.Migrate(db.DATABASE)

	kafka.SetUpKafka()
}

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: fibertools.ErrorHandler,
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	app.Use(fibertools.Recover())
	
	inventoryChannel := make(chan kafka.InventoryMessage)
	orderResponseChannel := make(chan kafka.OrderResponseMessage)

	// app.Post("/api/v1/stock", func (c *fiber.Ctx) error {
	// 	return functions.GetProductStock(c, inventoryChannel, orderResponseChannel)
	// })
	app.Post("/api/v1/order", func (c *fiber.Ctx) error {
		return functions.CreateOrder(c, inventoryChannel, orderResponseChannel)
	})

	var wg sync.WaitGroup

	wg.Add(1)
	go func () {
		for {
			select {
			case msg := <-kafka.InventoryPartition.Messages():
				var message kafka.InventoryMessage
				message.MessageID = string(msg.Key)
				if err := json.Unmarshal(msg.Value, &message); err != nil {
					log.Printf("Error unmarshaling message: %v", err)
                    continue
				}

				inventoryChannel <- message

			case err := <-kafka.InventoryPartition.Errors():
				log.Printf("Failed to consume message: %v", err)
			}
		}
	}()

	wg.Add(1)
	go func() {
		for {
			select {
			case msg := <-kafka.OrderPartition.Messages():
				var message kafka.OrderResponseMessage
				message.MessageID = string(msg.Key)
				if err := json.Unmarshal(msg.Value, &message); err != nil {
					log.Printf("Error unmarshaling message: %v", err)
                    continue
				}

				orderResponseChannel <- message
			case err := <-kafka.InventoryPartition.Errors():
				log.Printf("Failed to consume message: %v", err)
			}
		}
	}()

	wg.Add(1)
	go func() {
		if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
			log.Println("An error occured, shutting down gracefully. Error ", err)
			_ = app.Shutdown()
		}
	}()

	wg.Wait()
}