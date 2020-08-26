package repo

import (
	"github.com/pedroquerido/sword-challenge/service-tasks/pkg/task"
)

// TaskRepository ...
type TaskRepository interface {
	Save(task *task.Task) error
	Find(id string) (*task.Task, error)
	List() ([]*task.Task, error)
}
