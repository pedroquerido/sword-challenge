package task_test

import (
	"errors"
	"tasks-service/pkg/rand"
	"tasks-service/pkg/task"
	"testing"
	"time"
)

const (
	defaultUUID    = "21aaa98b-842c-484a-a925-79aef811e68d"
	failureBadUUID = "21AAA98B-842C-484A-A925-79AEF811E68D"
	failureNoUUID  = "no-uuid"

	defaultUserID = "user_id"

	defaultSummaryLength = 10
	failureSummaryLength = 2501

	failedTemplate  = "FAILED - expected %v but got %v\n"
	successTemplate = "PASSED - expected %v and got %v\n"
)

var (
	testTaskValidator *task.Validator

	defaultDate time.Time = time.Now()
)

func TestValidateTaskCorrectTask(t *testing.T) {

	testTask := &task.Task{
		ID:      defaultUUID,
		UserID:  defaultUserID,
		Summary: rand.String(defaultSummaryLength),
		Date:    defaultDate,
	}
	err := testTaskValidator.Validate(testTask)

	if err != nil {
		t.Errorf(failedTemplate, nil, err)
	} else {
		t.Logf(successTemplate, nil, err)
	}
}

func TestValidateTaskEmptyTask(t *testing.T) {

	testTask := &task.Task{}
	expectedResult := task.ErrorInvalidTask

	err := testTaskValidator.Validate(testTask)

	if err != nil && errors.Is(err, task.ErrorInvalidTask) {
		t.Logf(successTemplate, expectedResult, err)
	} else {
		t.Errorf(failedTemplate, expectedResult, err)
	}
}

func TestValidateTaskEmptyID(t *testing.T) {

	testTask := &task.Task{
		UserID:  defaultUserID,
		Summary: rand.String(defaultSummaryLength),
		Date:    defaultDate,
	}
	expectedResult := task.ErrorInvalidTask

	err := testTaskValidator.Validate(testTask)

	if err != nil && errors.Is(err, task.ErrorInvalidTask) {
		t.Logf(successTemplate, expectedResult, err)
	} else {
		t.Errorf(failedTemplate, expectedResult, err)
	}
}

func TestValidateTaskIDNoUUID(t *testing.T) {

	testTask := &task.Task{
		ID:      failureNoUUID,
		UserID:  defaultUserID,
		Summary: rand.String(defaultSummaryLength),
		Date:    defaultDate,
	}
	expectedResult := task.ErrorInvalidTask

	err := testTaskValidator.Validate(testTask)

	if err != nil && errors.Is(err, task.ErrorInvalidTask) {
		t.Logf(successTemplate, expectedResult, err)
	} else {
		t.Errorf(failedTemplate, expectedResult, err)
	}
}

func TestValidateTaskIDBadUUID(t *testing.T) {

	testTask := &task.Task{
		ID:      failureBadUUID,
		UserID:  defaultUserID,
		Summary: rand.String(defaultSummaryLength),
		Date:    defaultDate,
	}
	expectedResult := task.ErrorInvalidTask

	err := testTaskValidator.Validate(testTask)

	if err != nil && errors.Is(err, task.ErrorInvalidTask) {
		t.Logf(successTemplate, expectedResult, err)
	} else {
		t.Errorf(failedTemplate, expectedResult, err)
	}
}

func TestValidateTaskEmptyUserID(t *testing.T) {

	testTask := &task.Task{
		ID:      defaultUUID,
		Summary: rand.String(defaultSummaryLength),
		Date:    defaultDate,
	}
	expectedResult := task.ErrorInvalidTask

	err := testTaskValidator.Validate(testTask)

	if err != nil && errors.Is(err, task.ErrorInvalidTask) {
		t.Logf(successTemplate, expectedResult, err)
	} else {
		t.Errorf(failedTemplate, expectedResult, err)
	}
}

func TestValidateTaskEmptySummary(t *testing.T) {

	testTask := &task.Task{
		ID:     defaultUUID,
		UserID: defaultUserID,
		Date:   defaultDate,
	}
	expectedResult := task.ErrorInvalidTask

	err := testTaskValidator.Validate(testTask)

	if err != nil && errors.Is(err, task.ErrorInvalidTask) {
		t.Logf(successTemplate, expectedResult, err)
	} else {
		t.Errorf(failedTemplate, expectedResult, err)
	}
}

func TestValidateTaskSummaryTooLong(t *testing.T) {

	testTask := &task.Task{
		ID:      defaultUUID,
		UserID:  defaultUserID,
		Summary: rand.String(failureSummaryLength),
		Date:    defaultDate,
	}
	expectedResult := task.ErrorInvalidTask

	err := testTaskValidator.Validate(testTask)

	if err != nil && errors.Is(err, task.ErrorInvalidTask) {
		t.Logf(successTemplate, expectedResult, err)
	} else {
		t.Errorf(failedTemplate, expectedResult, err)
	}
}

func TestValidateTaskEmptyDate(t *testing.T) {

	testTask := &task.Task{
		ID:      defaultUUID,
		UserID:  defaultUserID,
		Summary: rand.String(defaultSummaryLength),
	}
	expectedResult := task.ErrorInvalidTask

	err := testTaskValidator.Validate(testTask)

	if err != nil && errors.Is(err, task.ErrorInvalidTask) {
		t.Logf(successTemplate, expectedResult, err)
	} else {
		t.Errorf(failedTemplate, expectedResult, err)
	}
}
