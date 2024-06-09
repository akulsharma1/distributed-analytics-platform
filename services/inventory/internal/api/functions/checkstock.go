package functions

import (
	"net/http"
	"strconv"

	"github.com/akulsharma1/distributed-analytics-platform/services/inventory/internal/api/models"
	"github.com/akulsharma1/distributed-analytics-platform/services/inventory/internal/db"
	"github.com/gofiber/fiber/v2"
)

func CheckStock(c *fiber.Ctx) error {
	productIdStr := c.Query("product_id")

	if productIdStr == "" {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error": "product_id query is empty",
			"message": "error checking stock",
		})
	}

	productID, err := strconv.ParseUint(productIdStr, 10, 32)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error": err,
			"message": "error converting productID to string",
		})
	}

	var variants []struct {
        Size     string  `json:"size"`
        Price    float64 `json:"price"`
        Quantity int     `json:"quantity"`
    }

	if err := db.DATABASE.Model(&models.Variant{}).Where("product_id = ?", uint(productID)).Scan(&variants).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error": err,
			"message": "error getting variants",
		})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"success": true,
		"data": variants,
		"message": "got stock",
	})
}