package app

import (
	"log"
	"os"
	"os/signal"
	"tasks-service/internal/config"
	"tasks-service/internal/repo/gorm"
	"tasks-service/internal/router"
	"tasks-service/internal/service"
	"tasks-service/pkg/server"

	"gopkg.in/go-playground/validator.v9"
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

	taskRepo, err := gorm.NewTaskRepository(&gorm.TaskRepositoryOptions{
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

	validate := validator.New()

	service := service.NewTaskService(taskRepo, validate)

	router := router.New(cfg.HTTP.Path, service, validate)

	httpServer := server.NewHTTPServer("tasks http", cfg.HTTP.Port, router.GetHTTPHandler())
	go httpServer.Run()

	waitForSignal()
	httpServer.Stop()

	return nil
}

func waitForSignal() {

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	s := <-signalChannel

	log.Printf("received interrupt signal: %v", s)
}
