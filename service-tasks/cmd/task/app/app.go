package app

import "github.com/pedroquerido/sword-challenge/service-tasks/internal/config"

type TaskApp struct{}

func NewTaskApp() *TaskApp {
	return &TaskApp{}
}

func (t *TaskApp) Run() error {

	config.Get()

	return nil
}
