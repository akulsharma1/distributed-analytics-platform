package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akulsharma1/distributed-analytics-platform/services/orders/internal/db"
	"github.com/ashah360/fibertools"
	"github.com/gofiber/fiber/v2"
)

func init() {
	db.SetUpDb()
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

	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Println("An error occured, shutting down gracefully. Error ", err)
		_ = app.Shutdown()
	}
}