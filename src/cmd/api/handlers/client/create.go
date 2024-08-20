package client

import (
	"Cafeteria/src/pkg/adapters"
	"github.com/gofiber/fiber/v2"
)

func (h Handler) CreateClient(c *fiber.Ctx) error {
	client := adapters.ClientParams{}
	if err := c.BodyParser(&client); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
		})
	}

	clientDomain, err := h.Service.CreateClient(client)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to create client",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": "Client created",
		"data":    clientDomain,
	})

}
