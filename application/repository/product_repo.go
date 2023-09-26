package repository

import (
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
