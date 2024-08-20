package order

import "Cafeteria/src/pkg/domain"

func (r Repository) SelectOrders() ([]domain.Order, error) {
	var orders []domain.Order
	err := r.DB.Preload("OrderDetail").Preload("OrderDetail.Product").Find(&orders).Error
	if err != nil {
		return []domain.Order{}, err
	}
	return orders, nil
}
