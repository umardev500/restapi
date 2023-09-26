package service

import "github.com/umardev500/restapi/domain"

type productService struct {
	repo domain.ProductRepository
}

func NewProductService(repo domain.ProductService) domain.ProductService {
	return &productService{
		repo: repo,
	}
}
