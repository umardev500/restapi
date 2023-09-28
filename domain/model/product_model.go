package model

type ProductStore struct {
	Name        string   `json:"name" validate:"required,min=15"`
	Description string   `json:"description" validate:"required,min=50"`
	Price       float32  `json:"price" validate:"required"`
	Stock       int64    `json:"stock" validate:"required"`
	Images      []string `json:"images" validate:"required"`
}
