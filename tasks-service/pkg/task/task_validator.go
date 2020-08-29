package task

import (
	"errors"
	"fmt"
	customError "tasks-service/pkg/error"

	"gopkg.in/go-playground/validator.v9"
)

var (
	// ErrorInvalidTask represents the error obtained by validating a task that fails to meet requirements
	ErrorInvalidTask = errors.New("invalid task")
)

// Validator ...
type Validator struct {
	validate *validator.Validate
}

// NewValidator ...
func NewValidator(validate *validator.Validate) *Validator {

	return &Validator{
		validate: validate,
	}
}

// Validate ...
func (v *Validator) Validate(task *Task) error {

	err := v.validate.Struct(task)
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

		return customError.NewError(customError.ErrorUnknown)
	}

	return nil
}
