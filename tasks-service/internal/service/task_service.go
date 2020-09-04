package service

import (
	"context"
	"errors"
	"time"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/repo"
	pkgError "github.com/pedroquerido/sword-challenge/tasks-service/pkg/error"
	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/task"

	"gopkg.in/go-playground/validator.v9"
)

var _ TaskService = (*taskService)(nil)

// TaskService represents the Use Case layer containing application specific business rules
type TaskService interface {
	Create(ctx context.Context, summary string, date time.Time) (string, error)
	List(ctx context.Context, userID *string) ([]*task.Task, error)
	Retrieve(ctx context.Context, taskID string) (*task.Task, error)
	Update(ctx context.Context, taskID string, summary *string, date *time.Time) error
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

// Create ...
func (s *taskService) Create(ctx context.Context, summary string, date time.Time) (string, error) {

	serviceContext, err := parseContext(ctx)
	if err != nil {
		return "", err
	}

	task := task.NewTask(serviceContext.UserID, summary, date)
	if err := s.validate.Validate(task); err != nil {
		return "", err
	}

	if err := s.repo.Insert(task); err != nil {
		return "", err
	}

	return task.ID, nil
}

// List ...
func (s *taskService) List(ctx context.Context, userID *string) ([]*task.Task, error) {

	serviceContext, err := parseContext(ctx)
	if err != nil {
		return nil, err
	}

	if userID == nil && !*serviceContext.IsManager {
		return nil, pkgError.NewError(ErrorUserNotAllowed)
	}

	return s.repo.Search(userID)
}

// Retrieve ...
func (s *taskService) Retrieve(ctx context.Context, taskID string) (*task.Task, error) {

	serviceContext, err := parseContext(ctx)
	if err != nil {
		return nil, err
	}

	task, err := s.repo.Find(taskID)
	if err != nil {

		if errors.Is(err, repo.ErrorNotFound) {
			return nil, pkgError.NewError(ErrorTaskNotFound)
		}

		return nil, err
	}

	if task.UserID != serviceContext.UserID && !*serviceContext.IsManager {
		return nil, pkgError.NewError(ErrorTaskNotFound)
	}

	return task, nil
}

// Update ...
func (s *taskService) Update(ctx context.Context, taskID string, summary *string, date *time.Time) error {

	serviceContext, err := parseContext(ctx)
	if err != nil {
		return err
	}

	task, err := s.Retrieve(ctx, taskID)
	if err != nil {
		return err
	}

	if task.UserID != serviceContext.UserID {
		return pkgError.NewError(ErrorTaskNotFound)
	}

	// Validate inputs
	if summary == nil && date == nil {
		return nil
	}

	if summary != nil {
		task.Summary = *summary
	}

	if date != nil {
		task.Date = *date
	}

	if err := s.validate.Validate(task); err != nil {
		return err
	}

	err = s.repo.Update(taskID, serviceContext.UserID, summary, date)
	if err != nil { // should not really happen with all the previous validations

		if errors.Is(err, repo.ErrorNotFound) {
			return pkgError.NewError(ErrorTaskNotFound)
		}

		return err
	}

	return nil
}

// DeleteTask ...
func (s *taskService) Delete(ctx context.Context, taskID string) error {

	serviceContext, err := parseContext(ctx)
	if err != nil {
		return err
	}

	if !*serviceContext.IsManager {
		return pkgError.NewError(ErrorUserNotAllowed)
	}

	err = s.repo.Delete(taskID)
	if err != nil {

		if errors.Is(err, repo.ErrorNotFound) {
			return pkgError.NewError(ErrorTaskNotFound)
		}

		return err
	}

	return nil
}
