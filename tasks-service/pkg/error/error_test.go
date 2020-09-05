package error_test

import (
	"errors"
	"fmt"
	"testing"

	pkgError "github.com/pedroquerido/sword-challenge/tasks-service/pkg/error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewError(t *testing.T) {
	t.Run("should return new error", func(t *testing.T) {

		testError := errors.New("some random error")
		err := pkgError.NewError(testError)
		require.NotNil(t, err)
		assert.NotEqual(t, pkgError.Error{}, err)
	})
}

func TestError_WithDetails(t *testing.T) {
	t.Run("should set details", func(t *testing.T) {

		testError := errors.New("some random error")
		details := []string{"random details"}

		err := pkgError.NewError(testError).WithDetails(details...)
		require.NotNil(t, err)
		assert.Equal(t, details, err.GetDetails())
	})
}

func TestError_GetDetails(t *testing.T) {
	t.Run("should get empty details", func(t *testing.T) {

		testError := errors.New("some random error")

		err := pkgError.NewError(testError)
		require.NotNil(t, err)
		assert.True(t, len(err.GetDetails()) == 0)
	})
	t.Run("should get correct details", func(t *testing.T) {

		testError := errors.New("some random error")
		details := []string{"random details"}

		err := pkgError.NewError(testError).WithDetails(details...)
		require.NotNil(t, err)
		assert.Equal(t, details, err.GetDetails())
	})
}

func TestError_Unwrap(t *testing.T) {
	t.Run("should return correct error", func(t *testing.T) {

		testError := errors.New("some random error")

		err := pkgError.NewError(testError)
		require.NotNil(t, err)
		assert.Equal(t, testError, err.Unwrap())
	})
}

func TestError_Error(t *testing.T) {
	t.Run("should return correct error string without details", func(t *testing.T) {

		errorString := "some random error"
		testError := errors.New(errorString)

		err := pkgError.NewError(testError)
		require.NotNil(t, err)
		assert.Equal(t, errorString, err.Error())
	})
	t.Run("should return correct error string with details", func(t *testing.T) {

		errorString := "some random error"
		details := []string{"random details"}
		testError := errors.New(errorString)

		err := pkgError.NewError(testError).WithDetails(details...)
		require.NotNil(t, err)
		assert.Equal(t, fmt.Sprintf("%s: %v", errorString, details), err.Error())
	})
}
