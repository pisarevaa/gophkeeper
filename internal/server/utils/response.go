package utils

import (
	"encoding/json"
	"net/http"

	"github.com/pisarevaa/gophkeeper/internal/server/model"
)

func SuccessResponse(w http.ResponseWriter) {
	success, err := json.Marshal(model.Success{Success: true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err = w.Write(success); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func ErrorResponse(w http.ResponseWriter, err error, statusCode int) {
	errorResponse, err := json.Marshal(model.Error{Error: err.Error()})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Error(w, string(errorResponse), statusCode)
}
