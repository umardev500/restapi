package model

type ProductStore struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float32  `json:"price"`
	Images      []string `json:"images"`
}
