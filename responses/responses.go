package responses

import "github.com/gofiber/fiber/v2"

func Error(c *fiber.Ctx, code int, message string) error {
	return c.
		Status(code).
		JSON(fiber.Map{"message": "could not create task", "error": message})
}
