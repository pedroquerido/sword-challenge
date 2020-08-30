package service

import (
	"context"
	"tasks-service/internal/repo"
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
func (s *TaskService) CreateTask(ctx context.Context, userID, summary string, date time.Time) (string, error) {

	task := task.NewTask(userID, summary, date)

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
