package router

import (
	"net/http"
	"tasks-service/internal/service"

	"github.com/gorilla/mux"
)

// Router ...
type Router struct {
	basePath string
	service  *service.TaskService
	router   *mux.Router
}

// New ...
func New(basePath string, service *service.TaskService) *Router {

	router := &Router{
		basePath: basePath,
		service:  service,
		router:   mux.NewRouter(),
	}

	router.setupRoutes()

	return router
}

func (r *Router) setupRoutes() {

}

// GetHTTPHandler ...
func (r *Router) GetHTTPHandler() http.Handler {

	return r.router
}
