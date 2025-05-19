package repository

import (
	"database/sql"
	"log"

	"github.com/Wuchinator/GoTasker/internal/models"
)

type PostgresTaskRepository struct {
	db *sql.DB
}

func NewPostgresTaskRepository(db *sql.DB) TaskRepository {
	return &PostgresTaskRepository{db: db}
}

func (r *PostgresTaskRepository) GetAll() ([]models.Task, error) {
	rows, err := r.db.Query(`
		SELECT id, title, description, completed, created_at, updated_at, user_id
		FROM tasks
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		if err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.Completed,
			&t.CreatedAt,
			&t.UpdatedAt,
			&t.UserID,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *PostgresTaskRepository) GetUserTasks(userID int) ([]models.Task, error) {
	rows, err := r.db.Query(`
		SELECT id, title, description, completed, created_at, updated_at
		FROM tasks
		WHERE user_id = $1
		ORDER BY created_at DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		if err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.Completed,
			&t.CreatedAt,
			&t.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *PostgresTaskRepository) GetByID(userID, taskID int) (*models.Task, error) {
	var t models.Task
	err := r.db.QueryRow(`
		SELECT id, title, description, completed, created_at, updated_at
		FROM tasks
		WHERE id = $1 AND user_id = $2
	`, taskID, userID).Scan(
		&t.ID,
		&t.Title,
		&t.Description,
		&t.Completed,
		&t.CreatedAt,
		&t.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (r *PostgresTaskRepository) Create(task *models.Task) error {
    log.Printf("Creating task: %+v", task) // Добавьте это
    err := r.db.QueryRow(`
        INSERT INTO tasks (title, description, completed, user_id)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at, updated_at
    `, task.Title, task.Description, task.Completed, task.UserID).Scan(
        &task.ID,
        &task.CreatedAt,
        &task.UpdatedAt,
    )
    if err != nil {
        log.Printf("Create task error: %v", err) // И это
    }
    return err
}

func (r *PostgresTaskRepository) Update(task *models.Task) error {
	_, err := r.db.Exec(`
		UPDATE tasks
		SET title = $1, description = $2, completed = $3, updated_at = NOW()
		WHERE id = $4 AND user_id = $5
	`, task.Title, task.Description, task.Completed, task.ID, task.UserID)
	return err
}

func (r *PostgresTaskRepository) Delete(userID, taskID int) error {
	_, err := r.db.Exec(`
		DELETE FROM tasks
		WHERE id = $1 AND user_id = $2
	`, taskID, userID)
	return err
}