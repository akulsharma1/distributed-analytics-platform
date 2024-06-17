package models

import "gorm.io/gorm"

type Product struct {
    gorm.Model
    Name     string    `gorm:"type:varchar(100);not null" json:"name"`
    Variants []Variant `json:"variants"`
}

type Variant struct {
    gorm.Model
    ProductID int    `gorm:"type:int;not null" json:"product_id"` // Part of the unique index
    Size      string  `gorm:"type:varchar(100);not null" json:"size"` // Part of the unique index
    Price     float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	
    Quantity  int     `gorm:"type:int;not null" json:"quantity"`
    Product   Product `gorm:"foreignKey:ProductID"` // Reference to Product

	OrderItems []OrderItem
}
