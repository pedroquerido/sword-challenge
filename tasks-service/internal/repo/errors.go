package repo

import "errors"

var (
	// ErrorNotFound represents the error obtained by trying to find a Task that does not exist in the repo
	ErrorNotFound = errors.New("not found")
)
