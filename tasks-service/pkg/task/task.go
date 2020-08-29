package task

import "time"

// Task ...
type Task struct {
	ID      string    `json:"id"`
	UserID  string    `json:"user_id"`
	Summary string    `json:"summary"`
	Date    time.Time `json:"date"`
}
