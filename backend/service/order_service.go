package service

import (
	"time"

	"github.com/sebsvt/prototype/repository"
)

type orderService struct {
	order_repo      repository.OrderRepository
	payment_service PaymentService
}

func NewOrderService(order_repo repository.OrderRepository, payment_srv PaymentService) OrderService {
	return orderService{order_repo: order_repo, payment_service: payment_srv}
}

// CreateNewOrder implements OrderService.
func (srv orderService) CreateNewOrder(entity OrderRequest) (*OrderResponse, error) {
	new_payment := PaymentRequest{
		Amount: entity.ProductCost * float64(entity.Duration),
	}
	payment_id, err := srv.payment_service.CreateNewPayment(new_payment)
	if err != nil {
		return nil, err
	}
	order, err := srv.order_repo.CreateNewOrder(repository.Order{
		CustomerID:  entity.CustomerID,
		ProductSKU:  entity.ProductSKU,
		ProductCost: entity.ProductCost,
		Duration:    entity.Duration,
		PaymentID:   payment_id,
		CreatedAt:   time.Now(),
	})
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
		IsPaid:      false,
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
