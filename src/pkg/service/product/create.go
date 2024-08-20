package product

import (
	"Cafeteria/src/pkg/adapters"
	"Cafeteria/src/pkg/domain"
)

func (s Service) CreateProduct(product adapters.ProductParams) (domain.Product, error) {
	productDomain, err := s.Repo.InsertProduct(product)
	if err != nil {
		return domain.Product{}, err
	}

	return productDomain, nil
}
