package main

import (
	"encoding/json"
	"net/http"
)

// Pair is one entry into the storage
type Pair struct {
	Key   string
	Value string
}

// Sends a JSON response with given values
func sendResponse(w http.ResponseWriter, code int, payload *map[string]string) {
	res, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
}

// Creates and sends an error response
func sendError(w http.ResponseWriter, code int, message string) {
	sendResponse(w, code, &map[string]string{"error": message})
}
