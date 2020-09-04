package response_test

import (
	"reflect"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router/response"
	"testing"
)

const (
	defaultErrorStatusCode = 400
	defaultErrorMessage    = "bad request"

	failedTemplate  = "FAILED - expected %v but got %v\n"
	successTemplate = "PASSED - expected %v and got %v\n"
)

var (
	defaultErrors = []string{"error1", "error2", "error3"}
)

func TestNewErrorResponseEqualsErrorResponse(t *testing.T) {

	errorResponse := response.NewErrorResponse(defaultErrorStatusCode, defaultErrorMessage, defaultErrors...)

	if errorResponse.Message != defaultErrorMessage {
		t.Errorf(failedTemplate, defaultErrorMessage, errorResponse.Message)
	} else {
		t.Logf(successTemplate, defaultErrorMessage, errorResponse.Message)
	}

	if errorResponse.Code != defaultErrorStatusCode {
		t.Errorf(failedTemplate, defaultErrorStatusCode, errorResponse.Code)
	} else {
		t.Logf(successTemplate, defaultErrorStatusCode, errorResponse.Code)
	}

	if reflect.DeepEqual(defaultErrors, errorResponse.Errors) {
		t.Logf(successTemplate, defaultErrors, errorResponse.Errors)
	} else {
		t.Errorf(failedTemplate, defaultErrors, errorResponse.Errors)
	}
}
