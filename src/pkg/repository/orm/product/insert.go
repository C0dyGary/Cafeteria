package product

import (
	"Cafeteria/src/pkg/adapters"
	"Cafeteria/src/pkg/domain"
)

func (r Repository) InsertProduct(product adapters.ProductParams) (domain.Product, error) {
	productDomain := domain.Product{
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Category:    product.Category,
	}

	err := r.DB.Create(&productDomain).Error
	if err != nil {
		return domain.Product{}, err
	}

	return productDomain, nil
}
