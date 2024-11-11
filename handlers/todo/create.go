package todo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/playsoil/todo-go/database"
	"github.com/playsoil/todo-go/models"
)

func CreateHandler(c *fiber.Ctx) error {
	task := new(models.Task)

	if err := c.BodyParser(task); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "body must be a valid json"})
	}

	if len(task.Title) == 0 {
		return c.Status(422).JSON(fiber.Map{"title": "title field is required"})
	}
	database.DB.DB.Create(&task)

	return c.JSON(fiber.Map{"data": task})
}
