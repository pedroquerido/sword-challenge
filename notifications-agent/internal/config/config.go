package config

import (
	"sync"
)

var (
	instance Values
	once     sync.Once
)

const (
	envVarNotDefined = "env var %s not defined\n"
)

// Values represents the structure of the application's configurations
type Values struct {
	RabbitMQ RabbitMQ
}

// Get returns the Values singleton
func Get() *Values {

	once.Do(func() {

		// load rabbitmq config
		instance.RabbitMQ.load()
	})

	return &instance
}
