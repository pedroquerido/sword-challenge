package gorm

import (
	"fmt"

	"tasks-service/internal/repo"
	"tasks-service/pkg/task"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver
)

var _ repo.TaskRepository = (*TaskRepository)(nil)

const (
	mySQL              = "mysql"
	mySQLConnectionURL = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"

	connectionError = "error connecting to db: %w"
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
		db, err = gorm.Open(mySQL, url)

	default:
		err = fmt.Errorf("unknown driver %s", options.Driver)
	}

	if err != nil {
		return nil, fmt.Errorf(connectionError, err)
	}

	return &TaskRepository{
		db: db,
	}, nil
}

// Save ...
func (r *TaskRepository) Save(task *task.Task) error {

	row := fromTask(task)
	err := row.validate()
	if err != nil {
		return err
	}

	err = r.db.Debug().Create(&row).Error
	if err != nil {
		return repo.ErrorUnknown.Wrap(err)
	}

	return nil
}

// Find ...
func (r *TaskRepository) Find(id string) (*task.Task, error) {

	row := taskRow{}

	if err := r.db.Debug().Where("id = ?", id).Take(&row).Error; err != nil {
		return &task.Task{}, repo.ErrorUnknown.Wrap(err)
	}

	return row.toTask(), nil
}

// List ...
func (r *TaskRepository) List() ([]*task.Task, error) {

	rows := []taskRow{}

	if err := r.db.Debug().Find(&rows).Error; err != nil {
		return []*task.Task{}, repo.ErrorUnknown.Wrap(err)
	}

	tasks := make([]*task.Task, 0, len(rows))
	for _, row := range rows {
		tasks = append(tasks, row.toTask())
	}

	return tasks, nil
}
