package todo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/playsoil/todo-go/database/repositories"
	"github.com/playsoil/todo-go/responses"
)

func ListHandler(c *fiber.Ctx) error {
	tasks, err := repositories.GetTaskList()

	if err != nil {
		return responses.Error(c, 500, "internal server error, "+err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": tasks,
	})
}
