package service

import (
	"errors"
	"fmt"
	customError "tasks-service/pkg/error"
	"tasks-service/pkg/task"

	"gopkg.in/go-playground/validator.v9"
)

// TaskValidator ...
type TaskValidator struct {
	validate *validator.Validate
}

// NewTaskValidator ...
func NewTaskValidator(validate *validator.Validate) *TaskValidator {

	return &TaskValidator{
		validate: validate,
	}
}

// Validate validates the content of a *task.Task
//		- returns task.ErrorInvalidTask if not valid, nil if valid and customError.ErrorUnknown if an unexpected problem occurred
func (v *TaskValidator) Validate(t *task.Task) error {

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

		return customError.NewError(ErrorUnknown).WithDetails(err.Error())
	}

	return nil
}
