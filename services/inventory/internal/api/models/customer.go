package models

import "gorm.io/gorm"

type Customer struct {
    gorm.Model
    Name  string `gorm:"type:varchar(100);not null"`
    Email string `gorm:"type:varchar(100);not null;unique;primaryKey"`
    Orders []Order
}