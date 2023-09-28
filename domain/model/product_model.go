package model

type ProductStore struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Price       float32  `json:"price" validate:"required"`
	Stock       int64    `json:"stock" validate:"required"`
	Images      []string `json:"images" validate:"required"`
}
