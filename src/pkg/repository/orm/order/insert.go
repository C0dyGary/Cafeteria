package order

import "Cafeteria/src/pkg/domain"

func (r Repository) InsertOrder(order domain.Order) (domain.Order, error) {
	err := r.DB.Create(&order).Error
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}
