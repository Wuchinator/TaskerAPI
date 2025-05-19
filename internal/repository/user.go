package repository

import (
    "database/sql"
    "github.com/Wuchinator/GoTasker/internal/models"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
    err := r.db.QueryRow(
        "INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id, created_at",
        user.Email, user.PasswordHash,
    ).Scan(&user.ID, &user.CreatedAt)
    return err
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
    var user models.User
    err := r.db.QueryRow(
        "SELECT id, email, password_hash, created_at FROM users WHERE email = $1",
        email,
    ).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt)
    
    if err == sql.ErrNoRows {
        return nil, nil
    }
    return &user, err
}