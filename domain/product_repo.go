package domain

import (
	"context"

	"github.com/umardev500/store/proto"
)

type ProductRepository interface {
	Create(ctx context.Context, p *proto.CreateProductRequest) error
}
