package models

import (
	"time"
	"errors"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" validate:"required,min=3,max=100"`
	Description string    `json:"description" validate:"max=500"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID    int         `json:"user_id"`
}

func (t *Task) Validate() error {
	if len(t.Title) < 3 {
		return errors.New("title must be at least 3 characters")
	}
	if len(t.Title) > 100 {
		return errors.New("title must be less than 100 characters")
	}
	if len(t.Description) > 500 {
		return errors.New("description must be less than 500 characters")
	}
	return nil
}