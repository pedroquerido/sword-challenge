package config

import (
	"log"
	"os"
)

const (
	rabbitPort     = "RABBITMQ_PORT"
	rabbitUser     = "RABBITMQ_USER"
	rabbitPassword = "RABBITMQ_PASSWORD"
	queueName      = "NOTIFICATIONS_QUEUE_NAME"
	exchange       = "EXCHANGE"
	bindingKey     = "BINDING_KEY"
)

// RabbitMQ contains the rabbitmq configuration values
type RabbitMQ struct {
	Port       string
	User       string
	Password   string
	QueueName  string
	Exchange   string
	BindingKey string
}

func (http *RabbitMQ) load() {

	if value, ok := os.LookupEnv(rabbitPort); ok {
		http.Port = value
	} else {
		log.Fatalf(envVarNotDefined, rabbitPort)
	}

	if value, ok := os.LookupEnv(rabbitUser); ok {
		http.User = value
	} else {
		log.Fatalf(envVarNotDefined, rabbitUser)
	}

	if value, ok := os.LookupEnv(rabbitPassword); ok {
		http.Password = value
	} else {
		log.Fatalf(envVarNotDefined, rabbitPassword)
	}

	if value, ok := os.LookupEnv(queueName); ok {
		http.QueueName = value
	} else {
		log.Fatalf(envVarNotDefined, queueName)
	}

	if value, ok := os.LookupEnv(exchange); ok {
		http.Exchange = value
	} else {
		log.Fatalf(envVarNotDefined, exchange)
	}

	if value, ok := os.LookupEnv(bindingKey); ok {
		http.BindingKey = value
	} else {
		log.Fatalf(envVarNotDefined, bindingKey)
	}
}
