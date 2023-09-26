package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/restapi/domain"
	"github.com/umardev500/restapi/domain/model"
	"github.com/umardev500/store/proto"
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
	err := ph.service.CreateProduct(ctx, &proto.CreateProductRequest{
		Product: &proto.ProductCreate{
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			ImageUrls:   product.Images,
		},
	})
	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON("ok")
}
