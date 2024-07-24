package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sebsvt/prototype/orchestra/services"
	"github.com/sebsvt/prototype/orchestra/utils"
)

type studioHandler struct {
	studio_srv services.StudioService
}

func NewStudioHandler(studio_srv services.StudioService) studioHandler {
	return studioHandler{studio_srv: studio_srv}
}

func (h studioHandler) OpenNewStudio(c *fiber.Ctx) error {
	claims, has := c.Locals("claims").(*utils.Claims)
	if !has || claims == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	var new_studio services.StudioRequest
	if err := c.BodyParser(&new_studio); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	user_id, err := strconv.Atoi(claims.Issuer)
	if err != nil {
		return err
	}
	studio_id, err := h.studio_srv.CreateNewStudio(new_studio, user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"studio_id": studio_id,
	})
}

func (h studioHandler) GetStudioFromSubDomain(c *fiber.Ctx) error {
	subdomain := c.Params("subdomain")
	studio, err := h.studio_srv.GetStudioBySubDomain(subdomain)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(studio)
}

func (h studioHandler) MyStudios(c *fiber.Ctx) error {
	claims, has := c.Locals("claims").(*utils.Claims)
	if !has || claims == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}
	user_id, err := strconv.Atoi(claims.Issuer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	my_studios, err := h.studio_srv.GetAllStudiosFromUser(user_id)
	if err != nil {
		return err
	}
	return c.JSON(my_studios)
}
