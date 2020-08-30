package router

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	errorEncodingJSON = "error encoding json: %s\n"
)

func writeJSON(w http.ResponseWriter, statusCode int, data interface{}) {

	w.WriteHeader(statusCode)

	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Printf(errorEncodingJSON, err.Error())
		}
	}
}
