package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerEmail string    `gorm:"type:varchar(100);not null;index"` // Foreign key for Customer, indexed for performance
    Customer      Customer  `gorm:"foreignKey:CustomerEmail;references:Email"` // Association to Customer via Email
    OrderItems    []OrderItem `json:"order_items"`
}

type OrderItem struct {
    gorm.Model
    OrderID   uint    `gorm:"primaryKey;autoIncrement:false;index"`  // Part of the composite primary key, and indexed for performance
    Order     Order   `gorm:"foreignKey:OrderID"`                    // Association to Order
    VariantID uint    `gorm:"primaryKey;autoIncrement:false;index"`  // Part of the composite primary key, and indexed for performance
    Variant   Variant `gorm:"foreignKey:VariantID"`                  // Association to Variant
    Quantity  int     `gorm:"type:int;not null"`                     // Quantity of the Variant ordered
}