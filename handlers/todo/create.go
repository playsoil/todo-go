package todo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/playsoil/todo-go/database/repositories"
	"github.com/playsoil/todo-go/models"
)

func CreateHandler(c *fiber.Ctx) error {
	task := new(models.Task)

	if err := c.BodyParser(task); err != nil {
		return ErrorResponse(c, 400, "request body must be a valid json")
	}

	if len(task.Title) == 0 {
		return ErrorResponse(c, 422, "title field is required")
	}

	createdTask, err := repositories.CreateTask(task)
	if err != nil {
		return ErrorResponse(c, 500, "internal server error, "+err.Error())
	}

	return c.JSON(fiber.Map{"data": createdTask})
}

func ErrorResponse(c *fiber.Ctx, code int, message string) error {
	return c.
		Status(code).
		JSON(fiber.Map{"message": "could not create task", "error": message})
}
