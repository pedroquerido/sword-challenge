package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/pedroquerido/sword-challenge/notifications-agent/internal/listener/rabbitmq"

	"github.com/pedroquerido/sword-challenge/notifications-agent/internal/config"
	"github.com/streadway/amqp"
)

const (
	rabbitMQConnectionURL = "amqp://%s:%s@%s%s"
)

// Agent ...
type Agent struct {
}

// NewAgent ...
func NewAgent() *Agent {
	return &Agent{}
}

// Run ...
func (a *Agent) Run() error {

	// load configs
	cfg := config.Get()

	// setup amqp connection
	connection, err := amqp.Dial(fmt.Sprintf(rabbitMQConnectionURL, cfg.RabbitMQ.User, cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host, cfg.RabbitMQ.Port))
	if err != nil {
		return err
	}
	defer connection.Close()

	// setup listener
	eventListener, err := rabbitmq.NewEventListener(connection, cfg.RabbitMQ.Exchange, cfg.RabbitMQ.QueueName,
		[]string{cfg.RabbitMQ.BindingKey}...)
	if err != nil {
		return err
	}

	// listen
	if err := eventListener.Listen(); err != nil {
		return err
	}

	// Wait for signal and cleanup afterwards
	waitForSignal()
	return nil
}

func waitForSignal() {

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	s := <-signalChannel

	log.Printf("received interrupt signal: %v", s)
}
