package request

import (
	"errors"
	"fmt"

	pkgError "github.com/pedroquerido/sword-challenge/tasks-service/pkg/error"

	"gopkg.in/go-playground/validator.v9"
)

var _ Validator = (*requestValidator)(nil)

// Validator ...
type Validator interface {
	Validate(req interface{}) error
}

type requestValidator struct {
	validate *validator.Validate
}

// NewValidator ...
func NewValidator(validate *validator.Validate) Validator {

	return &requestValidator{
		validate: validate,
	}
}

func (v *requestValidator) Validate(req interface{}) error {

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

			return pkgError.NewError(ErrorBadRequest).WithDetails(details...)
		}

		return err
	}

	return nil
}
