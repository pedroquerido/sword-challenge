package task

import "errors"

var (
	// ErrorInvalidTask represents the error obtained if a task fails to meet requirements
	ErrorInvalidTask = errors.New("invalid task")

	// ErrorPublishingTaskEvent represents the error obtained if a task event failed on publish
	ErrorPublishingTaskEvent = errors.New("failed to publish task event")
)
