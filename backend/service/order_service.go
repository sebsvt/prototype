package service

import (
	"errors"
	"io"
	"log"
	"mime/multipart"
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
func (srv orderService) CreateNewOrder(customer_id int, entity OrderRequest) (*OrderResponse, error) {
	new_payment := PaymentRequest{
		Amount: entity.ProductCost * float64(entity.Duration),
	}
	payment_id, err := srv.payment_service.CreateNewPayment(new_payment)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	order, err := srv.order_repo.CreateNewOrder(repository.Order{
		CustomerID:  customer_id,
		ProductSKU:  entity.ProductSKU,
		ProductCost: entity.ProductCost,
		Duration:    entity.Duration,
		PaymentID:   payment_id,
		IsPaid:      false,
		CreatedAt:   time.Now(),
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &OrderResponse{
		OrderID:     order.OrderID,
		CustomerID:  order.CustomerID,
		ProductSKU:  order.ProductSKU,
		ProductCost: order.ProductCost,
		Duration:    order.Duration,
		TotalCost:   order.ProductCost * float64(order.Duration),
		IsPaid:      order.IsPaid,
		CreatedAt:   order.CreatedAt,
	}, nil
}

// GetAllOrders implements OrderService.
func (srv orderService) GetAllOrders(user_id int) ([]OrderResponse, error) {
	var orders []OrderResponse
	order_query, err := srv.order_repo.GetAllOrder(user_id)
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
func (srv orderService) GetOrderFromOrderID(order_id int, user_id int) (*OrderResponse, error) {
	order, err := srv.order_repo.FromOrderID(order_id)
	if err != nil {
		return nil, err
	}
	if order.CustomerID != user_id {
		return nil, errors.New("permission denied")
	}
	return &OrderResponse{
		OrderID:     order.OrderID,
		CustomerID:  order.CustomerID,
		ProductSKU:  order.ProductSKU,
		ProductCost: order.ProductCost,
		Duration:    order.Duration,
		IsPaid:      order.IsPaid,
		TotalCost:   order.ProductCost * float64(order.Duration),
		CreatedAt:   order.CreatedAt,
	}, nil
}
func (srv orderService) CheckOut(order_id int, file *multipart.FileHeader) error {
	order, err := srv.order_repo.FromOrderID(order_id)
	if err != nil {
		return err
	}

	fileData, err := file.Open()
	if err != nil {
		return err
	}
	defer fileData.Close()

	byteData, err := io.ReadAll(fileData)
	if err != nil {
		return err
	}

	success, err := srv.payment_service.CheckPaymentSlip(order_id, byteData)
	if err != nil {
		log.Println("I am still here")
		return err
	}
	if !success {
		return errors.New("payment verification failed")
	}

	order.IsPaid = true
	err = srv.order_repo.UpdateOrder(*order)
	if err != nil {
		return err
	}

	return nil
}
