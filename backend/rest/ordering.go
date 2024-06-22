package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebsvt/prototype/domain/ordering"
)

type OrderRest struct {
	ordering_service ordering.OrderService
}

func NewOrderRest(ordering_srv ordering.OrderService) OrderRest {
	return OrderRest{ordering_service: ordering_srv}
}

func (h OrderRest) GetOrderFromReference(c *fiber.Ctx) error {
	ref := c.Params("ref")
	order, err := h.ordering_service.GetOrderByOrderReference(ref)
	if err != nil {
		return err
	}
	return c.JSON(order)
}

func (h OrderRest) CreateNewOrder(c *fiber.Ctx) error {
	var new_order ordering.OrderCreated
	if err := c.BodyParser(&new_order); err != nil {
		return err
	}
	response, err := h.ordering_service.CreateNewOrder(new_order)
	if err != nil {
		return err
	}
	return c.JSON(response)
}
