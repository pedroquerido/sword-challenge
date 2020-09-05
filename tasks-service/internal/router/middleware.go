package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router/response"
)

func setContentTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("here")

		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func recoverFromPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {

				log.Printf("recovered from panic: %v", err)

				errorResponse := response.NewErrorResponse(http.StatusInternalServerError, messageInternal)
				writeJSON(w, errorResponse.Code, errorResponse)
			}

		}()

		next.ServeHTTP(w, r)
	})
}

func requireHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		missingHeaders := make([]string, 0)
		missingHeader := "missing header %s"

		if user := r.Header.Get(headerUserID); user == "" {
			missingHeaders = append(missingHeaders, fmt.Sprintf(missingHeader, headerUserID))
		}

		if role := r.Header.Get(headerUserRole); role == "" {
			missingHeaders = append(missingHeaders, fmt.Sprintf(missingHeader, headerUserRole))
		}

		if len(missingHeaders) > 0 {
			errorResponse := response.NewErrorResponse(http.StatusBadRequest, messageBadRequest, missingHeaders...)
			writeJSON(w, errorResponse.Code, errorResponse)
		}

		next.ServeHTTP(w, r)
	})
}
