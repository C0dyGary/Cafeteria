package product

import "Cafeteria/src/pkg/domain"

func (s Service) GetProducts() ([]domain.Product, error) {
	products, err := s.Repo.SelectProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}
