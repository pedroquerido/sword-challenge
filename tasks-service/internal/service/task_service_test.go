package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/aes"

	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/task"

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

		returnErr := task.ErrorInvalidTask
		validator.EXPECT().
			Validate(gomock.Any()).
			Times(1).
			Return(returnErr)

		taskID, err := testService.Create(context.WithValue(context.Background(), service.ContextKey, testContext),
			"summary", time.Time{})
		require.NotNil(t, err)
		assert.True(t, errors.Is(err, service.ErrorInvalidTask))
		assert.Equal(t, taskID, "")
	})
	t.Run("should return error unexpected error - failure on encrypt", func(t *testing.T) {

		returnErr := errors.New("some random error")
		validator.EXPECT().
			Validate(gomock.Any()).
			Times(1).
			Return(nil)
		encryptor.EXPECT().
			Encrypt(gomock.Any()).
			Times(1).
			Return("", returnErr)

		taskID, err := testService.Create(context.WithValue(context.Background(), service.ContextKey, testContext),
			"summary", time.Now())
		require.NotNil(t, err)
		assert.True(t, errors.Is(err, service.ErrorUnexpectedError))
		assert.Equal(t, taskID, "")
	})
	t.Run("should return error unexpected error - failure on insert", func(t *testing.T) {

		returnErr := errors.New("some random error")
		validator.EXPECT().
			Validate(gomock.Any()).
			Times(1).
			Return(nil)
		encryptor.EXPECT().
			Encrypt(gomock.Any()).
			Times(1).
			Return("encrypted string", nil)
		repo.EXPECT().
			Insert(gomock.Any()).
			Times(1).
			Return(returnErr)

		taskID, err := testService.Create(context.WithValue(context.Background(), service.ContextKey, testContext),
			"summary", time.Now())
		require.NotNil(t, err)
		assert.True(t, errors.Is(err, service.ErrorUnexpectedError))
		assert.Equal(t, taskID, "")
	})
	t.Run("should return taskID", func(t *testing.T) {

		validator.EXPECT().
			Validate(gomock.Any()).
			Times(1).
			Return(nil)
		encryptor.EXPECT().
			Encrypt(gomock.Any()).
			Times(1).
			Return("encrypted string", nil)
		repo.EXPECT().
			Insert(gomock.Any()).
			Times(1).
			Return(nil)

		taskID, err := testService.Create(context.WithValue(context.Background(), service.ContextKey, testContext),
			"summary", time.Now())
		require.Nil(t, err)
		assert.NotNil(t, taskID)
	})
}

func TestTaskService_List(t *testing.T) {

	ctlr := gomock.NewController(t)
	repo := repoMock.NewMockTaskRepository(ctlr)
	validator := taskMock.NewMockValidator(ctlr)
	encryptor := aesMock.NewMockEncryptor(ctlr)
	testService := service.NewTaskService(repo, validator, encryptor)

	t.Run("should return error missing context", func(t *testing.T) {

		userID := "user_id"

		tasks, err := testService.List(context.Background(), &userID)
		require.NotNil(t, err)
		assert.True(t, errors.Is(err, service.ErrorMissingContext))
		assert.Equal(t, []*task.Task(nil), tasks)
	})
	t.Run("should return error user not allowed", func(t *testing.T) {

		isManager := false
		testContext := service.Context{
			UserID:    "user_id",
			IsManager: &isManager,
		}

		tasks, err := testService.List(context.WithValue(context.Background(), service.ContextKey, testContext),
			nil)
		require.NotNil(t, err)
		assert.True(t, errors.Is(err, service.ErrorUserNotAllowed))
		assert.Equal(t, []*task.Task(nil), tasks)
	})
	t.Run("should return unexpected error - failure on search", func(t *testing.T) {

		isManager := false
		userID := "user_id"
		testContext := service.Context{
			UserID:    userID,
			IsManager: &isManager,
		}

		returnTasks := []*task.Task(nil)
		returnErr := errors.New("some random error")

		repo.EXPECT().
			Search(gomock.Any()).
			Times(1).
			Return(returnTasks, returnErr)

		tasks, err := testService.List(context.WithValue(context.Background(), service.ContextKey, testContext),
			&userID)
		require.NotNil(t, err)
		assert.True(t, errors.Is(err, service.ErrorUnexpectedError))
		assert.Equal(t, returnTasks, tasks)
	})
	t.Run("should return unexpected error - failure on decrypt", func(t *testing.T) {

		isManager := false
		userID := "user_id"
		testContext := service.Context{
			UserID:    userID,
			IsManager: &isManager,
		}

		returnTasks := []*task.Task{
			task.New(userID, "summary", time.Now()),
		}
		returnErr := aes.ErrorEncryptedWithOtherEncryptor

		repo.EXPECT().
			Search(gomock.Any()).
			Times(1).
			Return(returnTasks, nil)
		encryptor.EXPECT().
			Decrypt(gomock.Any()).
			Times(1).
			Return("", returnErr)

		tasks, err := testService.List(context.WithValue(context.Background(), service.ContextKey, testContext),
			&userID)
		require.NotNil(t, err)
		assert.True(t, errors.Is(err, service.ErrorUnexpectedError))
		assert.Equal(t, []*task.Task(nil), tasks)
	})
	t.Run("should return no error - with decrypted task", func(t *testing.T) {

		isManager := false
		userID := "user_id"
		testContext := service.Context{
			UserID:    userID,
			IsManager: &isManager,
		}

		encryptedSummary := "encrypted"
		decryptedSummary := "decrypted"
		date := time.Now()
		returnTasksRepo := []*task.Task{
			&task.Task{
				ID:      "419d535a-ced5-4d5e-8885-49df44d3f5ff",
				UserID:  userID,
				Summary: encryptedSummary,
				Date:    date,
			},
		}
		returnTasksSvc := []*task.Task{
			&task.Task{
				ID:      "419d535a-ced5-4d5e-8885-49df44d3f5ff",
				UserID:  userID,
				Summary: decryptedSummary,
				Date:    date,
			},
		}

		repo.EXPECT().
			Search(gomock.Any()).
			Times(1).
			Return(returnTasksRepo, nil)
		encryptor.EXPECT().
			Decrypt(gomock.Any()).
			Times(1).
			Return(decryptedSummary, nil)

		tasks, err := testService.List(context.WithValue(context.Background(), service.ContextKey, testContext),
			&userID)
		require.Nil(t, err)
		assert.Equal(t, returnTasksSvc, tasks)
	})
	t.Run("should return no error - without decrypted task", func(t *testing.T) {

		isManager := true
		userID := "manager"
		testContext := service.Context{
			UserID:    userID,
			IsManager: &isManager,
		}

		encryptedSummary := "encrypted"
		decryptedSummary := "decrypted"
		date := time.Now()
		returnTasks := []*task.Task{
			&task.Task{
				ID:      "419d535a-ced5-4d5e-8885-49df44d3f5ff",
				UserID:  userID,
				Summary: encryptedSummary,
				Date:    date,
			},
		}

		repo.EXPECT().
			Search(gomock.Any()).
			Times(1).
			Return(returnTasks, nil)
		encryptor.EXPECT().
			Decrypt(gomock.Any()).
			Times(1).
			Return(decryptedSummary, nil)

		tasks, err := testService.List(context.WithValue(context.Background(), service.ContextKey, testContext),
			nil)
		require.Nil(t, err)
		assert.Equal(t, returnTasks, tasks)
	})
}

func TestTaskService_Retrieve(t *testing.T) {

	ctlr := gomock.NewController(t)
	repo := repoMock.NewMockTaskRepository(ctlr)
	validator := taskMock.NewMockValidator(ctlr)
	encryptor := aesMock.NewMockEncryptor(ctlr)
	testService := service.NewTaskService(repo, validator, encryptor)

	t.Run("should return error missing context", func(t *testing.T) {

		returnTask, err := testService.Retrieve(context.Background(), "419d535a-ced5-4d5e-8885-49df44d3f5ff")
		require.NotNil(t, err)
		assert.True(t, errors.Is(err, service.ErrorMissingContext))
		assert.Equal(t, (*task.Task)(nil), returnTask)
	})
	t.Run("should return unexpected error - failure on find", func(t *testing.T) {

		isManager := false
		testContext := service.Context{
			UserID:    "user_id",
			IsManager: &isManager,
		}

		returnErr := errors.New("some random error")

		repo.EXPECT().
			Find(gomock.Any()).
			Times(1).
			Return(nil, returnErr)

		returnTask, err := testService.Retrieve(context.WithValue(context.Background(), service.ContextKey, testContext),
			"419d535a-ced5-4d5e-8885-49df44d3f5ff")
		require.NotNil(t, err)
		assert.True(t, errors.Is(err, service.ErrorUnexpectedError))
		assert.Equal(t, (*task.Task)(nil), returnTask)
	})
	t.Run("should return error user not allowed", func(t *testing.T) {

		isManager := false
		testContext := service.Context{
			UserID:    "user_id",
			IsManager: &isManager,
		}

		returnTaskRepo := &task.Task{
			ID:      "419d535a-ced5-4d5e-8885-49df44d3f5ff",
			UserID:  "user_id2",
			Summary: "summary",
			Date:    time.Now(),
		}

		repo.EXPECT().
			Find(gomock.Any()).
			Times(1).
			Return(returnTaskRepo, nil)

		returnTaskSvc, err := testService.Retrieve(context.WithValue(context.Background(), service.ContextKey, testContext),
			"419d535a-ced5-4d5e-8885-49df44d3f5ff")
		require.NotNil(t, err)
		assert.True(t, errors.Is(err, service.ErrorUserNotAllowed))
		assert.Equal(t, (*task.Task)(nil), returnTaskSvc)
	})
	t.Run("should return error unexpected error", func(t *testing.T) {

		isManager := true
		userID := "manager"
		testContext := service.Context{
			UserID:    userID,
			IsManager: &isManager,
		}

		encryptedSummary := "encrypted"
		returnTaskRepo := &task.Task{
			ID:      "419d535a-ced5-4d5e-8885-49df44d3f5ff",
			UserID:  userID,
			Summary: encryptedSummary,
			Date:    time.Now(),
		}

		returnErr := aes.ErrorEncryptedWithOtherEncryptor

		repo.EXPECT().
			Find(gomock.Any()).
			Times(1).
			Return(returnTaskRepo, nil)
		encryptor.EXPECT().
			Decrypt(gomock.Any()).
			Times(1).
			Return("", returnErr)

		returnTaskSvc, err := testService.Retrieve(context.WithValue(context.Background(), service.ContextKey, testContext),
			"419d535a-ced5-4d5e-8885-49df44d3f5ff")
		require.NotNil(t, err)
		assert.True(t, errors.Is(err, service.ErrorUnexpectedError))
		assert.Equal(t, (*task.Task)(nil), returnTaskSvc)
	})
	t.Run("should return no error - without decrypted task", func(t *testing.T) {

		isManager := true
		testContext := service.Context{
			UserID:    "user_id",
			IsManager: &isManager,
		}

		encryptedSummary := "encrypted"
		decryptedSummary := "decrypted"
		returnTaskRepo := &task.Task{
			ID:      "419d535a-ced5-4d5e-8885-49df44d3f5ff",
			UserID:  "user_id2",
			Summary: encryptedSummary,
			Date:    time.Now(),
		}

		repo.EXPECT().
			Find(gomock.Any()).
			Times(1).
			Return(returnTaskRepo, nil)
		encryptor.EXPECT().
			Decrypt(gomock.Any()).
			Times(1).
			Return(decryptedSummary, nil)

		returnTaskSvc, err := testService.Retrieve(context.WithValue(context.Background(), service.ContextKey, testContext),
			"419d535a-ced5-4d5e-8885-49df44d3f5ff")
		require.Nil(t, err)
		assert.NotNil(t, returnTaskSvc)
		assert.Equal(t, returnTaskRepo, returnTaskSvc)
	})
	t.Run("should return no error - with decrypted task", func(t *testing.T) {

		isManager := false
		userID := "manager"
		testContext := service.Context{
			UserID:    userID,
			IsManager: &isManager,
		}

		encryptedSummary := "encrypted"
		decryptedSummary := "decrypted"
		date := time.Now()
		returnTaskRepo := &task.Task{
			ID:      "419d535a-ced5-4d5e-8885-49df44d3f5ff",
			UserID:  userID,
			Summary: encryptedSummary,
			Date:    date,
		}
		expectedReturnTaskSvc := &task.Task{
			ID:      "419d535a-ced5-4d5e-8885-49df44d3f5ff",
			UserID:  userID,
			Summary: decryptedSummary,
			Date:    date,
		}

		repo.EXPECT().
			Find(gomock.Any()).
			Times(1).
			Return(returnTaskRepo, nil)
		encryptor.EXPECT().
			Decrypt(gomock.Any()).
			Times(1).
			Return(decryptedSummary, nil)

		returnTaskSvc, err := testService.Retrieve(context.WithValue(context.Background(), service.ContextKey, testContext),
			"419d535a-ced5-4d5e-8885-49df44d3f5ff")
		require.Nil(t, err)
		assert.NotNil(t, returnTaskSvc)
		assert.Equal(t, expectedReturnTaskSvc, returnTaskSvc)
	})
}
