package repo

import "errors"

var (
	// ErrorInvalidSave represents the error obtained by trying to persist a Task that does not meet requirements
	ErrorInvalidSave = errors.New("invalid save")
	// ErrorNotFound represents the error obtained by trying to find a Task that does not exist in the repo
	ErrorNotFound = errors.New("not found")
)
