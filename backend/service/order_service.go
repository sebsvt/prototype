package service

import (
	"time"

	"github.com/sebsvt/prototype/repository"
)

type orderService struct {
	order_repo repository.OrderRepository
}

func NewOrderService(order_repo repository.OrderRepository) OrderService {
	return orderService{order_repo: order_repo}
}

// CreateNewOrder implements OrderService.
func (srv orderService) CreateNewOrder(entity OrderRequest) (*OrderResponse, error) {
	order, err := srv.order_repo.CreateNewOrder(repository.Order{
		CustomerID:  entity.CustomerID,
		ProductSKU:  entity.ProductSKU,
		ProductCost: entity.ProductCost,
		Duration:    entity.Duration,
		PaymentID:   1,
		CreatedAt:   time.Now(),
	})
	if err != nil {
		return nil, err
	}
	is_paid := true
	if order.PaymentID == 2 {
		is_paid = false
	}
	return &OrderResponse{
		OrderID:     order.OrderID,
		CustomerID:  order.CustomerID,
		ProductSKU:  order.ProductSKU,
		ProductCost: order.ProductCost,
		Duration:    order.Duration,
		TotalCost:   order.ProductCost * float64(order.Duration),
		IsPaid:      is_paid,
		CreatedAt:   order.CreatedAt,
	}, nil
}

// GetAllOrders implements OrderService.
func (srv orderService) GetAllOrders() ([]OrderResponse, error) {
	var orders []OrderResponse
	order_query, err := srv.order_repo.GetAllOrder()
	if err != nil {
		return nil, err
	}
	for _, order := range order_query {
		orders = append(orders, OrderResponse{
			OrderID:     order.OrderID,
			CustomerID:  order.CustomerID,
			ProductSKU:  order.ProductSKU,
			ProductCost: order.ProductCost,
			Duration:    order.Duration,
			TotalCost:   order.ProductCost * float64(order.Duration),
			CreatedAt:   order.CreatedAt,
		})
	}
	return orders, nil
}

// GetOrderFromOrderID implements OrderService.
func (srv orderService) GetOrderFromOrderID(order_id int) (*OrderResponse, error) {
	order, err := srv.order_repo.FromOrderID(order_id)
	if err != nil {
		return nil, err
	}
	return &OrderResponse{
		OrderID:     order.OrderID,
		CustomerID:  order.CustomerID,
		ProductSKU:  order.ProductSKU,
		ProductCost: order.ProductCost,
		Duration:    order.Duration,
		TotalCost:   order.ProductCost * float64(order.Duration),
		CreatedAt:   order.CreatedAt,
	}, nil
}
