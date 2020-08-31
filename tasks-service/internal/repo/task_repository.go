package repo

import (
	"tasks-service/pkg/task"
)

// TaskRepository ...
type TaskRepository interface {
	Save(task *task.Task) error
	Find(id string) (*task.Task, error)
	Search(limit *int, offset *int64, userID *string) (tasks []*task.Task, count int64, err error)
}
