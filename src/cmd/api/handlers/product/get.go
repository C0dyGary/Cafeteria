package product

import "github.com/gofiber/fiber/v2"

func (h Handler) GetProducts(c *fiber.Ctx) error {
	products, err := h.Service.GetProducts()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"products": products,
	})
}
