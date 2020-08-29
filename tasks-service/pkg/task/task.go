package task

import "time"

type Task struct {
	ID      string    `json:"id"`
	Summary string    `json:"summary"`
	Date    time.Time `json:"date"`
}
