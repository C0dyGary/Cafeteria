package client

import (
	"Cafeteria/src/pkg/adapters"
	"Cafeteria/src/pkg/domain"
)

func (r Repository) InsertClient(client adapters.ClientParams) (domain.Client, error) {
	clientDomain := domain.Client{
		Name:  client.Name,
		Phone: client.Phone,
	}
	err := r.DB.Create(&clientDomain).Error
	if err != nil {
		return domain.Client{}, err
	}
	return clientDomain, nil
}
