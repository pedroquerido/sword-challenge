package response

import (
	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/task"
)

// CreateTaskResponse ..
type CreateTaskResponse struct {
	baseResponse
	Data string `json:"data"`
}

// NewCreateTaskResponse ...
func NewCreateTaskResponse(statusCode int, message, data string) *CreateTaskResponse {

	return &CreateTaskResponse{
		baseResponse: newBaseResponse(message, statusCode),
		Data:         data,
	}
}

// ListTasksResponse ..
type ListTasksResponse struct {
	baseResponse
	Data []*task.Task `json:"data"`
}

// NewListTasksResponse ...
func NewListTasksResponse(statusCode int, message string, data []*task.Task) *ListTasksResponse {

	return &ListTasksResponse{
		baseResponse: newBaseResponse(message, statusCode),
		Data:         data,
	}
}

// ListUserTasksResponse ..
type ListUserTasksResponse struct {
	baseResponse
	Data []*task.Task `json:"data"`
}

// NewListUserTasksResponse ...
func NewListUserTasksResponse(statusCode int, message string, data []*task.Task) *ListUserTasksResponse {

	return &ListUserTasksResponse{
		baseResponse: newBaseResponse(message, statusCode),
		Data:         data,
	}
}

// RetrieveTaskResponse ..
type RetrieveTaskResponse struct {
	baseResponse
	Data *task.Task `json:"data"`
}

// NewRetrieveTaskResponse ...
func NewRetrieveTaskResponse(statusCode int, message string, data *task.Task) *RetrieveTaskResponse {

	return &RetrieveTaskResponse{
		baseResponse: newBaseResponse(message, statusCode),
		Data:         data,
	}
}

// UpdateTaskResponse ..
type UpdateTaskResponse struct {
	baseResponse
}

// NewUpdateTaskResponse ...
func NewUpdateTaskResponse(statusCode int, message string) *UpdateTaskResponse {

	return &UpdateTaskResponse{
		baseResponse: newBaseResponse(message, statusCode),
	}
}

// DeleteTaskResponse ..
type DeleteTaskResponse struct {
	baseResponse
}

// NewDeleteTaskResponse ...
func NewDeleteTaskResponse(statusCode int, message string) *DeleteTaskResponse {

	return &DeleteTaskResponse{
		baseResponse: newBaseResponse(message, statusCode),
	}
}
