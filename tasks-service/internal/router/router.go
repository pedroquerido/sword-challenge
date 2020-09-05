package router

import (
	"net/http"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router/request"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/service"

	"github.com/gorilla/mux"
)

const (
	routeTasks        = "/tasks"
	routeTasksID      = "/tasks/{taskID}"
	routeUsersIDTasks = "/users/{userID/tasks"

	methodPost   = "POST"
	methodGet    = "GET"
	methodPatch  = "PATCH"
	methodDelete = "DELETE"
)

// Router ...
type Router struct {
	basePath  string
	service   service.TaskService
	validator request.Validator
	router    *mux.Router
}

// New ...
func New(basePath string, service service.TaskService, validator request.Validator) *Router {

	r := mux.NewRouter()
	r.Use([]mux.MiddlewareFunc{recoverFromPanic, setContentTypeJSON, requireHeaders}...)

	router := &Router{
		basePath:  basePath,
		service:   service,
		router:    r,
		validator: validator,
	}

	return router
}

func (rt *Router) setupRoutes() {

	rt.router.HandleFunc(rt.basePath+routeTasks, rt.createTask).Methods(methodPost)

}

// GetHTTPHandler ...
func (rt *Router) GetHTTPHandler() http.Handler {

	return rt.router
}
