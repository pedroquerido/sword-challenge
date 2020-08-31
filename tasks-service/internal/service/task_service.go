package service

import (
	"context"
	"tasks-service/internal/repo"
	"tasks-service/pkg/task"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

var _ TaskService = (*taskService)(nil)

// TaskService represents the Use Case layer containing application specific business rules
type TaskService interface {
	Create(ctx context.Context, summary string, date time.Time) (string, error)
	List(ctx context.Context) ([]*task.Task, error)
	Retrieve(ctx context.Context, taskID string) (*task.Task, error)
	Update(ctx context.Context, taskID string) error
	Delete(ctx context.Context, taskID string) error
}

type taskService struct {
	repo     repo.TaskRepository
	validate *TaskValidator
}

// NewTaskService initializes and returns a new taskService implementation of TaskService
func NewTaskService(repo repo.TaskRepository, validate *validator.Validate) TaskService {

	return &taskService{
		repo:     repo,
		validate: NewTaskValidator(validate),
	}
}

// CreateTask ...
func (s *taskService) Create(ctx context.Context, summary string, date time.Time) (string, error) {

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
func (s *taskService) List(ctx context.Context) ([]*task.Task, error) {

	return nil, nil
}

// RetrieveTask ...
func (s *taskService) Retrieve(ctx context.Context, taskID string) (*task.Task, error) {

	return nil, nil
}

// UpdateTask ...
func (s *taskService) Update(ctx context.Context, taskID string) error {

	return nil
}

// DeleteTask ...
func (s *taskService) Delete(ctx context.Context, taskID string) error {

	return nil
}
