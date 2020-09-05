package error

import (
	"fmt"
)

// Error represents a simple structure with set methods for a more flexible error handling
type Error struct {
	error
	details []string
}

// NewError ...
func NewError(err error) Error {

	return Error{
		error: err,
	}
}

// Error returns error string representation
func (c Error) Error() string {

	if len(c.details) == 0 {
		return fmt.Sprintf("%s", c.error.Error())
	}

	return fmt.Sprintf("%s: %v", c.error.Error(), c.details)
}

// Unwrap returns underlying wrapped error
func (c Error) Unwrap() error {

	return c.error
}

// WithDetails sets context to error and is to be used in chain
func (c Error) WithDetails(details ...string) Error {

	c.details = make([]string, len(details))
	copy(c.details, details)
	return c
}

// GetDetails gets added context from error
func (c Error) GetDetails() []string {

	return c.details
}
