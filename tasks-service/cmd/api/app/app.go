package app

import (
	"github.com/pedroquerido/sword-challenge/service-tasks/internal/config"
	"github.com/pedroquerido/sword-challenge/service-tasks/internal/repo/gorm"
)

type taskAPI struct{}

// NewTaskAPI ...
func NewTaskAPI() *taskAPI {
	return &taskAPI{}
}

// Run ...
func (t *taskAPI) Run() error {

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

	return nil
}
