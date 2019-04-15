package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api", getValue).Methods("GET")
	r.HandleFunc("/api", putValue).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
