package order

import "Cafeteria/src/pkg/ports"

type Handler struct {
	Service ports.OrderService
}
