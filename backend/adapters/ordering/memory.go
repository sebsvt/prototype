package ordering

import (
	"sync"

	"github.com/google/uuid"
	domain "github.com/sebsvt/prototype/domain/ordering"
)

type OrderRepositoryMemory struct {
	orders map[uuid.UUID]domain.Order
	sync.Mutex
}

func NewOrderRepositoryMemory() *OrderRepositoryMemory {
	return &OrderRepositoryMemory{
		orders: make(map[uuid.UUID]domain.Order),
	}
}

// FromCustomerID implements ordering.OrderRepository.
func (repo *OrderRepositoryMemory) FromCustomerID(cust_id uuid.UUID) ([]domain.Order, error) {
	return nil, nil
}

// FromReference implements ordering.OrderRepository.
func (repo *OrderRepositoryMemory) FromReference(ref uuid.UUID) (*domain.Order, error) {
	if order, has := repo.orders[ref]; has {
		return &order, nil
	}
	return nil, domain.ErrOrderDoesNotExist
}

// Save implements ordering.OrderRepository.
func (repo *OrderRepositoryMemory) Save(new_order domain.Order) error {
	if _, has := repo.orders[new_order.Reference]; has {
		return domain.ErrOrderIsAlreadyExist
	}
	repo.Lock()
	repo.orders[new_order.Reference] = new_order
	repo.Unlock()
	return nil
}
