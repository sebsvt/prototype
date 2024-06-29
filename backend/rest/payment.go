package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebsvt/prototype/service"
)

type paymentRestAPI struct {
	payment_service service.PaymentService
}

func NewPaymentRestAPI(payment_service service.PaymentService) paymentRestAPI {
	return paymentRestAPI{payment_service: payment_service}
}

func (h paymentRestAPI) CreatePayment(c *fiber.Ctx) error {
	var payment_req service.PaymentRequest
	if err := c.BodyParser(&payment_req); err != nil {
		return err
	}
	payment_id, err := h.payment_service.CreateNewPayment(payment_req)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"payment_id": payment_id,
	})
}

func (h paymentRestAPI) GetPaymentFromID(c *fiber.Ctx) error {
	payment_id, err := c.ParamsInt("payment_id")
	if err != nil {
		return err
	}
	payment, err := h.payment_service.FromPaymentID(payment_id)
	if err != nil {
		return err
	}
	return c.JSON(payment)
}
