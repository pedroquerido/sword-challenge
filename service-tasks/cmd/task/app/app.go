package app

type TasksApp struct{}

func NewTasksApp() *TasksApp {
	return &TasksApp{}
}

func (t *TasksApp) Run() error {

	return nil
}
