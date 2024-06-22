package ordering

import (
	"log"
	"time"

	"github.com/google/uuid"
	domain "github.com/sebsvt/prototype/domain/ordering"
)

type orderingService struct {
	order_repository domain.OrderRepository
}

func NewOrderService(order_repo domain.OrderRepository) domain.OrderService {
	return orderingService{order_repository: order_repo}
}

// CreateNewOrder implements ordering.OrderService.
func (srv orderingService) CreateNewOrder(new_order domain.OrderCreated) (domain.OrderBase, error) {
	customer_id, err := uuid.Parse(new_order.CustomerID)
	if err != nil {
		log.Println(err)
		return domain.OrderBase{}, err
	}
	product_id, err := uuid.Parse(new_order.ProductID)
	if err != nil {
		log.Println(err)
		return domain.OrderBase{}, err
	}
	order := domain.NewOrder(customer_id, product_id)
	order.CreatedAt = time.Now()
	//product price must assing below before saving data
	err = srv.order_repository.Save(order)
	if err != nil {
		log.Println(err)
		return domain.OrderBase{}, err
	}
	return domain.OrderBase{
		Reference:  order.Reference.String(),
		CustomerID: order.CustomerID.String(),
		ProductID:  order.ProductID.String(),
		Price:      order.Price,
		IsPaid:     order.IsPaid,
		CreatedAt:  order.CreatedAt,
	}, nil
}

// GetOrderByOrderReference implements ordering.OrderService.
func (srv orderingService) GetOrderByOrderReference(reference string) (domain.OrderBase, error) {
	order_ref, err := uuid.Parse(reference)
	if err != nil {
		log.Println(err)
		return domain.OrderBase{}, err
	}
	order, err := srv.order_repository.FromReference(order_ref)
	if err != nil {
		log.Println(err)
		return domain.OrderBase{}, err
	}
	response_model := domain.OrderBase{
		Reference:  order.Reference.String(),
		CustomerID: order.CustomerID.String(),
		ProductID:  order.ProductID.String(),
		Price:      order.Price,
		IsPaid:     order.IsPaid,
		CreatedAt:  order.CreatedAt,
	}
	return response_model, nil
}
