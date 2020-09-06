package service

import (
	"context"
	"time"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/repo"
	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/aes"
	pkgError "github.com/pedroquerido/sword-challenge/tasks-service/pkg/error"
	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/task"
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
	repo      repo.TaskRepository
	validator task.Validator
	encryptor aes.Encryptor
}

// NewTaskService initializes and returns a new taskService implementation of TaskService
func NewTaskService(repo repo.TaskRepository, validator task.Validator, encryptor aes.Encryptor) TaskService {

	return &taskService{
		repo:      repo,
		validator: validator,
		encryptor: encryptor,
	}
}

// Create ...
func (s *taskService) Create(ctx context.Context, summary string, date time.Time) (string, error) {

	serviceContext, err := parseContext(ctx)
	if err != nil {
		return "", err
	}

	// validate task
	task := task.New(serviceContext.UserID, summary, date)
	if err := s.validator.Validate(task); err != nil {
		return "", parseExternalError(err)
	}

	// encrypt summary
	task.Summary, err = s.encryptor.Encrypt(task.Summary)
	if err != nil {
		return "", parseExternalError(err)
	}

	// persist
	if err := s.repo.Insert(task); err != nil {
		return "", parseExternalError(err)
	}

	return task.ID, nil
}

// List ...
func (s *taskService) List(ctx context.Context, userID *string) ([]*task.Task, error) {

	serviceContext, err := parseContext(ctx)
	if err != nil {
		return nil, err
	}

	// validate access
	if (userID == nil || *userID != serviceContext.UserID) && !*serviceContext.IsManager {
		return nil, pkgError.NewError(ErrorUserNotAllowed)
	}

	// search
	tasks, err := s.repo.Search(userID)
	if err != nil {
		return nil, parseExternalError(err)
	}

	// decrypt user tasks
	for i := range tasks {
		if tasks[i].UserID == serviceContext.UserID {
			tasks[i].Summary, err = s.encryptor.Decrypt(tasks[i].Summary)
			if err != nil {
				return nil, parseExternalError(err)
			}
		}
	}

	return tasks, nil
}

// Retrieve ...
func (s *taskService) Retrieve(ctx context.Context, taskID string) (*task.Task, error) {

	serviceContext, err := parseContext(ctx)
	if err != nil {
		return nil, err
	}

	// find task
	task, err := s.repo.Find(taskID)
	if err != nil {
		return nil, parseExternalError(err)
	}

	// validate access
	if task.UserID != serviceContext.UserID && !*serviceContext.IsManager {
		return nil, pkgError.NewError(ErrorUserNotAllowed)
	}

	// decrypt
	if task.UserID == serviceContext.UserID {
		task.Summary, err = s.encryptor.Decrypt(task.Summary)
		if err != nil {
			return nil, parseExternalError(err)
		}
	}

	return task, nil
}

// Update ...
func (s *taskService) Update(ctx context.Context, taskID string, summary *string, date *time.Time) error {

	serviceContext, err := parseContext(ctx)
	if err != nil {
		return err
	}

	// find
	task, err := s.repo.Find(taskID)
	if err != nil {
		return parseExternalError(err)
	}

	// validate access
	if task.UserID != serviceContext.UserID {
		return pkgError.NewError(ErrorUserNotAllowed)
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

	if err := s.validator.Validate(task); err != nil {
		return parseExternalError(err)
	}

	// encrypt summary
	task.Summary, err = s.encryptor.Encrypt(task.Summary)
	if err != nil {
		return parseExternalError(err)
	}

	// update
	err = s.repo.Update(taskID, summary, date)
	if err != nil {
		return parseExternalError(err)
	}

	return nil
}

// DeleteTask ...
func (s *taskService) Delete(ctx context.Context, taskID string) error {

	serviceContext, err := parseContext(ctx)
	if err != nil {
		return err
	}

	// validate access
	if !*serviceContext.IsManager {
		return pkgError.NewError(ErrorUserNotAllowed)
	}

	// delete
	err = s.repo.Delete(taskID)
	if err != nil {
		return parseExternalError(err)
	}

	return nil
}
