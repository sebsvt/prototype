package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	ordering_adapters "github.com/sebsvt/prototype/adapters/ordering"
	"github.com/sebsvt/prototype/rest"
	ordering_service "github.com/sebsvt/prototype/services/ordering"
)

func main() {
	mongodb_adapter, err := ordering_adapters.NewOrderRepositoryMongoDB(context.Background(), "mongodb://root:admin@localhost:27017")
	if err != nil {
		log.Panic(err)
	}
	srv := ordering_service.NewOrderService(mongodb_adapter)
	api_router := rest.NewOrderRest(srv)
	app := fiber.New()
	app.Get("/:ref", api_router.GetOrderFromReference)
	app.Post("/order", api_router.CreateNewOrder)
	app.Listen(":8000")
}
