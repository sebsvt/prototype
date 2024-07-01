package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
	"github.com/sebsvt/prototype/repository"
	"github.com/sebsvt/prototype/rest"
	"github.com/sebsvt/prototype/service"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := "root"
	dbPassword := os.Getenv("MARIADB_ROOT_PASSWORD")
	dbName := os.Getenv("MARIADB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	app := fiber.New()
	app.Use(cors.New())

	node_repo := repository.NewNodeRepositoryDB(db)
	node_serv := service.NewNodeService(node_repo)
	node_rest := rest.NewNodeRestAPI(node_serv)

	payment_repo := repository.NewPaymentRepositoryDB(db)
	payment_serv := service.NewPaymentService(payment_repo)
	payment_rest := rest.NewPaymentRestAPI(payment_serv)

	order_repo := repository.NewOrderRepositoryDB(db)
	order_serv := service.NewOrderService(order_repo, payment_serv)
	order_rest := rest.NewOrderRestAPI(order_serv)

	product_repo := repository.NewProductRepositoryDB(db)
	product_serv := service.NewProductService(product_repo)
	product_rest := rest.NewProductRestAPI(product_serv)

	api := app.Group("/api")

	api.Get("nodes/:node_id", node_rest.GetNodeFromNodeID)
	api.Get("nodes/", node_rest.GetAllNodes)
	api.Post("nodes/create", node_rest.CreateNewNode)

	api.Get("order/:order_id", order_rest.FromOrderID)
	api.Get("order/", order_rest.GetAllOrders)
	api.Post("order/create", order_rest.CreateNewOrder)

	api.Post("payments/", payment_rest.CreatePayment)
	api.Get("payments/:payment_id", payment_rest.GetPaymentFromID)

	api.Get("products/", product_rest.GetAllProducts)
	api.Get("products/:product_id", product_rest.GetProductFromID)
	api.Post("products/", product_rest.CreateNewProduct)

	app.Listen(":8000")
}
