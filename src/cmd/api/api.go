package api

import (
	"Cafeteria/src/cmd/api/routes"
	"Cafeteria/src/pkg/domain"
	"Cafeteria/src/pkg/repository/orm"
	"github.com/gofiber/fiber/v2"
)

func StartApi() {
	app := fiber.New()

	connect := orm.Connect()

	orm.Migrate([]interface{}{
		&domain.Client{},
		&domain.Order{},
		&domain.Product{},
		&domain.OrderDetail{},
	}, connect)

	router := routes.Router{
		BaseData: connect,
	}

	api := app.Group("/api/v1")
	router.ClientRoutes(api)
	router.ProductRoutes(api)
	router.OrderRoutes(api)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err := app.Listen("0.0.0.0:3000")
	if err != nil {
		panic(err)
	}
}
