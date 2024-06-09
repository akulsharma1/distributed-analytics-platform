package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name string `gorm:"type:varchar(100);not null"`
	Variants []Variant
}

type Variant struct {
    gorm.Model
    ProductID uint    `gorm:"primaryKey;autoIncrement:false"` // Part of the primary key, no auto increment
    Size      string  `gorm:"primaryKey;type:varchar(100)"`   // Part of the primary key
    Price     float64 `gorm:"type:decimal(10,2);not null"`
    Quantity  int     `gorm:"type:int;not null"`
    Product   Product `gorm:"foreignKey:ProductID"`           // Reference to Product
}
