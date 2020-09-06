package config

import (
	"log"
	"os"
)

const (
	httpPort = "HTTP_PORT"
)

// HTTP contains the http configuration values
type HTTP struct {
	Port string
}

func (http *HTTP) load() {

	if value, ok := os.LookupEnv(httpPort); ok {
		http.Port = value
	} else {
		log.Fatalf(envVarNotDefined, httpPort)
	}
}
