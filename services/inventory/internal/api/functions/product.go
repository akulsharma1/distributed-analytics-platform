package functions

import (
	"net/http"
	"strconv"

	"github.com/akulsharma1/distributed-analytics-platform/services/inventory/internal/api/models"
	"github.com/akulsharma1/distributed-analytics-platform/services/inventory/internal/db"
	"github.com/gofiber/fiber/v2"
)

func AddProduct (c *fiber.Ctx) error {
	var product = models.Product{}
	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "error unmarshaling json to product",
			"error": err,
		})
	}

	if err := db.DATABASE.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "failed to create product",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "added product to db",
	})
}

func GetProducts (c *fiber.Ctx) error {
	productIDStr := c.Query("product_id")

	if productIDStr == "" {
		products := []models.Product{}
		
		if err := db.DATABASE.Limit(20).Find(&products).Error; err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"message": "error getting products",
				"error": err,
			})
		}

		return c.Status(http.StatusOK).JSON(&fiber.Map{
			"success": true,
			"message": "got products",
			"data": products,
		}) 
	} else {
		productID, err := strconv.ParseUint(productIDStr, 10, 32)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid product_id",
			})
		}

		product := models.Product{}
		if err := db.DATABASE.Where("id = ?", uint(productID)).First(&product).Error; err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"message": "error getting product",
				"error": err,
			})
		}

		return c.Status(http.StatusOK).JSON(&fiber.Map{
			"success": true,
			"message": "got product",
			"data": product,
		})
	}
}