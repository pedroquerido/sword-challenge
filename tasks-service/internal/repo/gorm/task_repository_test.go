package gorm_test

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/repo"

	"github.com/stretchr/testify/assert"

	repoGorm "github.com/pedroquerido/sword-challenge/tasks-service/internal/repo/gorm"
	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/task"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbDriver   = "DB_DRIVER"
	dbHost     = "DB_HOST"
	dbPort     = "DB_PORT"
	dbUser     = "DB_USER"
	dbPassword = "DB_PASSWORD"
	dbName     = "DB_NAME"

	mySQLConnectionURL = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"
)

var (
	testRepo *repoGorm.TaskRepository
)

func TestMain(m *testing.M) {

	url := fmt.Sprintf(mySQLConnectionURL, os.Getenv(dbUser), os.Getenv(dbPassword),
		os.Getenv(dbHost), os.Getenv(dbPort), os.Getenv(dbName))
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to db: %s", err.Error())
	}

	if err := repoGorm.CreateTables(db); err != nil {
		log.Fatalf("error creating tables: %s", err.Error())
	}

	if err := repoGorm.Populate(db); err != nil {
		log.Fatalf("error populating db: %s", err.Error())
	}

	testRepo = repoGorm.NewTaskRepository(db)

	run := m.Run()

	if err := repoGorm.DropTables(db); err != nil {
		log.Fatalf("error dropping tables: %s", err.Error())
	}

	os.Exit(run)
}

func TestTaskRepository_Insert(t *testing.T) {

	testTaskID := "9b150634-0fef-4913-b80e-29e5f3bbafec"
	testTaskUserID := "user-test"
	testTaskSummary := "summary-test"
	testTaskDate := time.Now()

	t.Run("should insert task without error", func(t *testing.T) {

		task := &task.Task{
			ID:      testTaskID,
			UserID:  testTaskUserID,
			Summary: testTaskSummary,
			Date:    testTaskDate,
		}

		err := testRepo.Insert(task)
		assert.Nil(t, err)
	})
	t.Run("should return err", func(t *testing.T) {

		testTaskDate := time.Now()

		task := &task.Task{
			ID:      testTaskID,
			UserID:  testTaskUserID,
			Summary: testTaskSummary,
			Date:    testTaskDate,
		}

		err := testRepo.Insert(task)
		assert.NotNil(t, err)
	})
}

func TestTaskRepository_Search(t *testing.T) {

	testTaskUserID := "user-test"
	noTasksUserID := "test-user"

	t.Run("should find all tasks", func(t *testing.T) {

		tasks, err := testRepo.Search(nil)
		assert.Nil(t, err)
		assert.NotNil(t, tasks)
		assert.Equal(t, 3, len(tasks))
	})
	t.Run("should find only user tasks", func(t *testing.T) {

		tasks, err := testRepo.Search(&testTaskUserID)
		assert.Nil(t, err)
		assert.NotNil(t, tasks)
		assert.Equal(t, 1, len(tasks))
	})
	t.Run("should find no tasks", func(t *testing.T) {

		tasks, err := testRepo.Search(&noTasksUserID)
		assert.Nil(t, err)
		assert.NotNil(t, tasks)
		assert.Equal(t, 0, len(tasks))
	})
}

func TestTaskRepository_Find(t *testing.T) {

	testTaskID := "9b150634-0fef-4913-b80e-29e5f3bbafec"
	testTaskUserID := "user-test"
	testTaskSummary := "summary-test"
	testTaskDate := time.Now()
	notFoundTaskID := "b13c663c-40cf-4856-bc88-45818c1439ee"

	t.Run("should find task from task id", func(t *testing.T) {

		expectedTask := &task.Task{
			ID:      testTaskID,
			UserID:  testTaskUserID,
			Summary: testTaskSummary,
			Date:    testTaskDate,
		}

		task, err := testRepo.Find(testTaskID)
		assert.Nil(t, err)
		assert.NotNil(t, task)
		assert.Equal(t, expectedTask.ID, task.ID)
		assert.Equal(t, expectedTask.UserID, task.UserID)
		assert.Equal(t, expectedTask.Summary, task.Summary)
		//assert.True(t, expectedTask.Date.Equal(task.Date)) // TODO: comeback to this
	})
	t.Run("should return err not found", func(t *testing.T) {

		task, err := testRepo.Find(notFoundTaskID)
		assert.Nil(t, task)
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, repo.ErrorNotFound))
	})
}

func TestTaskRepository_Update(t *testing.T) {

	testTaskID := "9b150634-0fef-4913-b80e-29e5f3bbafec"
	testTaskUserID := "user-test"
	updateSummary := "summary-0"
	updateDate := time.Now()
	notFoundTaskID := "b13c663c-40cf-4856-bc88-45818c1439ee"
	notFoundUserID := "test-user"

	t.Run("should update summary and date on task with task id", func(t *testing.T) {

		err := testRepo.Update(testTaskID, testTaskUserID, &updateSummary, &updateDate)
		assert.Nil(t, err)

		task, err := testRepo.Find(testTaskID)
		assert.Nil(t, err)
		assert.Equal(t, updateSummary, task.Summary)
		// assert.Equal(t, updateDate, task.Date) // TODO: comeback to this
	})
	t.Run("should return err not found with nonexisting task", func(t *testing.T) {

		err := testRepo.Update(notFoundTaskID, testTaskUserID, &updateSummary, &updateDate)
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, repo.ErrorNotFound))
	})
	t.Run("should return err not found with task from other user", func(t *testing.T) {

		err := testRepo.Update(testTaskID, notFoundUserID, &updateSummary, &updateDate)
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, repo.ErrorNotFound))
	})
}

func TestTaskRepository_Delete(t *testing.T) {

	testTaskID := "9b150634-0fef-4913-b80e-29e5f3bbafec"

	t.Run("should delete task with task id", func(t *testing.T) {

		err := testRepo.Delete(testTaskID)
		assert.Nil(t, err)

		task, err := testRepo.Find(testTaskID)
		assert.Nil(t, task)
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, repo.ErrorNotFound))
	})
	t.Run("should return err not found", func(t *testing.T) {

		err := testRepo.Delete(testTaskID)
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, repo.ErrorNotFound))
	})
}
