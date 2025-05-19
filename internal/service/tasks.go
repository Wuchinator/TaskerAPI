package service

import (
	"github.com/Wuchinator/GoTasker/internal/models"
	"github.com/Wuchinator/GoTasker/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetUserTasks(userID int) ([]models.Task, error) {
    return s.repo.GetUserTasks(userID)
}

func (s *TaskService) Create(task *models.Task) error {
    return s.repo.Create(task) 
}

func (s *TaskService) GetByID(userID, taskID int) (*models.Task, error) {
    return s.repo.GetByID(userID, taskID)
}

func (s *TaskService) Update(task *models.Task) error {
	return s.repo.Update(task)
}

func (s *TaskService) Delete(userID, taskID int) error {
    return s.repo.Delete(userID, taskID)
}