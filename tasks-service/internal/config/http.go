package config

import (
	"log"
	"os"
)

const (
	httpPort = "HTTP_PORT"
	httpPath = "HTTP_PATH"
)

// HTTP contains the http configuration values
type HTTP struct {
	Port string
	Path string
}

func (http *HTTP) load() {

	if value, ok := os.LookupEnv(httpPort); ok {
		http.Port = value
	} else {
		log.Fatalf(envVarNotDefined, httpPort)
	}

	if value, ok := os.LookupEnv(httpPath); ok {
		http.Path = value
	} else {
		log.Fatalf(envVarNotDefined, httpPath)
	}
}
