package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebsvt/prototype/service"
)

type productRestAPI struct {
	product_srv service.ProductService
}

func NewProductRestAPI(product_srv service.ProductService) productRestAPI {
	return productRestAPI{product_srv: product_srv}
}

func (h productRestAPI) CreateNewProduct(c *fiber.Ctx) error {
	var product service.ProductRequest
	if err := c.BodyParser(&product); err != nil {
		return err
	}
	product_id, err := h.product_srv.CreateNewProduct(product)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"product_id": product_id,
	})
}

func (h productRestAPI) GetProductFromID(c *fiber.Ctx) error {
	product_id, err := c.ParamsInt("product_id")
	if err != nil {
		return err
	}
	product, err := h.product_srv.GetProductByID(product_id)
	if err != nil {
		return err
	}
	return c.JSON(product)
}

func (h productRestAPI) GetAllProducts(c *fiber.Ctx) error {
	products, err := h.product_srv.GetAllProducts()
	if err != nil {
		return err
	}
	return c.JSON(products)
}
