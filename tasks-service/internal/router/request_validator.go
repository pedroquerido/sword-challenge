package router

import (
	"errors"
	"fmt"
	customError "tasks-service/pkg/error"

	"gopkg.in/go-playground/validator.v9"
)

// RequestValidator ...
type RequestValidator struct {
	validate *validator.Validate
}

// NewRequestValidator ...
func NewRequestValidator(validate *validator.Validate) *RequestValidator {

	return &RequestValidator{
		validate: validate,
	}
}

// Validate validates the content of a request body content
//		- returns ErrorBadRequest if not valid, nil if valid and ErrorUnknown if an unexpected problem occurred
func (v *RequestValidator) Validate(req interface{}) error {

	err := v.validate.Struct(req)
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

			return customError.NewError(ErrorBadRequest).WithDetails(details...)
		}

		return err
	}

	return nil
}
