package router

import (
	"errors"
	"log"
	"net/http"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router/request"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router/response"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/service"
	pkgError "github.com/pedroquerido/sword-challenge/tasks-service/pkg/error"
)

const (
	messageBadRequest          = "bad request"
	messageForbidden           = "forbidden"
	messageNotFound            = "not found"
	messageUnprocessableEntity = "unprocessable entity"
	messageInternal            = "unknown error"
)

func buildErrorResponse(err error) *response.ErrorResponse {

	if err != nil {

		var (
			structuredError pkgError.Error
		)

		if errors.Is(err, request.ErrorBadRequest) {

			if errors.As(err, &structuredError) {
				return response.NewErrorResponse(http.StatusBadRequest, messageBadRequest, structuredError.GetDetails()...)
			}

			return response.NewErrorResponse(http.StatusBadRequest, messageBadRequest, err.Error())
		}

		if errors.Is(err, service.ErrorUserNotAllowed) {

			if errors.As(err, &structuredError) {
				return response.NewErrorResponse(http.StatusForbidden, messageForbidden)
			}

			return response.NewErrorResponse(http.StatusForbidden, messageForbidden)
		}

		if errors.Is(err, service.ErrorTaskNotFound) {

			if errors.As(err, &structuredError) {
				return response.NewErrorResponse(http.StatusNotFound, messageNotFound, structuredError.GetDetails()...)
			}

			return response.NewErrorResponse(http.StatusNotFound, messageNotFound, err.Error())
		}

		if errors.Is(err, service.ErrorInvalidTask) {

			if errors.As(err, &structuredError) {
				return response.NewErrorResponse(http.StatusUnprocessableEntity, messageUnprocessableEntity, structuredError.GetDetails()...)
			}

			return response.NewErrorResponse(http.StatusUnprocessableEntity, messageUnprocessableEntity, err.Error())
		}

		log.Printf("ERROR @HTTPRouter: unexpected error: %s", err.Error())
		return response.NewErrorResponse(http.StatusInternalServerError, messageInternal)
	}

	return nil
}
