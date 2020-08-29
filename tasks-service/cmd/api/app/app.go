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

	log.Printf("read configs: %v", cfg)

	_, err := gorm.NewTaskRepository(&gorm.TaskRepositoryOptions{
		Driver:   cfg.DB.Driver,
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Name:     cfg.DB.Name,
	})

	if err != nil {
		return err
	}

	log.Println("exiting... for now")

	blockForever()

	return nil
}

func blockForever() {
	select {}
}
