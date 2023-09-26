package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/restapi/domain"
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
	return c.SendStatus(fiber.StatusOK)
}
