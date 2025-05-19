package repository

import (
	"sync"
	
	"github.com/Wuchinator/GoTasker/internal/models"
)

type memoryTaskRepository struct {
	mu    sync.RWMutex
	tasks map[int]models.Task // key: task ID
	nextID int
}

func NewMemoryTaskRepository() TaskRepository {
	return &memoryTaskRepository{
		tasks:  make(map[int]models.Task),
		nextID: 1,
	}
}

func (r *memoryTaskRepository) GetAll() ([]models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]models.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *memoryTaskRepository) GetUserTasks(userID int) ([]models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var tasks []models.Task
	for _, task := range r.tasks {
		if task.UserID == userID {
			tasks = append(tasks, task)
		}
	}
	return tasks, nil
}

func (r *memoryTaskRepository) GetByID(userID, taskID int) (*models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[taskID]
	if !exists || task.UserID != userID {
		return nil, nil
	}
	return &task, nil
}

func (r *memoryTaskRepository) Create(task *models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	task.ID = r.nextID
	r.tasks[r.nextID] = *task
	r.nextID++
	return nil
}

func (r *memoryTaskRepository) Update(task *models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[task.ID]; !exists {
		return nil // или возвращайте ошибку "не найдено"
	}

	r.tasks[task.ID] = *task
	return nil
}

func (r *memoryTaskRepository) Delete(userID, taskID int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if task, exists := r.tasks[taskID]; exists && task.UserID == userID {
		delete(r.tasks, taskID)
	}
	return nil
}