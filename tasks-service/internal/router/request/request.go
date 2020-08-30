package request

import (
	"time"
)

// CreateTaskRequestBody ...
type CreateTaskRequestBody struct {
	Summary string    `json:"summary" validate:"required"`
	Date    time.Time `json:"date" validate:"required"`
}

// UpdateTaskRequestBody ...
type UpdateTaskRequestBody struct {
	Summary string    `json:"summary"`
	Date    time.Time `json:"date"`
}
