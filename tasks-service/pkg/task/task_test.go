package task_test

import (
	"testing"
	"time"

	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/task"
	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
)

func TestNew(t *testing.T) {
	t.Run("should return task with generated uuid", func(t *testing.T) {

		userID := "user_id"
		summary := "summary"
		date := time.Now()

		task := task.New(userID, summary, date)
		assert.NotNil(t, task)
		assert.Equal(t, task.UserID, userID)
		assert.Equal(t, task.Summary, summary)
		assert.Equal(t, task.Date, date)
		assert.NotNil(t, task.ID)

		generatedUUID, err := uuid.Parse(task.ID)
		assert.NotNil(t, generatedUUID)
		assert.Nil(t, err)
	})
}
