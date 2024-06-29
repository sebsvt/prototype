package service

import "time"

type OrderRequest struct {
	CustomerID  int     `json:"customer_id"`
	ProductSKU  string  `json:"product_sku"`
	ProductCost float64 `json:"product_cost"`
	Duration    int     `json:"duration"`
}

type OrderResponse struct {
	OrderID     int       `json:"order_id"`
	CustomerID  int       `json:"customer_id"`
	ProductSKU  string    `json:"product_sku"`
	ProductCost float64   `json:"product_cost"`
	Duration    int       `json:"duration"`
	TotalCost   float64   `json:"total_cost"`
	CreatedAt   time.Time `json:"created_at"`
}

type OrderService interface {
	CreateNewOrder(OrderRequest) (*OrderResponse, error)
	GetOrderFromOrderID(order_id int) (*OrderResponse, error)
	GetAllOrders() ([]OrderResponse, error)
}
