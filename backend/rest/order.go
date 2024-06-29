package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebsvt/prototype/service"
)

type orderRestAPI struct {
	order_srv service.OrderService
}

func NewOrderRestAPI(order_srv service.OrderService) orderRestAPI {
	return orderRestAPI{order_srv: order_srv}
}

func (h orderRestAPI) FromOrderID(c *fiber.Ctx) error {
	order_id, err := c.ParamsInt("order_id")
	if err != nil {
		return err
	}
	order, err := h.order_srv.GetOrderFromOrderID(order_id)
	if err != nil {
		return err
	}
	return c.JSON(order)
}

func (h orderRestAPI) GetAllOrders(c *fiber.Ctx) error {
	orders, err := h.order_srv.GetAllOrders()
	if err != nil {
		return err
	}
	return c.JSON(orders)
}

func (h orderRestAPI) CreateNewOrder(c *fiber.Ctx) error {
	var new_order service.OrderRequest
	if err := c.BodyParser(&new_order); err != nil {
		return err
	}
	order, err := h.order_srv.CreateNewOrder(new_order)
	if err != nil {
		return err
	}
	return c.JSON(order)
}
