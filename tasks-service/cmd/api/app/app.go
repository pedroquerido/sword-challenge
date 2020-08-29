package app

import (
	"log"
	"tasks-service/internal/config"

	"tasks-service/internal/repo/gorm"
)

// TaskAPI ...
type TaskAPI struct{}

// NewTaskAPI ...
func NewTaskAPI() *TaskAPI {
	return &TaskAPI{}
}

// Run ...
func (t *TaskAPI) Run() error {

	cfg := config.Get()

	_, err := gorm.NewTaskRepository(&gorm.TaskRepositoryOptions{
		Driver:   cfg.DB.Driver,
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.Port,
		Password: cfg.DB.Port,
		Name:     cfg.DB.Port,
	})

	if err != nil {
		return err
	}

	log.Println("exiting... for now")

	return nil
}
