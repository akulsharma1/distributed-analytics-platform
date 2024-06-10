package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func SetUpDb() *gorm.DB {
	url := os.Getenv("INVENTORY_URL")

	if url == "" {
		panic("invalid inventory url")
	}

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		panic("err connecting to inventory db: "+err.Error())
	}

	return db
}