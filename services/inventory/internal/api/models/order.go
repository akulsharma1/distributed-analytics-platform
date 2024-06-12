package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
    ProductName string `gorm:"type:varchar(100);not null"`
	CustomerEmail string `gorm:"type:varchar(100);not null;index"`
	Customer    Customer `gorm:"foreignKey:CustomerEmail;references:Email"`

	OrderItem []OrderItem
}

type OrderItem struct {
    gorm.Model
    OrderID   uint   `gorm:"not null;index;primaryKey"` // Composite primary key
    ProductID uint   `gorm:"not null;index;primaryKey"` // Composite primary key
    Size      string `gorm:"type:varchar(100);not null;primaryKey"` // Composite primary key
    Quantity  int    `gorm:"type:int;not null"`

    Order     Order  `gorm:"foreignKey:OrderID"` // Relation to Order
    Variant   Variant `gorm:"foreignKey:ProductID,Size;references:ProductID,Size"` // Relation to Variant
}