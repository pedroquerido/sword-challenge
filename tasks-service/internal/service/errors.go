package service

import "errors"

var (
	// ErrorInvalidTask represents the error obtained if a task fails to meet requirements
	ErrorInvalidTask = errors.New("invalid task")

	// ErrorTaskNotFound represents the error obtained if a task is not found based on input
	ErrorTaskNotFound = errors.New("task not found")

	// ErrorUserNotAllowed represents the error obtained if user is not allowed to perform the requested action
	ErrorUserNotAllowed = errors.New("user not allowed")

	// ErrorMissingContext represents the error obtained if Context is not found on the request
	ErrorMissingContext = errors.New("missing context")
)
