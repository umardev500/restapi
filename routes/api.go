package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/restapi/application/handler"
)

func LoadAllRoutes(app *fiber.App) {
	app.Route("api", loadApiRoutes)
}

func loadApiRoutes(router fiber.Router) {
	router.Route("product", loadProductRoutes)
}

func loadProductRoutes(router fiber.Router) {
	handler := handler.NewProductHandler()
	router.Post("/", handler.Create)
}
