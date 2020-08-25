package repo

import (
	"errors"
	"fmt"
)

// Known Error types
var (
	ErrorEmptyTask        = newError(errors.New("empty task"))
	ErrorEmptyTaskSummary = newError(errors.New("empty task summary"))
	ErrorEmptyTaskDate    = newError(errors.New("empty task date"))
)

// Error represents the error structure at repo level
type Error struct {
	error
}

func newError(err error) Error {

	return Error{
		error: err,
	}
}

// Error returns error string representation
func (c Error) Error() string {

	return c.error.Error()
}

// Unwrap returns underlying wrapped error
func (c Error) Unwrap() error {

	return c.error
}

// Wrap returns a wrapped (shadowed) error
func (c Error) Wrap(err error) Error {

	c.error = fmt.Errorf(c.Error()+": %w", err)
	return c
}
