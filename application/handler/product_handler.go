package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/restapi/domain"
)

type productHandler struct{}

func NewProductHandler() domain.ProductHanlder {
	return &productHandler{}
}

func (ph *productHandler) Create(c *fiber.Ctx) error {
	return nil
}
