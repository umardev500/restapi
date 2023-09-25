package domain

import "github.com/gofiber/fiber/v2"

type ProductHanlder interface {
	Create(c *fiber.Ctx) error
}
