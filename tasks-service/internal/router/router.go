package router

import (
	"net/http"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router/request"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/service"

	"github.com/gorilla/mux"
)

// Router ...
type Router struct {
	service   service.TaskService
	validator request.Validator
	router    *mux.Router
}

// New ...
func New(service service.TaskService, validator request.Validator) *Router {

	r := mux.NewRouter()
	r.Use([]mux.MiddlewareFunc{setContentTypeJSON, recoverFromPanic, logIncomingRequest, requireHeaders}...)

	router := &Router{
		service:   service,
		router:    r,
		validator: validator,
	}

	router.setupRoutes()

	return router
}

func (rt *Router) setupRoutes() {
	// /tasks
	tasksSubRouter := rt.router.PathPrefix("/tasks").Subrouter()
	tasksSubRouter.HandleFunc("", rt.createTask).Methods(http.MethodPost)
	tasksSubRouter.HandleFunc("", requireRoleManager(rt.listTasks)).Methods(http.MethodGet)
	tasksSubRouter.HandleFunc("/{taskID}", rt.retrieveTask).Methods(http.MethodGet)
	tasksSubRouter.HandleFunc("/{taskID}", rt.updateTask).Methods(http.MethodPatch)
	tasksSubRouter.HandleFunc("/{taskID}", requireRoleManager(rt.deleteTask)).Methods(http.MethodDelete)

	// /users
	usersSubRouter := rt.router.PathPrefix("/users").Subrouter()
	usersSubRouter.HandleFunc("/{userID}/tasks", rt.listUserTasks).Methods(http.MethodGet)
}

// GetHTTPHandler ...
func (rt *Router) GetHTTPHandler() http.Handler {

	return rt.router
}
