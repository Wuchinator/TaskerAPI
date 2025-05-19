package models

import "time"

type User struct {
    ID           int       `json:"id"`
    Email        string    `json:"email" validate:"required,email"`
    PasswordHash string    `json:"-"`
    CreatedAt    time.Time `json:"created_at"`
}