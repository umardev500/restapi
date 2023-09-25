package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/restapi/application/handler"
	"github.com/umardev500/restapi/config"
	"github.com/umardev500/store/proto"
)

func LoadAllRoutes(app *fiber.App) {
	app.Route("api", loadApiRoutes)
}

func loadApiRoutes(router fiber.Router) {
	router.Route("product", loadProductRoutes)
}

func loadProductRoutes(router fiber.Router) {
	conn := config.NewProductRPC()
	product := proto.NewProductServiceClient(conn)
	handler := handler.NewProductHandler(product)
	router.Post("/", handler.Create)
}
