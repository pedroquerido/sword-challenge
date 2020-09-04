package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/config"
	repoGorm "github.com/pedroquerido/sword-challenge/tasks-service/internal/repo/gorm"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/service"
	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/server"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gopkg.in/go-playground/validator.v9"
)

const (
	mySQL              = "mysql"
	mySQLConnectionURL = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"

	dbConnectionError = "error connecting to db: %w"
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

	// Connect to db
	db, err := connectToDB(cfg.DB)
	if err != nil {
		return err
	}

	// Setup repo
	taskRepo := repoGorm.NewTaskRepository(db)

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

func connectToDB(cfg config.Database) (db *gorm.DB, err error) {

	switch cfg.Driver {
	case mySQL:
		url := fmt.Sprintf(mySQLConnectionURL, cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
		db, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	default:
		err = fmt.Errorf("unknown driver %s", cfg.Driver)
	}

	if err != nil {
		return nil, fmt.Errorf(dbConnectionError, err)
	}

	return db, nil
}

func waitForSignal() {

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	s := <-signalChannel

	log.Printf("received interrupt signal: %v", s)
}
