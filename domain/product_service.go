package domain

import (
	"context"

	"github.com/umardev500/store/proto"
)

type ProductService interface {
	CreateProduct(ctx context.Context, p *proto.CreateProductRequest) error
}
