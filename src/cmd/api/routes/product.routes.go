package routes

import (
	product3 "Cafeteria/src/cmd/api/handlers/product"
	"Cafeteria/src/pkg/repository/orm/product"
	product2 "Cafeteria/src/pkg/service/product"
	"github.com/gofiber/fiber/v2"
)

func (r Router) ProductRoutes(app fiber.Router) fiber.Router {
	repositoryProduct := product.Repository{
		DB: r.BaseData,
	}

	serviceProduct := product2.Service{
		Repo: repositoryProduct,
	}

	handlersProduct := product3.Handler{
		Service: serviceProduct,
	}

	app.Post("/product", handlersProduct.CreateProduct)
	app.Get("/product", handlersProduct.GetProducts)
	return app
}
