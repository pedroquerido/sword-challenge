package mysql

import (
	"context"
	"time"

	"github.com/pedroquerido/sword-challenge/service-tasks/pkg/task"

	"github.com/pedroquerido/sword-challenge/service-tasks/internal/repo"
)

type taskRow struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"-"`
	TaskID    string    `gorm:"not null; unique;" json:"id"`
	Summary   string    `gorm:"size:2500;not null;" json:"summary"`
	Date      time.Time `gorm:"not null" json:"date"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func fromTask(ctx context.Context, task *task.Task) *taskRow {

	return &taskRow{
		TaskID:  task.ID,
		Summary: task.Summary,
		Date:    task.Date,
	}
}

func (r *taskRow) toTask() *task.Task {

	return &task.Task{
		ID:      r.TaskID,
		Summary: r.Summary,
		Date:    r.Date,
	}
}

func (r *taskRow) validate() error {

	if r == nil {
		return repo.ErrorEmptyTask
	}

	if r.Summary == "" {
		return repo.ErrorEmptyTaskSummary
	}

	if r.Date.IsZero() {
		return repo.ErrorEmptyTaskDate
	}

	return nil
}
