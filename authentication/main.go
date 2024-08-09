package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sebsvt/prototype/authentication/adapters"
	"github.com/sebsvt/prototype/authentication/configs"
	"github.com/sebsvt/prototype/authentication/handlers"
	"github.com/sebsvt/prototype/authentication/services"
	"github.com/sebsvt/prototype/authentication/utils"
)

func main() {
	configs.InitConfig()
	initTimeZone()
	db := initDB()
	defer db.Close()

	user_repo := adapters.NewUserRepositoryPSQLDB(db)
	auth_srv := services.NewAuthService(user_repo)
	auth_handler := handlers.NewAuthHandler(auth_srv)

	app := fiber.New()

	app.Use(cors.New(cors.ConfigDefault))
	app.Use("/api/auth/verify_token", utils.AuthRequired())

	api := app.Group("/api")

	auth_route := api.Group("/auth")

	auth_route.Post("/sign-in", auth_handler.SignIn)
	auth_route.Post("/sign-up", auth_handler.SignUp)
	auth_route.Get("/verify_token", auth_handler.VerityToken)

	app.Listen(":8000")
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func initDB() *sqlx.DB {
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
		configs.EnvConfig.DATABASE_USERNAME,
		configs.EnvConfig.DATABASE_PASSWORD,
		configs.EnvConfig.DATABASE_HOST,
		configs.EnvConfig.DATABASE_PORT,
		configs.EnvConfig.DATABASE_NAME,
		configs.EnvConfig.DATABASE_SSL_MODE,
	)
	fmt.Println(dsn)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
