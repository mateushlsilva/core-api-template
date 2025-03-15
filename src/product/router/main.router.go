package product_router

import (
	auth_middleware "github.com/cogniia/core-api-template/src/auth/middleware"
	product_handler "github.com/mateushlsilva/core-api-template/src/product/handler"
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	group := app.Group("/product")

	group.Get("/", product_handler.GetProducts)       
	group.Get("/:id", product_handler.GetProduct)     

	group.Post("/", auth_middleware.UserMiddleware, product_handler.CreateProduct)   
	group.Put("/:id", auth_middleware.UserMiddleware, product_handler.UpdateProduct) 
	group.Delete("/:id", auth_middleware.UserMiddleware, product_handler.DeleteProduct) 
}