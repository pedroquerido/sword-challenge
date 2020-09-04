package repo

import (
	"time"

	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/task"
)

// TaskRepository ...
type TaskRepository interface {
	Insert(task *task.Task) error
	Find(id string) (*task.Task, error)
	Search(userID *string) (tasks []*task.Task, err error)
	Update(id, userID string, summary *string, date *time.Time) error
	Delete(id string) error
}
