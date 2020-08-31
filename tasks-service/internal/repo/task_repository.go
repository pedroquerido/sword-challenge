package repo

import (
	"tasks-service/pkg/task"
	"time"
)

// TaskRepository ...
type TaskRepository interface {
	Insert(task *task.Task) error
	Find(id string) (*task.Task, error)
	Search(limit *int, offset *int64, userID *string) (tasks []*task.Task, count int64, err error)
	Update(id string, summary *string, date *time.Time) error
	Delete(id string) error
}
