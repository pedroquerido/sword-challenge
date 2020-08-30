package response

import (
	"testing"

	"github.com/google/uuid"
)

const (
	defaultStatusCode = 200
	defaultMessage    = "OK"

	failedTemplate  = "FAILED - expected %v but got %v\n"
	successTemplate = "PASSED - expected %v and got %v\n"
)

func TestNewBaseResponseEqualsBaseResponse(t *testing.T) {

	baseResponse := newBaseResponse(defaultMessage, defaultStatusCode)

	if baseResponse.Message != defaultMessage {
		t.Errorf(failedTemplate, defaultMessage, baseResponse.Message)
	} else {
		t.Logf(successTemplate, defaultMessage, baseResponse.Message)
	}

	if baseResponse.Code != defaultStatusCode {
		t.Errorf(failedTemplate, defaultStatusCode, baseResponse.Code)
	} else {
		t.Logf(successTemplate, defaultStatusCode, baseResponse.Code)
	}
}

func TestNewBaseResponseGeneratesID(t *testing.T) {

	const expectedID = "uuid"

	baseResponse := newBaseResponse(defaultMessage, defaultStatusCode)

	if baseResponse.ID == "" {
		t.Errorf(failedTemplate, expectedID, baseResponse.ID)
	} else {
		t.Logf(successTemplate, expectedID, baseResponse.ID)
	}
}

func TestNewBaseResponseIDisUUID(t *testing.T) {

	const expectedID = "valid uuid"

	baseResponse := newBaseResponse(defaultMessage, defaultStatusCode)

	if _, err := uuid.Parse(baseResponse.ID); err != nil {
		t.Errorf(failedTemplate, expectedID, err)
	} else {
		t.Logf(successTemplate, expectedID, baseResponse.ID)
	}
}
