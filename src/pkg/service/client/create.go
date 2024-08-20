package client

import (
	"Cafeteria/src/pkg/adapters"
	"Cafeteria/src/pkg/domain"
)

func (s Service) CreateClient(client adapters.ClientParams) (domain.Client, error) {
	clientDomain, err := s.Repo.InsertClient(client)
	if err != nil {
		return domain.Client{}, err
	}
	return clientDomain, nil
}
