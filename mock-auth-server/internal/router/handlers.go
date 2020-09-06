package router

import (
	"net/http"
)

func (rt *Router) authHandler(w http.ResponseWriter, r *http.Request) {

	apiKey := r.Header.Get(headerAPIKey)
	if apiKey == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user, ok := rt.apiKeyUser[apiKey]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	role, ok := rt.userRole[user]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set(headerUserID, user)
	w.Header().Set(headerUserRole, role)
	w.WriteHeader(http.StatusOK)
}
