package error

// Error represents a structure and set methods for better error handling
type Error struct {
	error
	details []string
}

func newError(err error) *Error {

	return &Error{
		error: err,
	}
}

// Error returns error string representation
func (c *Error) Error() string {

	return c.error.Error()
}

// Unwrap returns underlying wrapped error
func (c *Error) Unwrap() error {

	return c.error
}

// WithDetails adds details to the error
func (c *Error) WithDetails(details ...string) {

	c.details = details
}
