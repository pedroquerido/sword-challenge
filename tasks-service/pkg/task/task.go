package task

import (
	"time"

	"github.com/google/uuid"
)

// Task ...
type Task struct {
	ID      string    `json:"id" validate:"required,gt=0,uuid4"`
	UserID  string    `json:"user_id" validate:"required,gt=0"`
	Summary string    `json:"summary" validate:"required,gt=0,lte=2500"`
	Date    time.Time `json:"date" validate:"required"`
}

// New ...
func New(userID, summary string, date time.Time) *Task {

	return &Task{
		ID:      uuid.New().String(),
		UserID:  userID,
		Summary: summary,
		Date:    date,
	}
}
