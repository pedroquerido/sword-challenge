package repo

import (
	"context"

	"github.com/pedroquerido/sword-challenge/service-tasks/pkg/task"
)

type TaskRepository interface {
	Save(ctx context.Context, task *task.Task) error
}
