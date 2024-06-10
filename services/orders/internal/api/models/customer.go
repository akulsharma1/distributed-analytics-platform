package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model    // Includes fields ID, CreatedAt, UpdatedAt, DeletedAt
    Name  string  `gorm:"type:varchar(100);not null" json:"name"`
    Email string  `gorm:"type:varchar(100);not null;unique;primaryKey" json:"email"`
	Orders []Order
}