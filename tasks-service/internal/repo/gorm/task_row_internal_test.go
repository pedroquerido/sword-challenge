package gorm

import (
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/task"
	"github.com/stretchr/testify/assert"
)

func TestFromTask(t *testing.T) {
	t.Run("should return task row", func(t *testing.T) {
		task := &task.Task{
			ID:      uuid.New().String(),
			UserID:  "user-1",
			Summary: "summary-1",
			Date:    time.Now(),
		}
		taskRow := fromTask(task)
		assert.Equal(t, task, taskRow.toTask())
	})
	t.Run("should return empty task row", func(t *testing.T) {
		emptyTaskRow := fromTask(nil)
		assert.Equal(t, emptyTaskRow, &taskRow{})
	})
}

func TestTaskRow_TableName(t *testing.T) {
	t.Run("should return string", func(t *testing.T) {
		taskRow := fromTask(nil)
		assert.Equal(t, "tasks", taskRow.TableName())
	})
}

func TestTaskRow_BeforeCreate(t *testing.T) {
	t.Run("should set CreatedAt", func(t *testing.T) {
		taskRow := fromTask(nil)
		err := taskRow.BeforeCreate(nil)
		assert.Nil(t, err)
		assert.False(t, taskRow.CreatedAt.IsZero())
	})
}

func TestTaskRow_BeforeUpdate(t *testing.T) {
	t.Run("should set UpdatedAt", func(t *testing.T) {
		taskRow := fromTask(nil)
		err := taskRow.BeforeUpdate(nil)
		assert.Nil(t, err)
		assert.False(t, taskRow.UpdatedAt.IsZero())
	})
}
