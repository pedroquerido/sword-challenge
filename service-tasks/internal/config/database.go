package config

import (
	"log"
	"os"
)

const (
	dbDriver   = "DB_DRIVER"
	dbHost     = "DB_HOST"
	dbPort     = "DB_PORT"
	dbUser     = "DB_USER"
	dbPassword = "DB_PASSWORD"
	dbName     = "DB_NAME"
)

// Database contains the configurations for the db connection
type Database struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func (db *Database) load() {

	if value, ok := os.LookupEnv(dbDriver); ok {
		db.Driver = value
	} else {
		log.Fatalf(envVarNotDefined, dbDriver)
	}

	if value, ok := os.LookupEnv(dbHost); ok {
		db.Host = value
	} else {
		log.Fatalf(envVarNotDefined, dbHost)
	}

	if value, ok := os.LookupEnv(dbPort); ok {
		db.Port = value
	} else {
		log.Fatalf(envVarNotDefined, dbPort)
	}

	if value, ok := os.LookupEnv(dbUser); ok {
		db.User = value
	} else {
		log.Fatalf(envVarNotDefined, dbUser)
	}

	if value, ok := os.LookupEnv(dbPassword); ok {
		db.Password = value
	} else {
		log.Fatalf(envVarNotDefined, dbPassword)
	}

	if value, ok := os.LookupEnv(dbName); ok {
		db.Name = value
	} else {
		log.Fatalf(envVarNotDefined, dbName)
	}
}
