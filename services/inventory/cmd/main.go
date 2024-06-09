package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akulsharma1/distributed-analytics-platform/services/inventory/internal/api/functions"
	"github.com/akulsharma1/distributed-analytics-platform/services/inventory/internal/db"
	"github.com/ashah360/fibertools"
	"github.com/gofiber/fiber/v2"
)

func init() {
	db.DATABASE = db.SetUpDb()
}
func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: fibertools.ErrorHandler,
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app.Use(fibertools.Recover())

	app.Get("/api/v1/variants", functions.CheckStock)
	app.Get("/api/v1/stock", functions.CheckStock)

	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Println("An error occured, shutting down gracefully. Error ", err)
		_ = app.Shutdown()
	}
}