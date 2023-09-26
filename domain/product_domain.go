package domain

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/store/proto"
)

type ProductHanlder interface {
	Create(c *fiber.Ctx) error
}

type ProductService interface {
	CreateProduct(ctx context.Context, p *proto.CreateProductRequest) error
}

type ProductRepository interface {
	Create(ctx context.Context, p *proto.CreateProductRequest) error
}
