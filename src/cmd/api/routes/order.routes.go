package routes

import (
	order3 "Cafeteria/src/cmd/api/handlers/order"
	"Cafeteria/src/pkg/repository/orm/order"
	order2 "Cafeteria/src/pkg/service/order"
	"github.com/gofiber/fiber/v2"
)

func (r Router) OrderRoutes(app fiber.Router) fiber.Router {
	repoOrder := order.Repository{
		DB: r.BaseData,
	}

	serviceOrder := order2.Service{
		Repo: repoOrder,
	}

	handlerOrder := order3.Handler{
		Service: serviceOrder,
	}

	app.Post("/order", handlerOrder.CreateOrder)
	app.Get("/order", handlerOrder.GetOrders)

	return app

}
