package config

import (
	"log"
	"os"
	"sync"
)

var (
	instance Values
	once     sync.Once
)

const (
	envVarNotDefined = "env var %s not defined\n"

	encryptionKey = "AES_ENCRYPTION_KEY"
)

// Values represents the structure of the application's configurations
type Values struct {
	DB            Database
	HTTP          HTTP
	RabbitMQ      RabbitMQ
	EncryptionKey string
}

// Get returns the Values singleton
func Get() *Values {

	once.Do(func() {

		// load db
		instance.DB.load()

		// load http
		instance.HTTP.load()

		// load rabbitmq
		instance.RabbitMQ.load()

		// encryption key
		if value, ok := os.LookupEnv(encryptionKey); ok {
			instance.EncryptionKey = value
		} else {
			log.Fatalf(envVarNotDefined, encryptionKey)
		}
	})

	return &instance
}
