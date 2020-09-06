package task

// Publisher for publishing task related events
type Publisher interface {
	PublishTaskCreated(task *Task) error
}
