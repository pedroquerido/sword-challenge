package router

import (
	"net/http"
)

// Router ...
type Router struct {
	router     *http.ServeMux
	apiKeyUser map[string]string
	userRole   map[string]string
}

// New ...
func New(apiKeyUser, userRole map[string]string) *Router {

	router := http.NewServeMux()

	r := &Router{
		router:     router,
		apiKeyUser: apiKeyUser,
		userRole:   userRole,
	}

	r.router.HandleFunc("/auth", r.authHandler)

	return &Router{
		router: router,
	}
}

// GetHTTPHandler ...
func (rt *Router) GetHTTPHandler() http.Handler {

	return rt.router
}
