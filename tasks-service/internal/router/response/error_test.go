package response_test

import (
	"testing"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router/response"
	"github.com/stretchr/testify/assert"
)

func TestNewErrorResponse(t *testing.T) {

	t.Run("should return error response", func(t *testing.T) {

		message := "message"
		code := 400
		errors := []string{"error1", "error2", "error3"}

		errorResponse := response.NewErrorResponse(code, message, errors...)
		assert.NotNil(t, errorResponse)
		assert.Equal(t, message, errorResponse.Message)
		assert.Equal(t, code, errorResponse.Code)
		assert.Equal(t, errors, errorResponse.Errors)
	})
}
