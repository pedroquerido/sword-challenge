package repo

import (
	"tasks-service/pkg/task"
)

// TaskRepository ...
type TaskRepository interface {
	Save(task *task.Task) error
	Find(id string) (*task.Task, error)
	List() ([]*task.Task, error)
}
