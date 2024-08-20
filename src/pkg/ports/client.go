package ports

import (
	"Cafeteria/src/pkg/adapters"
	"Cafeteria/src/pkg/domain"
)

type ClientService interface {
	CreateClient(client adapters.ClientParams) (domain.Client, error)
	GetClients() ([]domain.Client, error)
	/*GetClient(id string) (domain.Client, error)

	UpdateClient(id string, client adapters.ClientParams) (domain.Client, error)
	DeleteClient(id string) error*/
}

type ClientRepository interface {
	InsertClient(client adapters.ClientParams) (domain.Client, error)
	SelectClients() ([]domain.Client, error)
	/*SelectClient(id string) (domain.Client, error)

	ModifyClient(id string, client adapters.ClientParams) (domain.Client, error)
	RemoveClient(id string) error*/
}
