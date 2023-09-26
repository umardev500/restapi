package service

import (
	"context"

	"github.com/umardev500/restapi/domain"
	"github.com/umardev500/restapi/domain/model"
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

func (ps *productService) CreateProduct(ctx context.Context, product model.ProductStore) error {
	var p *proto.CreateProductRequest = &proto.CreateProductRequest{
		Product: &proto.ProductCreate{
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			ImageUrls:   product.Images,
		},
	}

	err := ps.repo.Create(ctx, p)
	if err != nil {
		return err
	}
	return nil
}
