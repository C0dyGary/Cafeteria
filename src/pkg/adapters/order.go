package adapters

import "Cafeteria/src/pkg/domain"

type OrderParams struct {
	ClientID    uint                 `json:"client_id"`
	OrderDetail []domain.OrderDetail `json:"order_detail"`
}
