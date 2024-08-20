package routes

import (
	client3 "Cafeteria/src/cmd/api/handlers/client"
	"Cafeteria/src/pkg/repository/orm/client"
	client2 "Cafeteria/src/pkg/service/client"
	"github.com/gofiber/fiber/v2"
)

func (r Router) ClientRoutes(app fiber.Router) fiber.Router {

	repoClient := client.Repository{DB: r.BaseData}

	serviceClient := client2.Service{Repo: repoClient}

	handlerClient := client3.Handler{Service: serviceClient}

	app.Post("/client", handlerClient.CreateClient)
	app.Get("/client", handlerClient.GetClients)

	return app
}
