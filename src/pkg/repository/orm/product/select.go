package product

import "Cafeteria/src/pkg/domain"

func (r Repository) SelectProducts() ([]domain.Product, error) {
	var products []domain.Product

	err := r.DB.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}
