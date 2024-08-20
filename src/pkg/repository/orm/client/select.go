package client

import "Cafeteria/src/pkg/domain"

func (r Repository) SelectClients() ([]domain.Client, error) {
	var clients []domain.Client

	err := r.DB.Find(&clients).Error
	if err != nil {
		return nil, err
	}

	return clients, nil
}
