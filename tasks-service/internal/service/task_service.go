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

func (s *TaskService) createTask(ctx context.Context, userID, summary string, date time.Time) (string, error) {

	task := task.NewTask(userID, summary, date)

	if err := s.validate.Validate(task); err != nil {
		return "", err
	}

	// Call repo

	return task.ID, nil
}

func (s *TaskService) listTasks(ctx context.Context) ([]*task.Task, error) {

	return nil, nil
}

func (s *TaskService) retrieveTask(ctx context.Context, taskID string) (*task.Task, error) {

	return nil, nil
}

func (s *TaskService) updateTask(ctx context.Context, taskID string) error {

	return nil
}

func (s *TaskService) deleteTask(ctx context.Context, taskID string) error {

	return nil
}
