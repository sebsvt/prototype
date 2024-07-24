package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebsvt/prototype/orchestra/services"
)

type authHandler struct {
	auth_srv services.AuthService
}

func NewAuthHandler(auth_srv services.AuthService) authHandler {
	return authHandler{auth_srv: auth_srv}
}

func (h authHandler) SignIn(c *fiber.Ctx) error {
	var user_credential services.SignInRequest
	if err := c.BodyParser(&user_credential); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	access_token, err := h.auth_srv.SignIn(user_credential)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(access_token)
}

func (h authHandler) SignUp(c *fiber.Ctx) error {
	var user_credential services.SignUpRequest
	if err := c.BodyParser(&user_credential); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	access_token, err := h.auth_srv.SignUp(user_credential)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(access_token)
}

func (h authHandler) VerityToken(c *fiber.Ctx) error {
	claims := c.Locals("claims")
	return c.JSON(claims)
}
