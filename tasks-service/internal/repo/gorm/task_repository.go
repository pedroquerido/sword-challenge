package gorm

import (
	"errors"
	"time"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/repo"
	pkgError "github.com/pedroquerido/sword-challenge/tasks-service/pkg/error"
	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/task"

	"gorm.io/gorm"
)

var _ repo.TaskRepository = (*TaskRepository)(nil)

const (
	migrationError = "error migrating db: %w"
)

// TaskRepository ...
type TaskRepository struct {
	db *gorm.DB
}

// NewTaskRepository ...
func NewTaskRepository(db *gorm.DB) *TaskRepository {

	return &TaskRepository{
		db: db,
	}
}

// Insert ...
func (r *TaskRepository) Insert(task *task.Task) error {

	row := fromTask(task)

	if err := r.db.Create(&row).Error; err != nil {
		return err
	}

	return nil
}

// Find ...
func (r *TaskRepository) Find(id string) (*task.Task, error) {

	row := taskRow{}

	if err := r.db.Where("task_id = ?", id).Take(&row).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkgError.NewError(repo.ErrorNotFound)
		}

		return nil, err
	}

	return row.toTask(), nil
}

// Search ...
func (r *TaskRepository) Search(userID *string) (tasks []*task.Task, err error) {

	rows := []*taskRow{}

	query := r.db

	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}

	if err = query.Find(&rows).Error; err != nil {
		return nil, err
	}

	tasks = make([]*task.Task, 0, len(rows))
	for _, row := range rows {
		tasks = append(tasks, row.toTask())
	}

	return tasks, nil
}

// Update ...
func (r *TaskRepository) Update(id string, summary *string, date *time.Time) error {

	if summary == nil && date == nil {
		return nil
	}

	updateMap := make(map[string]interface{})

	if summary != nil {
		updateMap["summary"] = *summary
	}

	if date != nil {
		updateMap["date"] = *date
	}

	tx := r.db.Model(&taskRow{}).Where("task_id = ?", id).Updates(updateMap)

	if tx.Error != nil {

		return tx.Error
	}

	if tx.RowsAffected == 0 {

		return pkgError.NewError(repo.ErrorNotFound)
	}

	return nil
}

// Delete ...
func (r *TaskRepository) Delete(id string) error {

	tx := r.db.Where("task_id = ?", id).Delete(&taskRow{})

	if tx.Error != nil {

		return tx.Error
	}

	if tx.RowsAffected == 0 {

		return pkgError.NewError(repo.ErrorNotFound)
	}

	return nil
}
