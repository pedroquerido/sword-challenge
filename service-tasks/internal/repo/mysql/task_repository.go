package mysql

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/pedroquerido/sword-challenge/service-tasks/internal/repo"
	"github.com/pedroquerido/sword-challenge/service-tasks/pkg/task"
)

var _ repo.TaskRepository = (*TaskRepository)(nil)

// TaskRepository ...
type TaskRepository struct {
	db *gorm.DB
}

// Save ...
func (r *TaskRepository) Save(ctx context.Context, task *task.Task) error {

	row := fromTask(ctx, task)
	err := row.validate()
	if err != nil {
		return err
	}

	err = r.db.Debug().Create(&row).Error
	if err != nil {
		return repo.ErrorUnknown.Wrap(err)
	}

	return nil
}

// Find ...
func (r *TaskRepository) Find(ctx context.Context, id string) (*task.Task, error) {

	row := taskRow{}

	if err := r.db.Debug().Where("id = ?", id).Take(&row).Error; err != nil {
		return &task.Task{}, repo.ErrorUnknown.Wrap(err)
	}

	return row.toTask(), nil
}

// List ...
func (r *TaskRepository) List(ctx context.Context) (*[]task.Task, error) {

	rows := []taskRow{}

	if err := r.db.Debug().Find(&rows).Error; err != nil {
		return &[]task.Task{}, repo.ErrorUnknown.Wrap(err)
	}

	tasks := make([]task.Task, 0, len(rows))
	for _, row := range rows {
		tasks = append(tasks, *row.toTask())
	}

	return &tasks, nil
}
