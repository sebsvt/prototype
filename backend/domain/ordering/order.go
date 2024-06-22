package ordering

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrOrderAlreadyPaid  = errors.New("order's already paid")
	ErrOrderIsUnpaid     = errors.New("order is unpaid")
	ErrProductIsNotExist = errors.New("order's product does not exist")
)

type Order struct {
	Reference  uuid.UUID `bson:"reference"`
	CustomerID uuid.UUID `bson:"customer_id"`
	ProductID  uuid.UUID `bson:"product_id"`
	PaymentID  uuid.UUID `bson:"payment_id"`
	Price      float64   `bson:"price"`
	IsPaid     bool      `bson:"is_paid"`
	CreatedAt  string    `bson:"created_at"`
}

func NewOrder(customer_id, product_id uuid.UUID) Order {
	return Order{
		Reference:  uuid.New(),
		CustomerID: customer_id,
		ProductID:  product_id,
	}
}
