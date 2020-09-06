package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/pedroquerido/sword-challenge/notifications-agent/internal/config"
	"github.com/streadway/amqp"
)

const (
	// "amqp://guest:guest@localhost:5672"
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

	url := fmt.Sprintf(rabbitMQConnectionURL, cfg.RabbitMQ.User, cfg.RabbitMQ.Password, cfg.RabbitMQ.Host, cfg.RabbitMQ.Port)
	log.Println(url)

	connection, err := amqp.Dial(url)
	if err != nil {
		return err
	}
	defer connection.Close()

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
