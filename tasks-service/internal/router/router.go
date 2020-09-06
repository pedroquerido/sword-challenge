package router

import (
	"net/http"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router/request"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/service"

	"github.com/gorilla/mux"
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

	router.setupRoutes()

	return router
}

func (rt *Router) setupRoutes() {

	subRouter := rt.router.PathPrefix(rt.basePath).Subrouter()

	// /tasks
	tasksSubRouter := subRouter.PathPrefix("/tasks").Subrouter()
	tasksSubRouter.HandleFunc("", rt.createTask).Methods(http.MethodPost)
	tasksSubRouter.HandleFunc("", requireRoleManager(rt.listTasks)).Methods(http.MethodGet)
	tasksSubRouter.HandleFunc("/{taskID}", rt.retrieveTask).Methods(http.MethodGet)
	tasksSubRouter.HandleFunc("/{taskID}", rt.updateTask).Methods(http.MethodPatch)
	tasksSubRouter.HandleFunc("/{taskID}", requireRoleManager(rt.deleteTask)).Methods(http.MethodDelete)

	// /users
	usersSubRouter := subRouter.PathPrefix("/users").Subrouter()
	usersSubRouter.HandleFunc("/{userID}/tasks", rt.listUserTasks).Methods(http.MethodGet)
}

// GetHTTPHandler ...
func (rt *Router) GetHTTPHandler() http.Handler {

	return rt.router
}
