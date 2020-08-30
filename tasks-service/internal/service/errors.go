package service

import "errors"

var (
	// ErrorInvalidTask represents the error obtained if a task fails to meet requirements
	ErrorInvalidTask = errors.New("invalid task")

	// ErrorUnknown represents the default error
	ErrorUnknown = errors.New("unknown error")
)
