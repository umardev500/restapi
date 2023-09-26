package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/restapi/domain"
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := ph.service.CreateProduct(ctx, &proto.CreateProductRequest{
		Product: &proto.ProductCreate{
			Name:        "Car",
			Description: "Lorem ipsum",
			Price:       "5000",
			ImageUrls:   []string{"1"},
		},
	})
	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON("ok")
}
