package task_test

import (
	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/rand"
	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/task"
	"testing"
	"time"

	"github.com/google/uuid"
)

const (
	defaultUserID        = "user_id"
	defaultSummaryLength = 10

	failedTemplate  = "FAILED - expected %v but got %v\n"
	successTemplate = "PASSED - expected %v and got %v\n"
)

var (
	defaultDate time.Time = time.Now()
)

func TestNewTaskEqualsTask(t *testing.T) {

	testSummary := rand.String(defaultSummaryLength)

	task := task.NewTask(defaultUserID, testSummary, defaultDate)

	if task.UserID != defaultUserID {
		t.Errorf(failedTemplate, defaultUserID, task.UserID)
	} else {
		t.Logf(successTemplate, defaultUserID, task.UserID)
	}

	if task.Summary != testSummary {
		t.Errorf(failedTemplate, testSummary, task.Summary)
	} else {
		t.Logf(successTemplate, testSummary, task.Summary)
	}

	if task.Date != defaultDate {
		t.Errorf(failedTemplate, defaultDate, task.Date)
	} else {
		t.Logf(successTemplate, defaultDate, task.Date)
	}
}

func TestNewTaskGeneratesID(t *testing.T) {

	const expectedID = "uuid"

	task := task.NewTask(defaultUserID, rand.String(defaultSummaryLength), defaultDate)

	if task.ID == "" {
		t.Errorf(failedTemplate, expectedID, task.ID)
	} else {
		t.Logf(successTemplate, expectedID, task.ID)
	}
}

func TestNewTaskIDisUUID(t *testing.T) {

	const expectedID = "valid uuid"

	task := task.NewTask(defaultUserID, rand.String(defaultSummaryLength), defaultDate)

	if _, err := uuid.Parse(task.ID); err != nil {
		t.Errorf(failedTemplate, expectedID, err)
	} else {
		t.Logf(successTemplate, expectedID, task.ID)
	}
}
