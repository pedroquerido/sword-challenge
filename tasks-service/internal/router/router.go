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

// GetHTTPHandler ...
func (rt *Router) GetHTTPHandler() http.Handler {

	return rt.router
}
