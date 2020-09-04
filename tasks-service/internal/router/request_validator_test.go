package router_test

import (
	"errors"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router/request"
	"testing"
	"time"
)

const (
	defaultSummary = "12345"

	failedTemplate  = "FAILED - expected %v but got %v\n"
	successTemplate = "PASSED - expected %v and got %v\n"
)

var (
	testRequestValidator *router.RequestValidator

	defaultDate time.Time = time.Now()
)

func TestValidateRequestCorrectCreateTaskRequestBody(t *testing.T) {

	testRequest := &request.CreateTaskRequestBody{
		Summary: defaultSummary,
		Date:    defaultDate,
	}
	err := testRequestValidator.Validate(testRequest)

	if err != nil {
		t.Errorf(failedTemplate, nil, err)
	} else {
		t.Logf(successTemplate, nil, err)
	}
}

func TestValidateRequestEmptyCreateTaskRequestBody(t *testing.T) {

	testRequest := &request.CreateTaskRequestBody{}
	expectedResult := router.ErrorBadRequest

	err := testRequestValidator.Validate(testRequest)

	if err != nil && errors.Is(err, router.ErrorBadRequest) {
		t.Logf(successTemplate, expectedResult, err)
	} else {
		t.Errorf(failedTemplate, expectedResult, err)
	}
}

func TestValidateRequestCreateTaskRequestBodyEmptySummary(t *testing.T) {

	testRequest := &request.CreateTaskRequestBody{
		Date: defaultDate,
	}
	expectedResult := router.ErrorBadRequest

	err := testRequestValidator.Validate(testRequest)

	if err != nil && errors.Is(err, router.ErrorBadRequest) {
		t.Logf(successTemplate, expectedResult, err)
	} else {
		t.Errorf(failedTemplate, expectedResult, err)
	}
}

func TestValidateRequestCreateTaskRequestBodyEmptyDate(t *testing.T) {

	testRequest := &request.CreateTaskRequestBody{
		Summary: defaultSummary,
	}
	expectedResult := router.ErrorBadRequest

	err := testRequestValidator.Validate(testRequest)

	if err != nil && errors.Is(err, router.ErrorBadRequest) {
		t.Logf(successTemplate, expectedResult, err)
	} else {
		t.Errorf(failedTemplate, expectedResult, err)
	}
}
