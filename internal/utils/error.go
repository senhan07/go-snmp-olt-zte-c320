package utils

import (
	"encoding/json"
	"net/http"
)

// SendJSONResponse is a helper function to send a JSON response
func SendJSONResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

// ErrorBadRequest is a helper function to send a 400 Bad Request response
func ErrorBadRequest(w http.ResponseWriter, err error) {
	webResponse := ErrorResponse{
		Code:    http.StatusBadRequest,
		Status:  "Bad Request",
		Message: err.Error(),
	}
	SendJSONResponse(w, http.StatusBadRequest, webResponse)
}

// ErrorInternalServerError is a helper function to send a 500 Internal Server Error response
func ErrorInternalServerError(w http.ResponseWriter, err error) {
	webResponse := ErrorResponse{
		Code:    http.StatusInternalServerError,
		Status:  "Internal Server Error",
		Message: err.Error(),
	}
	SendJSONResponse(w, http.StatusInternalServerError, webResponse)
}

// ErrorNotFound is a helper function to send a 404 Not Found response
func ErrorNotFound(w http.ResponseWriter, err error) {
	webResponse := ErrorResponse{
		Code:    http.StatusNotFound,
		Status:  "Not Found",
		Message: err.Error(),
	}
	SendJSONResponse(w, http.StatusNotFound, webResponse)
}
