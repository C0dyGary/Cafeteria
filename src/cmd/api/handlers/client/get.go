package client

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func (h Handler) GetClients(c *fiber.Ctx) error {
	clients, err := h.Service.GetClients()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errors.New("error getting clients"),
		})
	}
	return c.JSON(clients)
}
