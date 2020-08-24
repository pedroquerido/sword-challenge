package app

type TasksApp struct{}

func NewTaskApp() *TasksApp {
	return &TasksApp{}
}

func (t *TasksApp) Run() error {

	return nil
}
