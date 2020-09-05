package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router/request"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router/response"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/service"
)

const (
	detailsBadDeserialization = "error parsing request body: %s"

	messageCreateTask = "task created"
)

func (rt *Router) createTask(w http.ResponseWriter, r *http.Request) {

	// Deserialize
	body := &request.CreateTaskRequestBody{}
	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		errResponse := response.NewErrorResponse(http.StatusBadRequest, messageBadRequest, fmt.Sprintf(detailsBadDeserialization, err.Error()))
		writeJSON(w, errResponse.Code, errResponse)
		return
	}

	// Validate Content
	err = rt.validator.Validate(body)
	if err != nil {
		errResponse := parseError(err)
		writeJSON(w, errResponse.Code, errResponse)
		return
	}

	// Call Service
	isManager := false
	if r.Header.Get("x-user-role") == headerUserRoleValueManager {
		isManager = true
	}

	serviceContext := service.Context{
		UserID:    r.Header.Get("x-user-id"),
		IsManager: &isManager,
	}
	taskID, err := rt.service.Create(context.WithValue(context.Background(), service.ContextKey, serviceContext), body.Summary, body.Date)
	if err != nil {
		errResponse := parseError(err)
		writeJSON(w, errResponse.Code, errResponse)
		return
	}

	// Build Response
	successResponse := response.NewCreateTaskResponse(http.StatusOK, messageCreateTask, taskID)
	writeJSON(w, successResponse.Code, successResponse)
}

func (rt *Router) listTasks(w http.ResponseWriter, r *http.Request) {

}

func (rt *Router) listUserTasks(w http.ResponseWriter, r *http.Request) {

}

func (rt *Router) retrieveTask(w http.ResponseWriter, r *http.Request) {

}

func (rt *Router) updateTask(w http.ResponseWriter, r *http.Request) {

}

func (rt *Router) deleteTask(w http.ResponseWriter, r *http.Request) {

}
