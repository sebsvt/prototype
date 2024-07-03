package rest

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sebsvt/prototype/service"
)

type userRestAPI struct {
	user_srv service.UserService
}

func NewUserRestAPI(user_srv service.UserService) userRestAPI {
	return userRestAPI{user_srv: user_srv}
}

func (h userRestAPI) Register(c *fiber.Ctx) error {
	var new_user service.UserRequest
	if err := c.BodyParser(&new_user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	user_id, err := h.user_srv.CreateNewUser(new_user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{
		"user_id": user_id,
	})
}

func (h userRestAPI) Login(c *fiber.Ctx) error {
	var user_crendetial service.UserLogin
	if err := c.BodyParser(&user_crendetial); err != nil {
		log.Println(err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": err.Error()})
	}

	access_token, err := h.user_srv.SignIn(user_crendetial)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"access_token": access_token, "type": "Bearer"})
}

func (h userRestAPI) GetCurrentUser(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	access_token, has := headers["Authorization"]
	if !has {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.New("authorization is required")})
	}
	user, err := h.user_srv.GetCurrentUser(access_token[0])
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}

func (h userRestAPI) GetUserFromID(c *fiber.Ctx) error {
	user_id, err := c.ParamsInt("user_id")
	if err != nil {
		return err
	}
	user, err := h.user_srv.GetUserFromID(user_id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}
