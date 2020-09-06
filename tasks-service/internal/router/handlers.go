package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router/request"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router/response"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/service"
	pkgError "github.com/pedroquerido/sword-challenge/tasks-service/pkg/error"
)

const (
	detailsBadDeserialization = "error parsing request body: %s"

	messageCreateTask = "task created"
)

func (rt *Router) createTask(w http.ResponseWriter, r *http.Request) {

	// deserialize body
	body := &request.CreateTaskRequestBody{}
	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		errResponse := buildErrorResponse(pkgError.NewError(request.ErrorBadRequest).
			WithDetails(fmt.Sprintf(detailsBadDeserialization, err.Error())))
		writeJSON(w, errResponse.Code, errResponse)
		return
	}

	// validate content
	err = rt.validator.Validate(body)
	if err != nil {
		errResponse := buildErrorResponse(err)
		writeJSON(w, errResponse.Code, errResponse)
		return
	}

	// call service
	taskID, err := rt.service.Create(context.WithValue(context.Background(), service.ContextKey, buildServiceContext(r)),
		body.Summary, body.Date)
	if err != nil {
		errResponse := buildErrorResponse(err)
		writeJSON(w, errResponse.Code, errResponse)
		return
	}

	// write success response
	successResponse := response.NewCreateTaskResponse(http.StatusOK, messageCreateTask, taskID)
	writeJSON(w, successResponse.Code, successResponse)
}

func (rt *Router) listTasks(w http.ResponseWriter, r *http.Request) {

	// read query params
	userID := r.URL.Query().Get("user_id")

	// call service
	tasks, err := rt.service.List(context.WithValue(context.Background(), service.ContextKey, buildServiceContext(r)),
		&userID)
	if err != nil {
		errResponse := buildErrorResponse(err)
		writeJSON(w, errResponse.Code, errResponse)
		return
	}

	// write success response
	successResponse := response.NewListTasksResponse(http.StatusOK, messageCreateTask, tasks)
	writeJSON(w, successResponse.Code, successResponse)
}

func (rt *Router) listUserTasks(w http.ResponseWriter, r *http.Request) {

	// read path variable
	vars := mux.Vars(r)
	userID := vars["userID"]

	// call service
	tasks, err := rt.service.List(context.WithValue(context.Background(), service.ContextKey, buildServiceContext(r)),
		&userID)
	if err != nil {
		errResponse := buildErrorResponse(err)
		writeJSON(w, errResponse.Code, errResponse)
		return
	}

	// write success response
	successResponse := response.NewListTasksResponse(http.StatusOK, messageCreateTask, tasks)
	writeJSON(w, successResponse.Code, successResponse)
}

func (rt *Router) retrieveTask(w http.ResponseWriter, r *http.Request) {

	// read path variable
	vars := mux.Vars(r)
	taskID := vars["taskID"]

	// call service
	task, err := rt.service.Retrieve(context.WithValue(context.Background(), service.ContextKey, buildServiceContext(r)),
		taskID)
	if err != nil {
		errResponse := buildErrorResponse(err)
		writeJSON(w, errResponse.Code, errResponse)
		return
	}

	// write success response
	successResponse := response.NewRetrieveTaskResponse(http.StatusOK, messageCreateTask, task)
	writeJSON(w, successResponse.Code, successResponse)
}

func (rt *Router) updateTask(w http.ResponseWriter, r *http.Request) {

	// read path variable
	vars := mux.Vars(r)
	taskID := vars["taskID"]

	// deserialize body
	body := &request.UpdateTaskRequestBody{}
	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {

		errResponse := buildErrorResponse(pkgError.NewError(request.ErrorBadRequest).
			WithDetails(fmt.Sprintf(detailsBadDeserialization, err.Error())))
		writeJSON(w, errResponse.Code, errResponse)
		return
	}

	// call service
	err = rt.service.Update(context.WithValue(context.Background(), service.ContextKey, buildServiceContext(r)),
		taskID, &body.Summary, &body.Date)
	if err != nil {
		errResponse := buildErrorResponse(err)
		writeJSON(w, errResponse.Code, errResponse)
		return
	}

	// write success response
	successResponse := response.NewUpdateTaskResponse(http.StatusOK, messageCreateTask)
	writeJSON(w, successResponse.Code, successResponse)

}

func (rt *Router) deleteTask(w http.ResponseWriter, r *http.Request) {

	// read path variable
	vars := mux.Vars(r)
	taskID := vars["taskID"]

	// call service
	err := rt.service.Delete(context.WithValue(context.Background(), service.ContextKey, buildServiceContext(r)),
		taskID)
	if err != nil {
		errResponse := buildErrorResponse(err)
		writeJSON(w, errResponse.Code, errResponse)
		return
	}

	// write success response
	successResponse := response.NewDeleteTaskResponse(http.StatusOK, messageCreateTask)
	writeJSON(w, successResponse.Code, successResponse)
}

func buildServiceContext(r *http.Request) service.Context {

	isManager := false
	if r.Header.Get(headerUserRole) == headerUserRoleValueManager {
		isManager = true
	}

	return service.Context{
		UserID:    r.Header.Get(headerUserID),
		IsManager: &isManager,
	}
}
