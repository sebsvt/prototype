package service

import (
	"mime/multipart"
	"time"
)

type OrderRequest struct {
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
	IsPaid      bool      `json:"is_paid"`
	CreatedAt   time.Time `json:"created_at"`
}

type OrderService interface {
	CreateNewOrder(int, OrderRequest) (*OrderResponse, error)
	GetOrderFromOrderID(order_id int, customer_id int) (*OrderResponse, error)
	GetAllOrders(user_id int) ([]OrderResponse, error)
	CheckOut(order_id int, file *multipart.FileHeader) error
}
