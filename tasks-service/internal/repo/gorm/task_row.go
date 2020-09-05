package gorm

import (
	"time"

	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/task"

	"gorm.io/gorm"
)

type taskRow struct {
	ID        uint64    `gorm:"colummn:id;primary_key;auto_increment;"`
	TaskID    string    `gorm:"colummn:task_id;not null; unique;"`
	UserID    string    `gorm:"collumn:user_id;not null;"`
	Summary   string    `gorm:"collumn:summary;not null;"`
	Date      time.Time `gorm:"collumn:date;not null"`
	CreatedAt time.Time `gorm:"collumn:created_at"`
	UpdatedAt time.Time `gorm:"collumn:updated_at"`
}

func fromTask(task *task.Task) *taskRow {

	if task == nil {
		return &taskRow{}
	}

	return &taskRow{
		TaskID:  task.ID,
		UserID:  task.UserID,
		Summary: task.Summary,
		Date:    task.Date,
	}
}

func (r *taskRow) toTask() *task.Task {

	return &task.Task{
		ID:      r.TaskID,
		UserID:  r.UserID,
		Summary: r.Summary,
		Date:    r.Date,
	}
}

func (r *taskRow) TableName() string {
	return "tasks"
}

func (r *taskRow) BeforeCreate(tx *gorm.DB) (err error) {

	r.CreatedAt = time.Now()

	return nil
}

func (r *taskRow) BeforeUpdate(tx *gorm.DB) (err error) {

	r.UpdatedAt = time.Now()

	return nil
}
