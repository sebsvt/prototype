package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sebsvt/prototype/orchestra/adapters"
	"github.com/sebsvt/prototype/orchestra/configs"
	"github.com/sebsvt/prototype/orchestra/handlers"
	"github.com/sebsvt/prototype/orchestra/services"
	"github.com/sebsvt/prototype/orchestra/utils"
)

func main() {
	configs.InitConfig()
	initTimeZone()
	db := initDB()
	defer db.Close()

	profile_repo := adapters.NewProfileRepositoryPSQLDB(db)
	profile_srv := services.NewProfileService(profile_repo)
	profile_handler := handlers.NewProfileHandler(profile_srv)

	user_repo := adapters.NewUserRepositoryPSQLDB(db)
	auth_srv := services.NewAuthService(user_repo, profile_srv)
	auth_handler := handlers.NewAuthHandler(auth_srv)

	studio_repo := adapters.NewStudioRepositoryPSQLDB(db)
	membership_repo := adapters.NewMembershipRepositoryPSQLDB(db)
	studio_srv := services.NewStudioService(studio_repo, membership_repo)
	studio_handler := handlers.NewStudioHandler(studio_srv)

	app := fiber.New()

	app.Use(cors.New(cors.ConfigDefault))
	app.Use("/api/auth/verify_token", utils.AuthRequired())
	app.Use("/api/studio/opening", utils.AuthRequired())
	app.Use("/api/studio/my-studio", utils.AuthRequired())

	api := app.Group("/api")

	auth_route := api.Group("/auth")
	profile_route := api.Group("/profile")
	studio_route := api.Group("/studio")

	auth_route.Post("/sign-in", auth_handler.SignIn)
	auth_route.Post("/sign-up", auth_handler.SignUp)
	auth_route.Get("/verify_token", auth_handler.VerityToken)

	profile_route.Get("/:user_id", profile_handler.GetUserProfileFromUserID)

	studio_route.Post("/opening", studio_handler.OpenNewStudio)
	studio_route.Get("/my-studio", studio_handler.MyStudios)
	studio_route.Get("/:subdomain", studio_handler.GetStudioFromSubDomain)

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
