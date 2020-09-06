package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/aes"
	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/task/rabbitmq"
	"github.com/streadway/amqp"

	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/task"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/config"
	repoGorm "github.com/pedroquerido/sword-challenge/tasks-service/internal/repo/gorm"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router/request"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/service"
	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/server"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gopkg.in/go-playground/validator.v9"
)

const (
	driverMySQL           = "mysql"
	mySQLConnectionURL    = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"
	rabbitMQConnectionURL = "amqp://%s:%s@%s%s"
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

	// load configs
	cfg := config.Get()

	// connect to db
	db, err := connectToDB(cfg.DB)
	if err != nil {
		return err
	}

	// setup repo
	taskRepo := repoGorm.NewTaskRepository(db)
	if t.migrateDB {
		if err := repoGorm.CreateTables(db); err != nil {
			return err
		}
	}

	// connect to rabbitmq
	connection, err := amqp.Dial(fmt.Sprintf(rabbitMQConnectionURL, cfg.RabbitMQ.User, cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host, cfg.RabbitMQ.Port))
	if err != nil {
		return err
	}
	defer connection.Close()

	// setup publisher
	taskPublisher, err := rabbitmq.NewTaskPublisher(connection, cfg.RabbitMQ.Exchange, cfg.RabbitMQ.TaskCreatedRoutingKey)
	if err != nil {
		return err
	}

	// create validators
	validate := validator.New()
	taskValidator := task.NewValidator(validate)
	requestValidator := request.NewValidator(validate)

	// create Encryptor
	encryptor, err := aes.NewEncryptor(cfg.EncryptionKey)
	if err != nil {
		return err
	}

	// setup business layer
	service := service.NewTaskService(taskRepo, taskValidator, encryptor, taskPublisher)

	// Setup router
	router := router.New(service, requestValidator)

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
	case driverMySQL:
		url := fmt.Sprintf(mySQLConnectionURL, cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
		db, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	default:
		err = fmt.Errorf("unknown driver %s", cfg.Driver)
	}

	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}

	return db, nil
}

func waitForSignal() {

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	s := <-signalChannel

	log.Printf("received interrupt signal: %v", s)
}
