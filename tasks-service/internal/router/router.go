package router

import (
	"tasks-service/internal/service"

	"github.com/gorilla/mux"
)

// Router ...
type Router struct {
	service *service.TaskService
	router  *mux.Router
}

// New ...
func New(service *service.TaskService) *Router {

	router := &Router{
		service: service,
		router:  mux.NewRouter(),
	}

	router.setupRoutes()

	return router
}

func (r *Router) setupRoutes() {

}
