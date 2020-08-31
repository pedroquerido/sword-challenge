package service

import "errors"

var (
	// ErrorInvalidTask represents the error obtained if a task fails to meet requirements
	ErrorInvalidTask = errors.New("invalid task")

	// ErrorMissingContext represents the error obtained if Context is not found on the request
	ErrorMissingContext = errors.New("missing context")
)
