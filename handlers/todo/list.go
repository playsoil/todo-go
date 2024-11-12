package todo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/playsoil/todo-go/database/repositories"
)

func ListHandler(c *fiber.Ctx) error {
	tasks := repositories.GetTaskList()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": tasks,
	})
}
