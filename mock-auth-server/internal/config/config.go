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
	HTTP  HTTP
	Users Users
}

// Get returns the Values singleton
func Get() *Values {

	once.Do(func() {

		// load http
		instance.HTTP.load()

		// load users config
		instance.Users.load()
	})

	return &instance
}
