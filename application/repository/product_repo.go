package repository

import (
	"context"

	"github.com/umardev500/restapi/domain"
	"github.com/umardev500/store/proto"
)

type productRepository struct {
	product proto.ProductServiceClient
}

func NewProductRepository(product proto.ProductServiceClient) domain.ProductRepository {
	return &productRepository{
		product: product,
	}
}

func (pr *productRepository) Create(ctx context.Context, p *proto.CreateProductRequest) error {
	_, err := pr.product.CreateProduct(ctx, p)
	if err != nil {
		return err
	}
	return nil
}
