package repositories

import (
	"github.com/playsoil/todo-go/database"
	"github.com/playsoil/todo-go/models"
)

func GetTaskList() ([]models.Task, error) {
	var tasks []models.Task
	if result := database.DB.DB.Find(&tasks); result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func CreateTask(task *models.Task) (*models.Task, error) {

	if err := database.DB.DB.Create(&task).Error; err != nil {
		return nil, err // return the error if database operation fails
	}
	return task, nil
}
