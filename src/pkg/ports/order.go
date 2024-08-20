package ports

import "Cafeteria/src/pkg/domain"

type OrderService interface {
	CreateOrder(order domain.Order) (domain.Order, error)
	GetOrders() ([]domain.Order, error)
	/*GetOrderById(id uint) (domain.Order, error)
	UpdateOrder(order domain.Order) (domain.Order, error)
	DeleteOrder(id uint) error*/
}

type OrderRepository interface {
	InsertOrder(order domain.Order) (domain.Order, error)
	SelectOrders() ([]domain.Order, error)
	GetProductByID(id uint, d *domain.Product) error
	GetClientByID(id uint, d *domain.Client) error
	/*SelectOrderById(id uint) (domain.Order, error)
	UpdateOrder(order domain.Order) (domain.Order, error)
	DeleteOrder(id uint) error*/
}
