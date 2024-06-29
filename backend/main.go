package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sebsvt/prototype/repository"
	"github.com/sebsvt/prototype/rest"
	"github.com/sebsvt/prototype/service"
)

func main() {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/prototype?parseTime=true")
	if err != nil {
		panic(err)
	}
	app := fiber.New()
	node_repo := repository.NewNodeRepositoryDB(db)
	node_serv := service.NewNodeService(node_repo)
	node_rest := rest.NewNodeRestAPI(node_serv)
	order_repo := repository.NewOrderRepositoryDB(db)
	order_serv := service.NewOrderService(order_repo)
	order_rest := rest.NewOrderRestAPI(order_serv)

	api := app.Group("/api")

	api.Get("nodes/:node_id", node_rest.GetNodeFromNodeID)
	api.Get("nodes/", node_rest.GetAllNodes)
	api.Post("nodes/create", node_rest.CreateNewNode)

	api.Get("order/:order_id", order_rest.FromOrderID)
	api.Get("order/", order_rest.GetAllOrders)
	api.Post("order/create", order_rest.CreateNewOrder)

	api.Post("payments/", func(c *fiber.Ctx) error {
		return nil
	})
	app.Listen(":8000")
}
