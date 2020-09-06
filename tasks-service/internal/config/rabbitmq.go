package config

import (
	"log"
	"os"
)

const (
	rabbitHost     = "RABBITMQ_HOST"
	rabbitPort     = "RABBITMQ_PORT"
	rabbitUser     = "RABBITMQ_USER"
	rabbitPassword = "RABBITMQ_PASSWORD"
	exchange       = "EXCHANGE"
	routingKey     = "TASK_CREATED_ROUTING_KEY"
)

// RabbitMQ contains the rabbitmq configuration values
type RabbitMQ struct {
	Host                  string
	Port                  string
	User                  string
	Password              string
	Exchange              string
	TaskCreatedRoutingKey string
}

func (http *RabbitMQ) load() {

	if value, ok := os.LookupEnv(rabbitHost); ok {
		http.Host = value
	} else {
		log.Fatalf(envVarNotDefined, rabbitHost)
	}

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

	if value, ok := os.LookupEnv(exchange); ok {
		http.Exchange = value
	} else {
		log.Fatalf(envVarNotDefined, exchange)
	}

	if value, ok := os.LookupEnv(routingKey); ok {
		http.TaskCreatedRoutingKey = value
	} else {
		log.Fatalf(envVarNotDefined, routingKey)
	}
}
