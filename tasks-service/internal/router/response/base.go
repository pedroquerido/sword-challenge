package response

import "github.com/google/uuid"

type baseResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func newBaseResponse(message string, code int) baseResponse {

	return baseResponse{
		ID:      uuid.New().String(),
		Message: message,
		Code:    code,
	}
}
