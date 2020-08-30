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
type TaskAPI struct {
	migrateDB bool
}

// NewTaskAPI ...
func NewTaskAPI(migrateDB bool) *TaskAPI {
	return &TaskAPI{
		migrateDB: migrateDB,
	}
}

// Run ...
func (t *TaskAPI) Run() error {

	// Load configs
	cfg := config.Get()

	// Setup repo
	taskRepo, err := gorm.NewTaskRepository(&gorm.TaskRepositoryOptions{
		Driver:   cfg.DB.Driver,
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Name:     cfg.DB.Name,
		Migrate:  t.migrateDB,
	})
	if err != nil {
		return err
	}

	// Create validator
	validate := validator.New()

	// Setup business layer
	service := service.NewTaskService(taskRepo, validate)

	// Setup router
	router := router.New(cfg.HTTP.Path, service, validate)

	// Create and start HTTP server
	httpServer := server.NewHTTPServer("tasks http", cfg.HTTP.Port, router.GetHTTPHandler())
	go httpServer.Run()

	// Wait for signal and cleanup afterwards
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
