package models

import "gorm.io/gorm"

type Product struct {
    gorm.Model
    Name     string    `gorm:"type:varchar(100);not null" json:"name"`
    Variants []Variant `json:"variants"`
}

type Variant struct {
    gorm.Model
    ProductID uint    `gorm:"index:idx_product_size,unique;autoIncrement:false" json:"product_id"` // Part of the unique index
    Size      string  `gorm:"index:idx_product_size,unique;type:varchar(100)" json:"size"` // Part of the unique index
    Price     float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	
    Quantity  int     `gorm:"type:int;not null" json:"quantity"`
    Product   Product `gorm:"foreignKey:ProductID"` // Reference to Product

	OrderItems []OrderItem
}