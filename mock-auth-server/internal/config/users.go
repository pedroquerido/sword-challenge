package config

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Users represents the users configured in the yaml file
type Users struct {
	APIKeyUser map[string]string `yaml:"api-key-user,omitempty"`
	UserRole   map[string]string `yaml:"user-role,omitempty"`
}

func (u *Users) load() {

	const errorLoadingUsers = "error loading users: %s\n"

	filename, err := filepath.Abs("./users.yaml")
	if err != nil {
		log.Fatalf(errorLoadingUsers, err.Error())
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf(errorLoadingUsers, err.Error())
	}

	err = yaml.Unmarshal(file, u)
	if err != nil {
		log.Fatalf(errorLoadingUsers, err.Error())
	}

}
