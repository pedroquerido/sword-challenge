package error

// Error represents a structure and set methods for better error handling
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

	return c.error.Error()
}

// Unwrap returns underlying wrapped error
func (c Error) Unwrap() error {

	return c.error
}

// SetDetails sets context to error
func (c Error) SetDetails(details ...string) {

	c.details = details
}

// WithDetails sets context to error and is to be used in chain
func (c Error) WithDetails(details ...string) Error {

	c.details = details
	return c
}

// GetDetails gets added context from error
func (c Error) GetDetails() []string {

	return c.details
}
