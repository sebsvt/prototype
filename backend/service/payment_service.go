package service

import (
	"errors"
	"time"

	"github.com/sebsvt/prototype/repository"
)

type paymentService struct {
	payment_repo repository.PaymentRepository
}

func NewPaymentService(payment_repo repository.PaymentRepository) PaymentService {
	return paymentService{payment_repo: payment_repo}
}

// CreateNewPayment implements PaymentService.
func (srv paymentService) CreateNewPayment(payment PaymentRequest) (int, error) {
	new_payment_transaction := repository.Payment{
		Receiver:   "Saharat Muksarn",
		Amount:     payment.Amount,
		IsVerified: false,
		CreatedAt:  time.Now(),
	}
	payment_id, err := srv.payment_repo.CreateNewPayment(new_payment_transaction)
	if err != nil {
		return 0, err
	}
	return payment_id, nil
}

// FromPaymentID implements PaymentService.
func (srv paymentService) FromPaymentID(payment_id int) (*PaymentResponse, error) {
	payment, err := srv.payment_repo.FromPaymentID(payment_id)
	if err != nil {
		return nil, err
	}
	return &PaymentResponse{
		PaymentID:            payment_id,
		Sender:               payment.Sender,
		Receiver:             payment.Receiver,
		Amount:               payment.Amount,
		IsVerified:           payment.IsVerified,
		TransactionRef:       payment.TransactionRef,
		TransactionTimeStamp: payment.TransactionTimeStamp,
		CreatedAt:            payment.CreatedAt,
	}, nil
}

// VerifyPayment implements PaymentService.
func (srv paymentService) VerifyPayment(paymend_id int) error {
	payment, err := srv.payment_repo.FromPaymentID(paymend_id)
	if err != nil {
		return err
	}
	if payment.IsVerified {
		return errors.New("payment is verified")
	}
	payment.IsVerified = true
	if err := srv.payment_repo.UpdatePayment(*payment); err != nil {
		return err
	}
	return nil
}
