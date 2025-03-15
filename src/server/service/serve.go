package server_service

import (
	"fmt"

	"github.com/cogniia/core-api-template/src/config/env"
	user_router "github.com/cogniia/core-api-template/src/user/router"
	product_router "github.com/mateushlsilva/core-api-template/src/product/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pterm/pterm"
)

func Serve() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	makeDocs(app)
	user_router.Route(app)
	product_router.Route(app)

	err := app.Listen(fmt.Sprintf(":%s", env.ServerPort))
	pterm.DefaultLogger.Fatal(
		fmt.Sprintf("%v", err),
	)
}
