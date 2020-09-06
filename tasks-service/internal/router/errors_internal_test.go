package router

import (
	"errors"
	"net/http"
	"testing"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router/request"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router/response"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/service"
	pkgError "github.com/pedroquerido/sword-challenge/tasks-service/pkg/error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildErrorResponse(t *testing.T) {

	t.Run("should return nil", func(t *testing.T) {

		errorResponse := buildErrorResponse(nil)
		require.Nil(t, errorResponse)
	})
	t.Run("should return error response internal server error", func(t *testing.T) {

		err := errors.New("some random error")
		expectResponse := response.NewErrorResponse(http.StatusInternalServerError, messageInternal)

		errorResponse := buildErrorResponse(err)
		require.NotNil(t, errorResponse)
		assert.Equal(t, expectResponse.Code, errorResponse.Code)
		assert.Equal(t, expectResponse.Code, errorResponse.Code)
		assert.Equal(t, expectResponse.Errors, errorResponse.Errors)
	})
	t.Run("should return error response bad request - not structured error", func(t *testing.T) {

		err := request.ErrorBadRequest
		expectResponse := response.NewErrorResponse(http.StatusBadRequest, messageBadRequest, []string{request.ErrorBadRequest.Error()}...)

		errorResponse := buildErrorResponse(err)
		require.NotNil(t, errorResponse)
		assert.Equal(t, expectResponse.Code, errorResponse.Code)
		assert.Equal(t, expectResponse.Code, errorResponse.Code)
		assert.Equal(t, expectResponse.Errors, errorResponse.Errors)
	})
	t.Run("should return error response bad request - structured error", func(t *testing.T) {

		details := "some details"
		err := pkgError.NewError(request.ErrorBadRequest).WithDetails(details)
		expectResponse := response.NewErrorResponse(http.StatusBadRequest, messageBadRequest, []string{details}...)

		errorResponse := buildErrorResponse(err)
		require.NotNil(t, errorResponse)
		assert.Equal(t, expectResponse.Code, errorResponse.Code)
		assert.Equal(t, expectResponse.Code, errorResponse.Code)
		assert.Equal(t, expectResponse.Errors, errorResponse.Errors)
	})
	t.Run("should return error response forbidden - not structured error", func(t *testing.T) {

		err := service.ErrorUserNotAllowed
		expectResponse := response.NewErrorResponse(http.StatusForbidden, messageForbidden)

		errorResponse := buildErrorResponse(err)
		require.NotNil(t, errorResponse)
		assert.Equal(t, expectResponse.Code, errorResponse.Code)
		assert.Equal(t, expectResponse.Code, errorResponse.Code)
		assert.Equal(t, expectResponse.Errors, errorResponse.Errors)
	})
	t.Run("should return error response forbidden - structured error", func(t *testing.T) {

		details := "some details"
		err := pkgError.NewError(service.ErrorUserNotAllowed).WithDetails(details)
		expectResponse := response.NewErrorResponse(http.StatusForbidden, messageForbidden)

		errorResponse := buildErrorResponse(err)
		require.NotNil(t, errorResponse)
		assert.Equal(t, expectResponse.Code, errorResponse.Code)
		assert.Equal(t, expectResponse.Code, errorResponse.Code)
		assert.Equal(t, expectResponse.Errors, errorResponse.Errors)
	})
	t.Run("should return error response not found - not structured error", func(t *testing.T) {

		err := service.ErrorTaskNotFound
		expectResponse := response.NewErrorResponse(http.StatusNotFound, messageNotFound, []string{service.ErrorTaskNotFound.Error()}...)

		errorResponse := buildErrorResponse(err)
		require.NotNil(t, errorResponse)
		assert.Equal(t, expectResponse.Code, errorResponse.Code)
		assert.Equal(t, expectResponse.Code, errorResponse.Code)
		assert.Equal(t, expectResponse.Errors, errorResponse.Errors)
	})
	t.Run("should return error response not found - structured error", func(t *testing.T) {

		details := "some details"
		err := pkgError.NewError(service.ErrorTaskNotFound).WithDetails(details)
		expectResponse := response.NewErrorResponse(http.StatusNotFound, messageNotFound, []string{details}...)

		errorResponse := buildErrorResponse(err)
		require.NotNil(t, errorResponse)
		assert.Equal(t, expectResponse.Code, errorResponse.Code)
		assert.Equal(t, expectResponse.Code, errorResponse.Code)
		assert.Equal(t, expectResponse.Errors, errorResponse.Errors)
	})
	t.Run("should return error response unprocessable entity - not structured error", func(t *testing.T) {

		err := service.ErrorInvalidTask
		expectResponse := response.NewErrorResponse(http.StatusUnprocessableEntity, messageUnprocessableEntity, []string{service.ErrorInvalidTask.Error()}...)

		errorResponse := buildErrorResponse(err)
		require.NotNil(t, errorResponse)
		assert.Equal(t, expectResponse.Code, errorResponse.Code)
		assert.Equal(t, expectResponse.Code, errorResponse.Code)
		assert.Equal(t, expectResponse.Errors, errorResponse.Errors)
	})
	t.Run("should return error response unprocessable entity - structured error", func(t *testing.T) {

		details := "some details"
		err := pkgError.NewError(service.ErrorInvalidTask).WithDetails(details)
		expectResponse := response.NewErrorResponse(http.StatusUnprocessableEntity, messageUnprocessableEntity, []string{details}...)

		errorResponse := buildErrorResponse(err)
		require.NotNil(t, errorResponse)
		assert.Equal(t, expectResponse.Code, errorResponse.Code)
		assert.Equal(t, expectResponse.Code, errorResponse.Code)
		assert.Equal(t, expectResponse.Errors, errorResponse.Errors)
	})
}
