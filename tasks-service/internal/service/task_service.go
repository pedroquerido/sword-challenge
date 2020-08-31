package service

import (
	"context"
	"tasks-service/internal/repo"
	pkgError "tasks-service/pkg/error"
	"tasks-service/pkg/task"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

// TaskService represents the Use Case layer containing application specific business rules
type TaskService struct {
	repo     repo.TaskRepository
	validate *TaskValidator
}

// NewTaskService initializes and returns a new TaskService instance
func NewTaskService(repo repo.TaskRepository, validate *validator.Validate) *TaskService {

	return &TaskService{
		repo:     repo,
		validate: NewTaskValidator(validate),
	}
}

// CreateTask ...
func (s *TaskService) CreateTask(ctx context.Context, summary string, date time.Time) (string, error) {

	serviceContext, err := parseContext(ctx)
	if err != nil {
		return "", err
	}

	task := task.NewTask(serviceContext.UserID, summary, date)
	if err := s.validate.Validate(task); err != nil {
		return "", err
	}

	// Call repo

	return task.ID, nil
}

// ListTasks ...
func (s *TaskService) ListTasks(ctx context.Context) ([]*task.Task, error) {

	return nil, nil
}

// RetrieveTask ...
func (s *TaskService) RetrieveTask(ctx context.Context, taskID string) (*task.Task, error) {

	return nil, nil
}

// UpdateTask ...
func (s *TaskService) UpdateTask(ctx context.Context, taskID string) error {

	return nil
}

// DeleteTask ...
func (s *TaskService) DeleteTask(ctx context.Context, taskID string) error {

	return nil
}

func parseContext(ctx context.Context) (*Context, error) {

	context := ctx.Value(ContextKey)

	serviceContext, ok := context.(Context)
	if !ok {
		return nil, pkgError.NewError(ErrorMissingContext)
	}

	return &serviceContext, nil
}
