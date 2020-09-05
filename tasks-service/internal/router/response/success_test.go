package response_test

import (
	"testing"
	"time"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router/response"
	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/task"
	"github.com/stretchr/testify/assert"
)

func TestNewCreateTaskResponse(t *testing.T) {

	t.Run("should return create task response", func(t *testing.T) {

		message := "message"
		code := 201
		data := "task_id"

		successResponse := response.NewCreateTaskResponse(code, message, data)
		assert.NotNil(t, successResponse)
		assert.Equal(t, message, successResponse.Message)
		assert.Equal(t, code, successResponse.Code)
		assert.Equal(t, data, successResponse.Data)
	})
}

func TestNewListTasksResponse(t *testing.T) {

	t.Run("should return list tasks response", func(t *testing.T) {

		message := "message"
		code := 200
		data := []*task.Task{
			&task.Task{
				ID:      "task_id",
				UserID:  "user_id",
				Summary: "summary",
				Date:    time.Now(),
			},
		}

		successResponse := response.NewListTasksResponse(code, message, data)
		assert.NotNil(t, successResponse)
		assert.Equal(t, message, successResponse.Message)
		assert.Equal(t, code, successResponse.Code)
		assert.Equal(t, data, successResponse.Data)
	})
}

func TestNewListUserTasksResponse(t *testing.T) {

	t.Run("should return list user tasks response", func(t *testing.T) {

		message := "message"
		code := 200
		data := []*task.Task{
			&task.Task{
				ID:      "task_id",
				UserID:  "user_id",
				Summary: "summary",
				Date:    time.Now(),
			},
		}

		successResponse := response.NewListUserTasksResponse(code, message, data)
		assert.NotNil(t, successResponse)
		assert.Equal(t, message, successResponse.Message)
		assert.Equal(t, code, successResponse.Code)
		assert.Equal(t, data, successResponse.Data)
	})
}

func TestNewRetrieveTaskResponse(t *testing.T) {

	t.Run("should return retrieve task response", func(t *testing.T) {

		message := "message"
		code := 200
		data := &task.Task{
			ID:      "task_id",
			UserID:  "user_id",
			Summary: "summary",
			Date:    time.Now(),
		}

		successResponse := response.NewRetrieveTaskResponse(code, message, data)
		assert.NotNil(t, successResponse)
		assert.Equal(t, message, successResponse.Message)
		assert.Equal(t, code, successResponse.Code)
		assert.Equal(t, data, successResponse.Data)
	})
}

func TestNewUpdateTaskResponse(t *testing.T) {

	t.Run("should return update task response", func(t *testing.T) {

		message := "message"
		code := 200

		successResponse := response.NewUpdateTaskResponse(code, message)
		assert.NotNil(t, successResponse)
		assert.Equal(t, message, successResponse.Message)
		assert.Equal(t, code, successResponse.Code)
	})
}

func TestNewDeleteTaskResponse(t *testing.T) {

	t.Run("should return update task response", func(t *testing.T) {

		message := "message"
		code := 200

		successResponse := response.NewDeleteTaskResponse(code, message)
		assert.NotNil(t, successResponse)
		assert.Equal(t, message, successResponse.Message)
		assert.Equal(t, code, successResponse.Code)
	})
}
