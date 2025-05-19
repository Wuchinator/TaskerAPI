package repository

import "github.com/Wuchinator/GoTasker/internal/models"

type TaskRepository interface {
	GetAll() ([]models.Task, error)
	GetUserTasks(userID int) ([]models.Task, error)
	GetByID(userID, taskID int) (*models.Task, error)
	Create(task *models.Task) error
	Update(task *models.Task) error
	Delete(userID, taskID int) error
}