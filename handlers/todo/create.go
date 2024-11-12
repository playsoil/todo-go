package todo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/playsoil/todo-go/database/repositories"
	"github.com/playsoil/todo-go/models"
	"github.com/playsoil/todo-go/responses"
)

func CreateHandler(c *fiber.Ctx) error {
	task := new(models.Task)

	if err := c.BodyParser(task); err != nil {
		return responses.Error(c, 400, "request body must be a valid json")
	}

	if len(task.Title) == 0 {
		return responses.Error(c, 422, "title field is required")
	}

	createdTask, err := repositories.CreateTask(task)
	if err != nil {
		return responses.Error(c, 500, "internal server error, "+err.Error())
	}

	return responses.Success(c, createdTask)
}
