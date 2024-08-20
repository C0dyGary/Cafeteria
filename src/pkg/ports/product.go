package ports

import (
	"Cafeteria/src/pkg/adapters"
	"Cafeteria/src/pkg/domain"
)

type ProductService interface {
	CreateProduct(product adapters.ProductParams) (domain.Product, error)
	GetProducts() ([]domain.Product, error)
}

type ProductRepository interface {
	InsertProduct(product adapters.ProductParams) (domain.Product, error)
	SelectProducts() ([]domain.Product, error)
}
