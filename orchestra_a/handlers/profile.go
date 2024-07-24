package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebsvt/prototype/orchestra/services"
)

type profileHandler struct {
	profile_srv services.ProfileService
}

func NewProfileHandler(profile_srv services.ProfileService) profileHandler {
	return profileHandler{profile_srv: profile_srv}
}

func (h profileHandler) GetUserProfileFromUserID(c *fiber.Ctx) error {
	user_id, err := c.ParamsInt("user_id")
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	profile, err := h.profile_srv.GetProfileFromUserID(user_id)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(profile)
}
