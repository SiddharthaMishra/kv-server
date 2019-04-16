package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// App manages the server
type App struct {

	// storage.go::Storage manages the actual internal storage
	Storage

	// router to manage the routes
	router *mux.Router
}

// Initializes the app with the required data
func initApp() *App {
	a := &App{
		Storage: Storage{hashmap: make(map[string]string)},
	}

	// creates the routes for the app
	a.createRoutes()

	return a
}

// Create routes for the server
func (a *App) createRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/api", a.postHandler).Methods("POST")
	r.HandleFunc("/api/{key}", a.getHandler).Methods("GET")
	a.router = r
}

func (a *App) postHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	// fits the request body inside a pair, i.e, a (string, string) struct
	var entry Pair
	if err := decoder.Decode(&entry); err != nil {
		// if the format is incorrect
		sendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// inserts the data
	a.Storage.putValue(&entry)

	// send back the response
	sendResponse(w, http.StatusCreated, map[string]string{"Key": entry.Key, "Value": entry.Value})
}

func (a *App) getHandler(w http.ResponseWriter, r *http.Request) {
	// gets key form url
	key := mux.Vars(r)["key"]

	// get value from storage
	value, ok := a.Storage.getValue(key)

	// if value not present
	if !ok {
		sendError(w, http.StatusNotFound, "Key not found")
		return
	}

	// send back the response
	sendResponse(w, http.StatusOK, map[string]string{"Key": key, "Value": value})

}
