package todo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/playsoil/todo-go/database"
	"github.com/playsoil/todo-go/models"
)

func ListHandler(c *fiber.Ctx) error {
	tasks := []models.Task{}
	database.DB.DB.Find(&tasks)
	return c.Status(200).JSON(fiber.Map{
		"data": tasks,
	})
}
