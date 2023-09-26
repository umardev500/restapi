package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/restapi/domain"
	"github.com/umardev500/restapi/domain/model"
)

type productHandler struct {
	service domain.ProductService
}

func NewProductHandler(service domain.ProductService) domain.ProductHanlder {
	return &productHandler{
		service: service,
	}
}

func (ph *productHandler) Create(c *fiber.Ctx) error {
	var product *model.ProductStore = new(model.ProductStore)
	if err := c.BodyParser(product); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := ph.service.CreateProduct(ctx, *product)
	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON("ok")
}
