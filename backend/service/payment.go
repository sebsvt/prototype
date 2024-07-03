package service

import "time"

type PaymentRequest struct {
	Amount float64 `json:"amount"`
}

type PaymentResponse struct {
	PaymentID            int       `json:"payment_id"`
	Sender               string    `json:"sender"`
	Receiver             string    `json:"receiver"`
	Amount               float64   `json:"amount"`
	IsVerified           bool      `json:"is_verified"`
	TransactionRef       string    `json:"transaction_ref"`
	TransactionTimeStamp string    `json:"transaction_time_stamp"`
	CreatedAt            time.Time `json:"created_at"`
}

type PaymentService interface {
	CreateNewPayment(payment PaymentRequest) (int, error)
	FromPaymentID(payment_id int) (*PaymentResponse, error)
	VerifyPayment(paymend_id int) error
	CheckPaymentSlip(orderID int, fileData []byte) (bool, error)
}
