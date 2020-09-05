package request

import "errors"

var (
	// ErrorBadRequest represents the error obtained from validating a request that does not meet requirements
	ErrorBadRequest = errors.New("bad request")
)
