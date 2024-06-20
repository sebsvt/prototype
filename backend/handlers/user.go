package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebsvt/prototype/services"
)

type userHandler struct {
	userSrv services.UserService
}

func NewUserHandler(userSrv services.UserService) userHandler {
	return userHandler{userSrv: userSrv}
}

func (h userHandler) CreateNewUserAccount(c *fiber.Ctx) error {
	var entity services.CreateNewUserRequest
	if err := c.BodyParser(&entity); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	res, err := h.userSrv.CreateUserAccount(entity)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h userHandler) GetUserFromID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	user, err := h.userSrv.GetUser(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(user)
}

type LoginSchme struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h userHandler) Login(c *fiber.Ctx) error {
	var credentials LoginSchme
	if err := c.BodyParser(&credentials); err != nil {
		return err
	}

	accessToken, err := h.userSrv.SignIn(credentials.Email, credentials.Password)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"access_token": accessToken,
		"type":         "Bearer",
	})
}
