package errors

import (
	"encoding/json"
	"net/http"
)

type AppError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Respond(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.StatusCode)
	json.NewEncoder(w).Encode(e)
}

func NewBadRequestError(message string) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: http.StatusBadRequest,
	}
}

