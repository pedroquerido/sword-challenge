package gorm

import (
	"errors"
	"fmt"

	"tasks-service/internal/repo"
	pkgError "tasks-service/pkg/error"
	"tasks-service/pkg/task"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _ repo.TaskRepository = (*TaskRepository)(nil)

const (
	mySQL              = "mysql"
	mySQLConnectionURL = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"

	connectionError = "error connecting to db: %w"
	migrationError  = "error migrating db: %w"
)

// TaskRepository ...
type TaskRepository struct {
	db *gorm.DB
}

// TaskRepositoryOptions ...
type TaskRepositoryOptions struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Migrate  bool
}

// NewTaskRepository ...
func NewTaskRepository(options *TaskRepositoryOptions) (*TaskRepository, error) {

	var (
		err error
		db  *gorm.DB
	)

	switch options.Driver {
	case mySQL:
		url := fmt.Sprintf(mySQLConnectionURL, options.User, options.Password, options.Host, options.Port, options.Name)
		db, err = gorm.Open(mysql.Open(url), &gorm.Config{})

	default:
		err = fmt.Errorf("unknown driver %s", options.Driver)
	}

	if err != nil {
		return nil, fmt.Errorf(connectionError, err)
	}

	if options.Migrate {
		if err := db.AutoMigrate(&taskRow{}); err != nil {
			return nil, fmt.Errorf(migrationError, err)
		}
	}

	return &TaskRepository{
		db: db,
	}, nil
}

// Save ...
func (r *TaskRepository) Save(task *task.Task) error {

	row := fromTask(task)

	if err := r.db.Create(&row).Error; err != nil {

		if errors.Is(err, gorm.ErrInvalidField) {
			return pkgError.NewError(repo.ErrorInvalidSave).WithDetails(err.Error())
		}

		return err
	}

	return nil
}

// Find ...
func (r *TaskRepository) Find(id string) (*task.Task, error) {

	row := taskRow{}

	if err := r.db.Where("task_id = ?", id).Take(&row).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkgError.NewError(repo.ErrorNotFound).WithDetails(err.Error())
		}

		return nil, err
	}

	return row.toTask(), nil
}

// Search ...
func (r *TaskRepository) Search(limit *int, offset *int64, userID *string) (tasks []*task.Task, count int64, err error) {

	rows := []*taskRow{}

	query := r.db
	countQuery := r.db.Model(&taskRow{})

	if userID != nil {
		query = query.Where("user_id = ?", *userID)
		countQuery = countQuery.Where("user_id = ?", *userID)
	}

	if limit != nil {
		query = query.Limit(*limit)
	}

	if offset != nil {
		query = query.Offset(int(*offset))
	}

	if err = query.Find(&rows).Error; err != nil {
		return nil, 0, err
	}

	tasks = make([]*task.Task, 0, len(rows))
	for _, row := range rows {
		tasks = append(tasks, row.toTask())
	}

	countQuery.Count(&count)

	return tasks, count, nil
}
