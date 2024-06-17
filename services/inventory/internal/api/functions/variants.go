package functions

import (
	"net/http"
	"strconv"

	"github.com/akulsharma1/distributed-analytics-platform/services/inventory/internal/api/models"
	"github.com/akulsharma1/distributed-analytics-platform/services/inventory/internal/db"
	"github.com/gofiber/fiber/v2"
)

func AddVariant(c *fiber.Ctx) error {
	productIDStr := c.Params("product_id")

	productID, err := strconv.ParseUint(productIDStr, 10, 32)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "invalid product_id",
        })
    }

	variant := models.Variant{}
	if err := c.BodyParser(&variant); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "error parsing variant data",
        })
    }

	variant.ProductID = int(productID)
	
	if err := db.DATABASE.Create(&variant).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "failed to add variant to product",
			"error":   err.Error(),
		})
	}
	
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "added variant to product",
	})
}