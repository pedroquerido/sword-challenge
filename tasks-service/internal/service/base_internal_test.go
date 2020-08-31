package service

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

const (
	defaultUserID   = "user1"
	defaultUserRole = "tech"

	defaultBadContext = "bad context"

	failedTemplate  = "FAILED - expected %v but got %v\n"
	successTemplate = "PASSED - expected %v and got %v\n"
)

var (
	defaultContext = Context{
		UserID:   defaultUserID,
		UserRole: defaultUserRole,
	}
)

func TestParseContextCorrectContext(t *testing.T) {

	testContext, err := parseContext(context.WithValue(context.Background(), ContextKey, defaultContext))

	if err != nil {
		t.Errorf(failedTemplate, nil, err)
	} else {
		t.Logf(successTemplate, nil, err)
	}

	if testContext == nil {
		t.Errorf(failedTemplate, defaultContext, nil)
	} else {

		if !reflect.DeepEqual(*testContext, defaultContext) {
			t.Errorf(failedTemplate, defaultContext, *testContext)
		} else {
			t.Logf(successTemplate, defaultContext, *testContext)
		}
	}
}

func TestParseContextBadContext(t *testing.T) {

	_, err := parseContext(context.WithValue(context.Background(), ContextKey, defaultBadContext))
	expectedResult := ErrorMissingContext

	if err != nil && errors.Is(err, ErrorMissingContext) {
		t.Logf(successTemplate, expectedResult, err)
	} else {
		t.Errorf(failedTemplate, expectedResult, err)
	}
}

func TestParseContextMissingUserID(t *testing.T) {

	_, err := parseContext(context.WithValue(context.Background(), ContextKey, Context{UserRole: defaultUserRole}))
	expectedResult := ErrorMissingContext

	if err != nil && errors.Is(err, ErrorMissingContext) {
		t.Logf(successTemplate, expectedResult, err)
	} else {
		t.Errorf(failedTemplate, expectedResult, err)
	}
}

func TestParseContextMissingUserRole(t *testing.T) {

	_, err := parseContext(context.WithValue(context.Background(), ContextKey, Context{UserID: defaultUserID}))
	expectedResult := ErrorMissingContext

	if err != nil && errors.Is(err, ErrorMissingContext) {
		t.Logf(successTemplate, expectedResult, err)
	} else {
		t.Errorf(failedTemplate, expectedResult, err)
	}
}
