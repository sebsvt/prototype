package repository

import "github.com/jmoiron/sqlx"

type paymentRepositoryDB struct {
	db *sqlx.DB
}

func NewPaymentRepositoryDB(db *sqlx.DB) PaymentRepository {
	return paymentRepositoryDB{db: db}
}

// CreateNewPayment implements PaymentRepository.
func (repo paymentRepositoryDB) CreateNewPayment(entity Payment) (int, error) {
	panic("unimplemented")
}

// FromPaymentID implements PaymentRepository.
func (repo paymentRepositoryDB) FromPaymentID(payment_id int) (*Payment, error) {
	panic("unimplemented")
}

// UpdatePayment implements PaymentRepository.
func (repo paymentRepositoryDB) UpdatePayment(entity Payment) error {
	panic("unimplemented")
}
