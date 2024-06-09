package db

import (
	"github.com/akulsharma1/distributed-analytics-platform/services/inventory/internal/api/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
    db.AutoMigrate(&models.Product{}, &models.Variant{})
}