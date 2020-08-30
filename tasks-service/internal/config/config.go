package config

import "sync"

var (
	instance Values
	once     sync.Once
)

const (
	envVarNotDefined = "env var %s not defined\n"
)

// Values ...
type Values struct {
	DB   Database
	HTTP HTTP
}

// Get ...
func Get() *Values {

	once.Do(func() {

		instance.DB.load()
		instance.HTTP.load()
	})

	return &instance
}
