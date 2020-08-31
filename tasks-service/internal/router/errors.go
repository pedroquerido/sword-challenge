package router

import (
	"errors"
	"log"
	"net/http"
	"tasks-service/internal/router/response"
	"tasks-service/internal/service"
	pkgError "tasks-service/pkg/error"
)

const (
	messageBadRequest          = "bad request"
	messageUnprocessableEntity = "unprocessable entity"
	messageInternal            = "unknown error"
)

var (
	// ErrorBadRequest represents the error obtained from validating a request body that does not meet requirements
	ErrorBadRequest = errors.New("bad request")

	// ErrorUnknown represents the default error
	ErrorUnknown = errors.New("unknown error")
)

func parseError(err error) *response.ErrorResponse {

	if err != nil {

		var (
			structuredError pkgError.Error
		)

		if errors.As(err, &structuredError) {

			if errors.Is(err, ErrorBadRequest) {
				return response.NewErrorResponse(http.StatusBadRequest, messageBadRequest, structuredError.GetDetails()...)
			}

			if errors.Is(err, service.ErrorInvalidTask) {
				return response.NewErrorResponse(http.StatusUnprocessableEntity, messageUnprocessableEntity, structuredError.GetDetails()...)
			}
		}

	}

	log.Printf("%s: %s", messageInternal, err.Error())
	return response.NewErrorResponse(http.StatusInternalServerError, messageInternal)
}
