package ordering

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrOrderDoesNotExist   = errors.New("order does not exits")
	ErrOrderIsAlreadyExist = errors.New("order's already exist")
)

type OrderRepository interface {
	FromReference(uuid.UUID) (*Order, error)
	FromCustomerID(uuid.UUID) ([]Order, error)
	Save(Order) error
}
