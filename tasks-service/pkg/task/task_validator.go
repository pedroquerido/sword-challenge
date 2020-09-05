package task

import (
	"errors"
	"fmt"

	pkgError "github.com/pedroquerido/sword-challenge/tasks-service/pkg/error"
	"gopkg.in/go-playground/validator.v9"
)

var _ Validator = (*taskValidator)(nil)

// Validator ...
type Validator interface {
	Validate(t *Task) error
}

type taskValidator struct {
	validate *validator.Validate
}

// NewValidator ...
func NewValidator(validate *validator.Validate) Validator {

	return &taskValidator{
		validate: validate,
	}
}

// Validate ...
func (v *taskValidator) Validate(t *Task) error {

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

			return pkgError.NewError(ErrorInvalidTask).WithDetails(details...)
		}

		return err
	}

	return nil
}
