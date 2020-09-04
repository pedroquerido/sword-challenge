package service_test

import (
	"errors"
	"testing"
	"time"

	"math/rand"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/service"
	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/task"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
)

const (
	testCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	testStringlength = 10
)

func generateRandomString(t *testing.T) string {
	t.Helper()

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, testStringlength)
	for i := range b {
		b[i] = testCharset[seededRand.Intn(len(testCharset))]
	}

	return string(b)
}

func TestTaskValidator_Validate(t *testing.T) {

	testTaskValidator := service.NewTaskValidator(validator.New())

	t.Run("should return no error", func(t *testing.T) {

		testTask := &task.Task{
			ID:      "21aaa98b-842c-484a-a925-79aef811e68d",
			UserID:  "user_id",
			Summary: generateRandomString(t),
			Date:    time.Now(),
		}

		err := testTaskValidator.Validate(testTask)
		assert.Nil(t, err)
	})
	t.Run("should return error invalid task - empty task", func(t *testing.T) {

		testTask := &task.Task{}

		err := testTaskValidator.Validate(testTask)
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, service.ErrorInvalidTask))
	})
	t.Run("should return error invalid task - empty task_id", func(t *testing.T) {

		testTask := &task.Task{
			UserID:  "user_id",
			Summary: generateRandomString(t),
			Date:    time.Now(),
		}

		err := testTaskValidator.Validate(testTask)
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, service.ErrorInvalidTask))
	})
	t.Run("should return error invalid task - empty task_id", func(t *testing.T) {

		testTask := &task.Task{
			UserID:  "user_id",
			Summary: generateRandomString(t),
			Date:    time.Now(),
		}

		err := testTaskValidator.Validate(testTask)
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, service.ErrorInvalidTask))
	})
	t.Run("should return error invalid task - task_id not uuid", func(t *testing.T) {

		testTask := &task.Task{
			ID:      "task_id",
			UserID:  "user_id",
			Summary: generateRandomString(t),
			Date:    time.Now(),
		}

		err := testTaskValidator.Validate(testTask)
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, service.ErrorInvalidTask))
	})
	t.Run("should return error invalid task - task_id invalid uuid", func(t *testing.T) {

		testTask := &task.Task{
			ID:      "21aaa98b-842c-484a-a925-79aef811e68D",
			UserID:  "user_id",
			Summary: generateRandomString(t),
			Date:    time.Now(),
		}

		err := testTaskValidator.Validate(testTask)
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, service.ErrorInvalidTask))
	})
	t.Run("should return error invalid task - empty user_id", func(t *testing.T) {

		testTask := &task.Task{
			ID:      "21aaa98b-842c-484a-a925-79aef811e68d",
			Summary: generateRandomString(t),
			Date:    time.Now(),
		}

		err := testTaskValidator.Validate(testTask)
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, service.ErrorInvalidTask))
	})
	t.Run("should return error invalid task - empty summary", func(t *testing.T) {

		testStringlength = 2501

		testTask := &task.Task{
			ID:     "21aaa98b-842c-484a-a925-79aef811e68d",
			UserID: "user_id",
			Date:   time.Now(),
		}

		err := testTaskValidator.Validate(testTask)
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, service.ErrorInvalidTask))
	})
	t.Run("should return error invalid task - summary too long", func(t *testing.T) {

		testStringlength = 2501

		testTask := &task.Task{
			ID:      "21aaa98b-842c-484a-a925-79aef811e68d",
			UserID:  "user_id",
			Summary: generateRandomString(t),
			Date:    time.Now(),
		}

		err := testTaskValidator.Validate(testTask)
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, service.ErrorInvalidTask))

		testStringlength = 10
	})
	t.Run("should return error invalid task - empty date", func(t *testing.T) {

		testStringlength = 2501

		testTask := &task.Task{
			ID:      "21aaa98b-842c-484a-a925-79aef811e68d",
			UserID:  "user_id",
			Summary: generateRandomString(t),
		}

		err := testTaskValidator.Validate(testTask)
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, service.ErrorInvalidTask))
	})
}
