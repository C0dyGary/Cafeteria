package order

import (
	"Cafeteria/src/pkg/domain"
	"errors"
)

func (s Service) CreateOrder(order domain.Order) (domain.Order, error) {
	var client domain.Client
	err := s.Repo.GetClientByID(order.ClientID, &client)
	if err != nil {
		return domain.Order{}, err
	}
	if client.ID == 0 {
		return domain.Order{}, errors.New("client not found")
	}
	order.Client = client
	total := 0.0
	for i, detail := range order.OrderDetail {
		var product domain.Product
		err := s.Repo.GetProductByID(detail.ProductID, &product)
		if err != nil {
			return domain.Order{}, err
		}
		if product.ID == 0 {
			return domain.Order{}, errors.New("product not found")
		}

		order.OrderDetail[i].UnitPrice = product.Price
		order.OrderDetail[i].Total = product.Price * float64(detail.Quantity)
		total += order.OrderDetail[i].Total
	}

	return s.Repo.InsertOrder(order)
}
