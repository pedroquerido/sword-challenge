package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/service"

	repoMock "github.com/pedroquerido/sword-challenge/tasks-service/internal/repo/mock"
	aesMock "github.com/pedroquerido/sword-challenge/tasks-service/pkg/aes/mock"
	taskMock "github.com/pedroquerido/sword-challenge/tasks-service/pkg/task/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	gomock "github.com/golang/mock/gomock"
)

func TestTaskService_Create(t *testing.T) {

	ctlr := gomock.NewController(t)
	repo := repoMock.NewMockTaskRepository(ctlr)
	validator := taskMock.NewMockValidator(ctlr)
	encryptor := aesMock.NewMockEncryptor(ctlr)
	testService := service.NewTaskService(repo, validator, encryptor)

	isManager := false
	testContext := service.Context{
		UserID:    "user_id",
		IsManager: &isManager,
	}

	t.Run("should return error missing context", func(t *testing.T) {
		taskID, err := testService.Create(context.Background(), "summary", time.Now())
		require.NotNil(t, err)
		assert.True(t, errors.Is(err, service.ErrorMissingContext))
		assert.Equal(t, taskID, "")
	})
	t.Run("should return error invalid task", func(t *testing.T) {

		returnErr := []error{errors.New("validation error")}
		validator.EXPECT().
			Validate(gomock.Any()).
			Times(1).
			Return(returnErr)

		taskID, err := testService.Create(context.WithValue(context.Background(), service.ContextKey, testContext),
			"summary", time.Time{})
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, service.ErrorInvalidTask))
		assert.Equal(t, taskID, "")
	})
	t.Run("should return error unexpected error", func(t *testing.T) {

		returnErr := errors.New("some random error")
		validator.EXPECT().
			Validate(gomock.Any()).
			Times(1).
			Return(nil)
		repo.EXPECT().
			Insert(gomock.Any()).
			Times(1).
			Return(returnErr)

		taskID, err := testService.Create(context.WithValue(context.Background(), service.ContextKey, testContext),
			"summary", time.Now())
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, service.ErrorUnexpectedError))
		assert.Equal(t, taskID, "")
	})
	t.Run("should return taskID", func(t *testing.T) {

		validator.EXPECT().
			Validate(gomock.Any()).
			Times(1).
			Return(nil)
		repo.EXPECT().
			Insert(gomock.Any()).
			Times(1).
			Return(nil)

		taskID, err := testService.Create(context.WithValue(context.Background(), service.ContextKey, testContext),
			"summary", time.Now())
		assert.Nil(t, err)
		assert.NotNil(t, taskID)
	})
}
