package main

import (
	"encoding/json"
	"net/http"
)

func sendResponse(w http.ResponseWriter, status int, v interface{}) {
	response, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func sendSuccessResponse(w http.ResponseWriter, status int, str string) {
	sendResponse(w, status, map[string]string{"success": str})
}

func sendErrorResponse(w http.ResponseWriter, status int, str string) {
	sendResponse(w, status, map[string]string{"error": str})
}

func sendNotFoundError(w http.ResponseWriter, r *http.Request) {
	sendErrorResponse(w, http.StatusNotFound, "Page not found")
}
