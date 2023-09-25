package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/restapi/domain"
	"github.com/umardev500/store/proto"
)

type productHandler struct {
	product proto.ProductServiceClient
}

func NewProductHandler(product proto.ProductServiceClient) domain.ProductHanlder {
	return &productHandler{
		product: product,
	}
}

func (ph *productHandler) Create(c *fiber.Ctx) error {
	return nil
}
