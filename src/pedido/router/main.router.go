package pedido_router

import (
	auth_middleware "github.com/cogniia/core-api-template/src/auth/middleware"
	pedido_handler "github.com/cogniia/core-api-template/src/pedido/handler"
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	group := app.Group("/pedido")

	mainRoutes(group)
	authRoutes(group)
}

func mainRoutes(group fiber.Router) {
	group.Get("/me", auth_middleware.UserMiddleware, pedido_handler.GetCurrentPedido)

	group.Post("/", pedido_handler.CreatePedido)
}
