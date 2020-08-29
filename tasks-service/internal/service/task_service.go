package service

import (
	"context"
	"tasks-service/internal/repo"
	"tasks-service/pkg/task"
)

// TaskService represents the Use Case layer containing application specific business rules
type TaskService struct {
	repo repo.TaskRepository
}

// NewTaskService initializes and returns a new TaskService instance
func NewTaskService(repo repo.TaskRepository) *TaskService {

	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) createTask(ctx context.Context) (string, error) {

	return "", nil
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
