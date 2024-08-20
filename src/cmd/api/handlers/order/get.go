package order

import "github.com/gofiber/fiber/v2"

func (h Handler) GetOrders(c *fiber.Ctx) error {
	orders, err := h.Service.GetOrders()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "failed to get orders",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "orders retrieved successfully",
		"orders":  orders,
	})
}
