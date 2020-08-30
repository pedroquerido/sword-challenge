package gorm

import (
	"time"

	"tasks-service/pkg/task"

	"gorm.io/gorm"
)

type taskRow struct {
	ID        uint64    `gorm:"colummn:id;primary_key;auto_increment" json:"-"`
	TaskID    string    `gorm:"colummn:task_id;not null; unique;" json:"id"`
	Summary   string    `gorm:"collumn:summary;size:2500;not null;" json:"summary"`
	Date      time.Time `gorm:"collumn:date;not null" json:"date"`
	CreatedAt time.Time `gorm:"collumn:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"collumn:updated_at" json:"updated_at"`
}

func fromTask(task *task.Task) *taskRow {

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

func (r *taskRow) BeforeCreate(tx *gorm.DB) (err error) {

	r.CreatedAt = time.Now().UTC()

	return nil
}

func (r *taskRow) BeforeUpdate(tx *gorm.DB) (err error) {

	r.UpdatedAt = time.Now().UTC()

	return nil
}
