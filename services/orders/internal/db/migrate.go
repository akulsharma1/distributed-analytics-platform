package db

import (
	// "github.com/akulsharma1/distributed-analytics-platform/services/orders/internal/api/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.Exec(`
	CREATE TABLE IF NOT EXISTS customers (
		id SERIAL NOT NULL,
		created_at TIMESTAMPTZ,
		updated_at TIMESTAMPTZ,
		deleted_at TIMESTAMPTZ,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE,
		PRIMARY KEY (email)
	);`)
	
	db.Exec(`
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		created_at TIMESTAMPTZ,
        updated_at TIMESTAMPTZ,
        deleted_at TIMESTAMPTZ,
		name VARCHAR(100) NOT NULL
	);`)

	db.Exec(`
	CREATE TABLE IF NOT EXISTS variants (
		id SERIAL NOT NULL,
		created_at TIMESTAMPTZ,
		updated_at TIMESTAMPTZ,
		deleted_at TIMESTAMPTZ,
		product_id INT NOT NULL,
		size VARCHAR(100) NOT NULL,
		price decimal(10,2) NOT NULL,
		quantity INT NOT NULL,
		PRIMARY KEY (product_id, size),
    	FOREIGN KEY (product_id) REFERENCES products(id)
	);`)

	db.Exec(`
    CREATE TABLE IF NOT EXISTS orders (
        id SERIAL PRIMARY KEY,
        created_at TIMESTAMPTZ,
        updated_at TIMESTAMPTZ,
        deleted_at TIMESTAMPTZ,
        product_name VARCHAR(100) NOT NULL,
        quantity INT NOT NULL,
        customer_email VARCHAR(100) NOT NULL REFERENCES customers(email)
    );`)

	db.Exec(`
	CREATE TABLE IF NOT EXISTS orderitems (
		id SERIAL NOT NULL,
		created_at TIMESTAMPTZ,
		updated_at TIMESTAMPTZ,
		deleted_at TIMESTAMPTZ,
		order_id INT NOT NULL,
		product_id INT NOT NULL,
    	size VARCHAR(100) NOT NULL,
		quantity INT NOT NULL,
		PRIMARY KEY (order_id, product_id, size),
		FOREIGN KEY (order_id) REFERENCES orders(id),
		FOREIGN KEY (product_id, size) REFERENCES variants(product_id, size)
	);`)

	// db.Migrator()
}