package order

import (
	"Cafeteria/src/pkg/domain"
	"github.com/gofiber/fiber/v2"
)

func (h Handler) CreateOrder(c *fiber.Ctx) error {
	var order domain.Order
	if err := c.BodyParser(order); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}
	order, err := h.Service.CreateOrder(order)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "Order created successfully",
		"order":   order,
	})
}
