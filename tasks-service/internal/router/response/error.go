package response

// ErrorResponse ...
type ErrorResponse struct {
	baseResponse
	Errors []string `json:"errors,omitempty"`
}

// NewErrorResponse ...
func NewErrorResponse(statusCode int, message string, errors ...string) *ErrorResponse {

	return &ErrorResponse{
		baseResponse: newBaseResponse(message, statusCode),
		Errors:       errors,
	}
}
