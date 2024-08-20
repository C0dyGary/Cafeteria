package order

import "Cafeteria/src/pkg/domain"

func (s Service) GetOrders() ([]domain.Order, error) {
	return s.Repo.SelectOrders()
}
