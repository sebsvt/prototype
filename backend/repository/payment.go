package repository

import "time"

type Payment struct {
	PaymentID            int       `db:"payment_id"`
	Sender               string    `db:"sender"`
	Receiver             string    `db:"receiver"`
	Amount               float64   `db:"amount"`
	IsVerified           bool      `db:"is_verified"`
	TransactionRef       string    `db:"transaction_ref"`
	TransactionTimeStamp string    `db:"transaction_time_stamp"`
	CreatedAt            time.Time `db:"created_at"`
}

type PaymentRepository interface {
	CreateNewPayment(entity Payment) (int, error)
	FromPaymentID(payment_id int) (*Payment, error)
	UpdatePayment(entity Payment) error
}
