package product

import (
	"Cafeteria/src/pkg/adapters"
	"github.com/gofiber/fiber/v2"
)

func (h Handler) CreateProduct(c *fiber.Ctx) error {
	productParam := adapters.ProductParams{}

	if err := c.BodyParser(&productParam); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	product, err := h.Service.CreateProduct(productParam)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"product": product,
	})

}
