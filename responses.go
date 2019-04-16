package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func sendResponse(w http.ResponseWriter, code int, payload map[string]string) {
	res, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
}

func sendError(w http.ResponseWriter, code int, message string) {
	sendResponse(w, code, map[string]string{"error": message})
}

func printMap(x map[string]string) {
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
}
