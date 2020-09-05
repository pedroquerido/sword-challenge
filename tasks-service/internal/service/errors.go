package service

import (
	"errors"
	"log"

	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/task"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/repo"
	pkgError "github.com/pedroquerido/sword-challenge/tasks-service/pkg/error"
)

var (
	// ErrorInvalidTask represents the error obtained if a task fails to meet requirements
	ErrorInvalidTask = errors.New("invalid task")

	// ErrorTaskNotFound represents the error obtained if a requested task is not found
	ErrorTaskNotFound = errors.New("task not found")

	// ErrorUserNotAllowed represents the error obtained if user is not allowed to perform the requested action
	ErrorUserNotAllowed = errors.New("user not allowed")

	// ErrorMissingContext represents the error obtained if Context is not found on the request
	ErrorMissingContext = errors.New("missing context")

	// ErrorUnexpectedError represents the error obtained if something unexpected happened
	ErrorUnexpectedError = errors.New("unexpected error")
)

func parseExternalError(err error) error {

	if err != nil {

		var (
			structuredError pkgError.Error
		)

		if errors.Is(err, task.ErrorInvalidTask) {

			if errors.As(err, &structuredError) {
				return pkgError.NewError(ErrorInvalidTask).WithDetails(structuredError.GetDetails()...)
			}

			return pkgError.NewError(ErrorInvalidTask).WithDetails(err.Error())
		}

		if errors.Is(err, repo.ErrorNotFound) {

			if errors.As(err, &structuredError) {
				return pkgError.NewError(ErrorTaskNotFound).WithDetails(structuredError.GetDetails()...)
			}

			return pkgError.NewError(ErrorTaskNotFound).WithDetails(err.Error())

		}

		log.Printf("ERROR @TaskService: unexpected error: %s", err.Error())
		return pkgError.NewError(ErrorUnexpectedError)
	}

	return nil
}
