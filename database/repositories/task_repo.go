package repositories

import (
	"github.com/playsoil/todo-go/database"
	"github.com/playsoil/todo-go/models"
)

func GetTaskList() []models.Task {
	var tasks []models.Task
	database.DB.DB.Find(&tasks)
	return tasks
}

func CreateTask(task *models.Task) (*models.Task, error) {

	if err := database.DB.DB.Create(&task).Error; err != nil {
		return nil, err // return the error if database operation fails
	}
	return task, nil
}
