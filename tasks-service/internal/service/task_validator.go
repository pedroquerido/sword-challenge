package service

import (
	"errors"
	"fmt"

	customError "github.com/pedroquerido/sword-challenge/tasks-service/pkg/error"
	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/task"

	"gopkg.in/go-playground/validator.v9"
)

// TaskValidator ...
type TaskValidator interface {
	// Validate validates the content of a *task.Task - returns ErrorInvalidTask if not valid, nil if valid
	Validate(t *task.Task) error
}

type taskValidator struct {
	validate *validator.Validate
}

// NewTaskValidator ...
func NewTaskValidator(validate *validator.Validate) TaskValidator {

	return &taskValidator{
		validate: validate,
	}
}

// Validate ...
func (v *taskValidator) Validate(t *task.Task) error {

	err := v.validate.Struct(t)
	if err != nil {

		var (
			validationErrors validator.ValidationErrors
		)

		if errors.As(err, &validationErrors) {

			details := make([]string, 0, len(validationErrors))
			for _, fieldErr := range validationErrors {

				if fieldErr == nil {
					continue
				}

				details = append(details, fmt.Sprintf("%s", fieldErr))
			}

			return customError.NewError(ErrorInvalidTask).WithDetails(details...)
		}

		return err
	}

	return nil
}
