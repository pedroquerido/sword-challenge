package service

import (
	"errors"
	"testing"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/repo"
	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/task"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseExternalError(t *testing.T) {

	t.Run("should return nil", func(t *testing.T) {

		err := parseExternalError(nil)
		require.Nil(t, err)
	})
	t.Run("should return unexpected error", func(t *testing.T) {

		testError := errors.New("some random error")

		err := parseExternalError(testError)
		require.NotNil(t, err)
		assert.True(t, errors.Is(err, ErrorUnexpectedError))
	})
	t.Run("should return error invalid task", func(t *testing.T) {

		testError := task.ErrorInvalidTask

		err := parseExternalError(testError)
		require.NotNil(t, err)
		assert.True(t, errors.Is(err, ErrorInvalidTask))
	})
	t.Run("should return error task not found", func(t *testing.T) {

		testError := repo.ErrorNotFound

		err := parseExternalError(testError)
		require.NotNil(t, err)
		assert.True(t, errors.Is(err, ErrorTaskNotFound))
	})
}
