package request_test

import (
	"errors"
	"testing"
	"time"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/go-playground/validator.v9"
)

func TestValidator_Validate_CreateTaskRequestBody(t *testing.T) {

	testRequestValidator := request.NewValidator(validator.New())

	t.Run("should return no error", func(t *testing.T) {

		testRequest := &request.CreateTaskRequestBody{
			Summary: "summary",
			Date:    time.Now(),
		}

		err := testRequestValidator.Validate(testRequest)
		require.Nil(t, err)
	})
	t.Run("should return error invalid request - empty request", func(t *testing.T) {

		err := testRequestValidator.Validate(&request.CreateTaskRequestBody{})
		require.NotNil(t, err)
		assert.True(t, errors.Is(err, request.ErrorBadRequest))
	})
	t.Run("should return error invalid request - empty summary", func(t *testing.T) {

		testRequest := &request.CreateTaskRequestBody{
			Date: time.Now(),
		}

		err := testRequestValidator.Validate(testRequest)
		require.NotNil(t, err)
		assert.True(t, errors.Is(err, request.ErrorBadRequest))
	})
	t.Run("should return error invalid request - empty date", func(t *testing.T) {

		testRequest := &request.CreateTaskRequestBody{
			Summary: "summary",
		}

		err := testRequestValidator.Validate(testRequest)
		require.NotNil(t, err)
		assert.True(t, errors.Is(err, request.ErrorBadRequest))
	})
}
