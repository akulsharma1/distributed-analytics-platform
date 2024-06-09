package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name string `gorm:"type:varchar(100);not null"`
	Variants []Variant
}

type Variant struct {
	gorm.Model
	ProductID uint    `gorm:"not null"`                          // Foreign key
	Price     float64 `gorm:"type:decimal(10,2);not null"`       // Price as a decimal
	Quantity  int     `gorm:"type:int;not null"`
	Size      string  `gorm:"type:varchar(100)"`                 // Size description (e.g., S, M, L, XL)
}
