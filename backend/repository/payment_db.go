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
	query := "insert into payments (sender, receiver, amount, is_verified, transaction_ref, transaction_time_stamp, created_at) values (?, ?, ?, ?, ?, ?, ?)"
	result, err := repo.db.Exec(
		query,
		entity.Sender,
		entity.Receiver,
		entity.Amount,
		entity.IsVerified,
		entity.TransactionRef,
		entity.TransactionTimeStamp,
		entity.CreatedAt,
	)
	if err != nil {
		return 0, err
	}
	payment_id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(payment_id), nil
}

// FromPaymentID implements PaymentRepository.
func (repo paymentRepositoryDB) FromPaymentID(payment_id int) (*Payment, error) {
	var payment Payment
	query := "select payment_id, sender, receiver, amount, is_verified, transaction_ref, transaction_time_stamp, created_at from payments where payment_id=?"
	if err := repo.db.Get(&payment, query, payment_id); err != nil {
		return nil, err
	}
	return &payment, nil
}

// UpdatePayment implements PaymentRepository.
func (repo paymentRepositoryDB) UpdatePayment(entity Payment) error {
	query := `update payments
		set sender=?, receiver=?, amount=?, is_verified=?, transaction_ref=?, transaction_time_stamp=?, created_at=?
		where payment_id=?`
	_, err := repo.db.Exec(
		query,
		entity.Sender,
		entity.Receiver,
		entity.Amount,
		entity.IsVerified,
		entity.TransactionRef,
		entity.TransactionTimeStamp,
		entity.CreatedAt,
		entity.PaymentID,
	)
	return err
}
