package service

type PaymentService interface {
	CreateNewPayment()
	VerifyPayment(paymend_id int)
}
