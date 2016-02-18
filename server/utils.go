package server

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func writeBadRequest(w http.ResponseWriter, message string) {
	if message == "" {
		message = "Bad Request"
	}
	writeJSON(w, http.StatusBadRequest, ErrorResponse{Code: 400, Message: message})
}

func writeNotFound(w http.ResponseWriter, message string) {
	if message == "" {
		message = "Not Found"
	}
	writeJSON(w, http.StatusNotFound, ErrorResponse{Code: 404, Message: message})
}

func writeInternalServerError(w http.ResponseWriter, e error) {
	writeJSON(w, http.StatusInternalServerError, ErrorResponse{Code: 500, Message: e.Error()})
}

func writeOk(w http.ResponseWriter, e interface{}) {
	writeJSON(w, http.StatusOK, e)
}

func writeJSON(w http.ResponseWriter, status int, e interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		panic(err)
	}
}
