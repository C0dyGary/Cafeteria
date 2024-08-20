package client

import "Cafeteria/src/pkg/domain"

func (s Service) GetClients() ([]domain.Client, error) {
	clients, err := s.Repo.SelectClients()
	if err != nil {
		return nil, err
	}
	return clients, nil
}
