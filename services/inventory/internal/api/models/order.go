package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerEmail string `gorm:"type:varchar(100);not null;index"`
	Customer    Customer `gorm:"foreignKey:CustomerEmail;references:Email"`

	OrderItems []OrderItem
}

type OrderItem struct {
    gorm.Model
    OrderID   int   `gorm:"not null;index;primaryKey"` // Composite primary key
    VariantID int   `gorm:"not null;index;primaryKey"` // Composite primary key
    Quantity  int    `gorm:"type:int;not null"`

    Order     Order  `gorm:"foreignKey:OrderID"` // Relation to Order
    Variant   Variant `gorm:"foreignKey:VariantID;references:id"` // Relation to Variant
}
