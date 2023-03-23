package api

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func writeJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func writeErrorResponse(w http.ResponseWriter, status int, err error, message string) {
	response := errorResponse{
		Error:   err.Error(),
		Message: message,
	}
	writeJSONResponse(w, status, response)
}