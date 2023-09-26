package service

import (
	"context"

	"github.com/umardev500/restapi/domain"
	"github.com/umardev500/store/proto"
)

type productService struct {
	repo domain.ProductRepository
}

func NewProductService(repo domain.ProductRepository) domain.ProductService {
	return &productService{
		repo: repo,
	}
}

func (ps *productService) CreateProduct(ctx context.Context, p *proto.CreateProductRequest) error {
	err := ps.repo.Create(ctx, p)
	if err != nil {
		return err
	}
	return nil
}
