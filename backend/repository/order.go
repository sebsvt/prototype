package repository

import "time"

type Order struct {
	OrderID     int       `db:"order_id"`
	CustomerID  int       `db:"customer_id"`
	ProductSKU  string    `db:"product_sku"`
	ProductCost float64   `db:"product_cost"`
	Duration    int       `db:"duration"`
	CreatedAt   time.Time `db:"created_at"`
}

type OrderRepository interface {
	CreateNewOrder(entity Order) (*Order, error)
	FromOrderID(order_id int) (*Order, error)
	GetAllOrder() ([]Order, error)
}
