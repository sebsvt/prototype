package rest

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sebsvt/prototype/service"
)

type orderRestAPI struct {
	order_srv    service.OrderService
	user_service service.UserService
}

func NewOrderRestAPI(order_srv service.OrderService, user_service service.UserService) orderRestAPI {
	return orderRestAPI{order_srv: order_srv, user_service: user_service}
}

func (h orderRestAPI) FromOrderID(c *fiber.Ctx) error {
	order_id, err := c.ParamsInt("order_id")
	if err != nil {
		return err
	}
	headers := c.GetReqHeaders()
	access_token, has := headers["Authorization"]
	if !has {
		log.Println("not has")
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Authorization header is required",
		})
	}
	user_db, err := h.user_service.GetCurrentUser(access_token[0])
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": err,
		})
	}
	order, err := h.order_srv.GetOrderFromOrderID(order_id, user_db.UserID)
	if err != nil {
		return err
	}
	return c.JSON(order)
}

func (h orderRestAPI) GetAllOrders(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	access_token, has := headers["Authorization"]
	if !has {
		log.Println("not has")
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Authorization header is required",
		})
	}
	user_db, err := h.user_service.GetCurrentUser(access_token[0])
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": err,
		})
	}
	orders, err := h.order_srv.GetAllOrders(user_db.UserID)
	if err != nil {
		log.Println(err)
		return err
	}
	return c.JSON(orders)
}

func (h orderRestAPI) CreateNewOrder(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	access_token, has := headers["Authorization"]
	log.Println("access_tplem", access_token)
	if !has {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Authorization header is required",
		})
	}
	user_db, err := h.user_service.GetCurrentUser(access_token[0])
	if err != nil {
		log.Println("auth where is current user")
		return err
	}
	var new_order service.OrderRequest
	if err := c.BodyParser(&new_order); err != nil {
		log.Println("parser mai dai")
		return err
	}
	order, err := h.order_srv.CreateNewOrder(user_db.UserID, new_order)
	if err != nil {
		log.Println("create mai dai")
		return err
	}
	return c.JSON(order)
}

func (h orderRestAPI) CheckOut(c *fiber.Ctx) error {
	order_id, err := c.ParamsInt("order_id")
	if err != nil {
		return err
	}

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	err = h.order_srv.CheckOut(order_id, file)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Order checked out successfully"})
}
