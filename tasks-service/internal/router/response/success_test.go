package response_test

import (
	"reflect"
	"tasks-service/internal/router/response"
	"tasks-service/pkg/task"
	"testing"
	"time"
)

const (
	defaultSuccessStatusCode = 200
	defaultSuccessMessage    = "OK"

	defaultTaskID1 = "8ec7e95c-53b9-480f-9af6-2a7156de8173"
	defaultTaskID2 = "9c000a80-b55c-45c3-a14e-3219717667da"

	defaultUserID1 = "user1"
	defaultUserID2 = "user2"

	defaultSummary1 = "12345"
	defaultSummary2 = "abcdefghijkl"
)

var (
	defaultDate1 = time.Now()
	defaultDate2 = time.Now().Add(-time.Duration(10))

	defaultTask = &task.Task{
		ID:      defaultTaskID1,
		UserID:  defaultUserID1,
		Summary: defaultSummary1,
		Date:    defaultDate1,
	}

	defaultTasksList = []*task.Task{
		defaultTask,
		&task.Task{
			ID:      defaultTaskID2,
			UserID:  defaultUserID2,
			Summary: defaultSummary2,
			Date:    defaultDate2,
		},
	}
)

func TestNewCreateTaskResponseEqualsCreateTaskResponse(t *testing.T) {

	successResponse := response.NewCreateTaskResponse(defaultSuccessStatusCode, defaultSuccessMessage, defaultTaskID1)

	if successResponse.Message != defaultSuccessMessage {
		t.Errorf(failedTemplate, defaultSuccessMessage, successResponse.Message)
	} else {
		t.Logf(successTemplate, defaultSuccessMessage, successResponse.Message)
	}

	if successResponse.Code != defaultSuccessStatusCode {
		t.Errorf(failedTemplate, defaultSuccessStatusCode, successResponse.Code)
	} else {
		t.Logf(successTemplate, defaultSuccessStatusCode, successResponse.Code)
	}

	if successResponse.Data != defaultTaskID1 {
		t.Errorf(failedTemplate, defaultTaskID1, successResponse.Data)
	} else {
		t.Logf(successTemplate, defaultTaskID1, successResponse.Data)
	}
}

func TestNewListTasksResponseEqualsListTasksResponse(t *testing.T) {

	successResponse := response.NewListTasksResponse(defaultSuccessStatusCode, defaultSuccessMessage, defaultTasksList)

	if successResponse.Message != defaultSuccessMessage {
		t.Errorf(failedTemplate, defaultSuccessMessage, successResponse.Message)
	} else {
		t.Logf(successTemplate, defaultSuccessMessage, successResponse.Message)
	}

	if successResponse.Code != defaultSuccessStatusCode {
		t.Errorf(failedTemplate, defaultSuccessStatusCode, successResponse.Code)
	} else {
		t.Logf(successTemplate, defaultSuccessStatusCode, successResponse.Code)
	}

	if reflect.DeepEqual(successResponse.Data, defaultTasksList) {
		t.Logf(successTemplate, defaultTasksList, successResponse.Data)
	} else {
		t.Errorf(failedTemplate, defaultTasksList, successResponse.Data)
	}
}

func TestNewListUserTasksResponseEqualsListUserTasksResponse(t *testing.T) {

	successResponse := response.NewListUserTasksResponse(defaultSuccessStatusCode, defaultSuccessMessage, defaultTasksList)

	if successResponse.Message != defaultSuccessMessage {
		t.Errorf(failedTemplate, defaultSuccessMessage, successResponse.Message)
	} else {
		t.Logf(successTemplate, defaultSuccessMessage, successResponse.Message)
	}

	if successResponse.Code != defaultSuccessStatusCode {
		t.Errorf(failedTemplate, defaultSuccessStatusCode, successResponse.Code)
	} else {
		t.Logf(successTemplate, defaultSuccessStatusCode, successResponse.Code)
	}

	if reflect.DeepEqual(successResponse.Data, defaultTasksList) {
		t.Logf(successTemplate, defaultTasksList, successResponse.Data)
	} else {
		t.Errorf(failedTemplate, defaultTasksList, successResponse.Data)
	}
}

func TestNewRetrieveTaskResponseEqualsRetrieveTaskResponse(t *testing.T) {

	successResponse := response.NewRetrieveTaskResponse(defaultSuccessStatusCode, defaultSuccessMessage, defaultTask)

	if successResponse.Message != defaultSuccessMessage {
		t.Errorf(failedTemplate, defaultSuccessMessage, successResponse.Message)
	} else {
		t.Logf(successTemplate, defaultSuccessMessage, successResponse.Message)
	}

	if successResponse.Code != defaultSuccessStatusCode {
		t.Errorf(failedTemplate, defaultSuccessStatusCode, successResponse.Code)
	} else {
		t.Logf(successTemplate, defaultSuccessStatusCode, successResponse.Code)
	}

	if reflect.DeepEqual(successResponse.Data, defaultTask) {
		t.Logf(successTemplate, defaultTask, successResponse.Data)
	} else {
		t.Errorf(failedTemplate, defaultTask, successResponse.Data)
	}
}

func TestNewUpdateTaskResponseEqualsUpdateTaskResponse(t *testing.T) {

	successResponse := response.NewUpdateTaskResponse(defaultSuccessStatusCode, defaultSuccessMessage)

	if successResponse.Message != defaultSuccessMessage {
		t.Errorf(failedTemplate, defaultSuccessMessage, successResponse.Message)
	} else {
		t.Logf(successTemplate, defaultSuccessMessage, successResponse.Message)
	}

	if successResponse.Code != defaultSuccessStatusCode {
		t.Errorf(failedTemplate, defaultSuccessStatusCode, successResponse.Code)
	} else {
		t.Logf(successTemplate, defaultSuccessStatusCode, successResponse.Code)
	}
}

func TestNewDeleteTaskResponseEqualsDeleteTaskResponse(t *testing.T) {

	successResponse := response.NewDeleteTaskResponse(defaultSuccessStatusCode, defaultSuccessMessage)

	if successResponse.Message != defaultSuccessMessage {
		t.Errorf(failedTemplate, defaultSuccessMessage, successResponse.Message)
	} else {
		t.Logf(successTemplate, defaultSuccessMessage, successResponse.Message)
	}

	if successResponse.Code != defaultSuccessStatusCode {
		t.Errorf(failedTemplate, defaultSuccessStatusCode, successResponse.Code)
	} else {
		t.Logf(successTemplate, defaultSuccessStatusCode, successResponse.Code)
	}
}
