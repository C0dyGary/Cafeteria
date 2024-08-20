package ports

import "Cafeteria/src/pkg/domain"

type OrderDetailService interface {
	CreateOrderDetail(orderDetail domain.OrderDetail) (domain.OrderDetail, error)
	GetOrderDetails() ([]domain.OrderDetail, error)
}

type OrderDetailRepository interface {
	InsertOrderDetail(orderDetail domain.OrderDetail) (domain.OrderDetail, error)
	SelectOrderDetails() ([]domain.OrderDetail, error)
}
