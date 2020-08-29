package task_test

import (
	"os"
	"tasks-service/pkg/rand"
	"tasks-service/pkg/task"
	"testing"

	"github.com/google/uuid"

	"gopkg.in/go-playground/validator.v9"
)

func TestMain(m *testing.M) {

	testTaskValidator = task.NewValidator(validator.New())

	os.Exit(m.Run())
}

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
